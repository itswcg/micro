package dao

import (
	"context"
	"database/sql"
	"github.com/prometheus/common/log"
	"leaf/internal/model"
	"time"
)

const (
	_all       = "select biz_tag,max_id,step,`desc`,update_time from segment"
	_maxSeq    = "select max_id from segment where biz_tag=?"
	_updateMax = "update segment set max_id=? where biz_tag=? and max_id=?"
)

// All get all seq
func (d *dao) All(c context.Context) (seqs map[string]*model.Segment, err error) {
	rows, err := d.db.Query(c, _all)
	if err != nil {
		log.Error("d.db.Query(%s) error(%v)", _all, err)
		return
	}
	seqs = make(map[string]*model.Segment)
	defer rows.Close()
	for rows.Next() {
		s := new(model.Segment)
		if err = rows.Scan(&s.BizTag, &s.MaxSeq, &s.Step, &s.Desc, &s.UpdateTime); err != nil {
			log.Error("rows.Scan error(%v)")
			return
		}
		s.LastTimestamp = time.Now()
		seqs[s.BizTag] = s
	}
	return
}

// MaxSeq return current max seq.
func (d *dao) MaxSeq(c context.Context, bizTag string) (maxSeq int64, err error) {
	row := d.db.QueryRow(c, _maxSeq, bizTag)
	if err = row.Scan(&maxSeq); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
	}
	return
}

// UpMaxSeq update max seq by businessID.
func (d *dao) UpMaxSeq(c context.Context, bizTag string, maxSeq, lastSeq int64) (rows int64, err error) {
	res, err := d.db.Exec(c, _updateMax, maxSeq, bizTag, lastSeq)
	if err != nil {
		log.Error("d.db.Exec(%s, %d, %d) error(%v)", _updateMax, maxSeq, bizTag)
		return
	}
	return res.RowsAffected()
}
