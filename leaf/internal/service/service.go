package service

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/ecode"
	"leaf/api"
	"leaf/internal/dao"
	"leaf/internal/model"
	"sync"
	"time"
)

const _nextStepRetry = 3

// Service service.
type Service struct {
	ac            *paladin.Map
	dao           dao.Dao
	seqs          map[string]*model.Segment
	bl            sync.RWMutex // 读写锁
	lastTimestamp time.Time
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	s = &Service{
		ac:            ac,
		dao:           dao.New(),
		seqs:          make(map[string]*model.Segment),
		lastTimestamp: time.Now(),
	}
	s.loadSeqs()
	//go s.loadProc()
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}

// load seq from mysql
func (s *Service) loadSeqs() (err error) {
	var seqs map[string]*model.Segment
	if seqs, err = s.dao.All(context.TODO()); err != nil {
		return
	}
	for tag, seq := range seqs {
		s.bl.Lock()
		if _, ok := s.seqs[tag]; !ok {
			s.seqs[tag] = seq
			s.bl.Unlock()
			continue
		}
		s.bl.Unlock()
	}
	return
}

func (s *Service) loadProc() {
	for {
		time.Sleep(time.Duration(5))
		s.loadSeqs()
	}
}

// next id
func (s *Service) NextID(ctx context.Context, req *api.LeafReq) (res *api.LeafReply, err error) {
	s.bl.RLock()
	seq, ok := s.seqs[req.BizTag]
	s.bl.RUnlock()
	if !ok {
		err = ecode.NothingFound
		return
	}
	seq.Mutex.Lock()
	if seq.CurSeq == 0 || seq.CurSeq+1 > seq.MaxSeq {
		if err = s.nextStep(ctx, seq); err != nil {
			seq.Mutex.Unlock()
			return
		}
	}
	seq.CurSeq++
	id := int64(seq.CurSeq)
	seq.Mutex.Unlock()

	return &api.LeafReply{Id: id}, nil
}

func (s *Service) nextStep(ctx context.Context, seq *model.Segment) (err error) {
	var (
		n             int
		lastSeq, rows int64
	)
	for {
		if lastSeq, err = s.dao.MaxSeq(ctx, seq.BizTag); err != nil {
			return
		}

		if rows, err = s.dao.UpMaxSeq(ctx, seq.BizTag, lastSeq+seq.Step, lastSeq); err != nil {
			return
		}

		if rows > 0 {
			seq.CurSeq = lastSeq
			seq.MaxSeq = lastSeq + seq.Step
			break
		}

		if n++; n > _nextStepRetry {
			err = fmt.Errorf("get the next step failed(%s)", seq.BizTag)
			return
		}
	}
	return
}
