package model

import (
	"sync"
	"time"
)

type Segment struct {
	BizTag        string
	CurSeq        int64
	MaxSeq        int64
	Step          int64
	Desc          string
	UpdateTime    time.Time
	LastTimestamp time.Time
	Mutex         sync.Mutex
}
