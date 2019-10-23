package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"leaf/api"
	"leaf/internal/dao"
	"leaf/internal/model"
	"sync"
)

// Service service.
type Service struct {
	ac   *paladin.Map
	dao  dao.Dao
	seqs map[string]*model.Segment
	bl   sync.RWMutex
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	s = &Service{
		ac:  ac,
		dao: dao.New(),
	}
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

// next id
func (s *Service) NextID(ctx context.Context, req *api.LeafReq) (res *api.LeafReply, err error) {
	//
}
