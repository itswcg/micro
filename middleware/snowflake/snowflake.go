package snowflake

import (
	"errors"
	"sync"
	"time"
)

// fork https://github.com/holdno/snowFlakeByGo

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	epoch       int64 = 1572883200000 // 25岁生日了
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workId    int64
	number    int64
}

func New(workId int64) (*Worker, error) {
	if workId < 0 || workId >= workerMax {
		return nil, errors.New("workId error")
	}

	return &Worker{
		timestamp: 0,
		workId:    0,
		number:    0,
	}, nil
}

func (w *Worker) GetID() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++

		if w.number > numberMax {
			for w.timestamp >= now {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.timestamp = now
		w.number = 0
	}

	ID := int64((now-epoch)<<timeShift | (w.workId << workerShift) | (w.number))
	return ID
}
