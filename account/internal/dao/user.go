package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/itswcg/micro/account/api"
	"github.com/prometheus/common/log"
)

const (
	_userSharding int64 = 2
)

const (
	_selInfoById     = "select mid,name,sex,face from user_%d where mid=?"
	_selProfileById  = "select mid,name,sex,face,email,phone,join_time from user_%d where mid=?"
	_selPasswordById = "select password from user_%d where mid=?"

	_insertInfo = "insert into user_%d(mid,name,password) values (?,?,?) on duplicate key update name=?,password=?"
	_updateInfo = "update user_%d set %s=? where mid=?"
)

func (d *dao) hit(mid int64) int64 {
	return mid % _userSharding
}

func (d *dao) GetInfo(ctx context.Context, mid int64) (r *api.Info, err error) {
	r = &api.Info{}
	row := d.db.QueryRow(ctx, fmt.Sprintf(_selInfoById, d.hit(mid)), mid)
	if err = row.Scan(&r.Mid, &r.Name, &r.Sex, &r.Face); err != nil {
		if err == sql.ErrNoRows {
			r = nil
			err = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
	}
	return
}

func (d *dao) GetProfile(ctx context.Context, mid int64) (r *api.Profile, err error) {
	r = &api.Profile{}
	row := d.db.QueryRow(ctx, fmt.Sprintf(_selProfileById, d.hit(mid)), mid)
	if err = row.Scan(&r.Mid, &r.Name, &r.Sex, &r.Face, &r.Email, &r.Phone, &r.JoinTime); err != nil {
		if err == sql.ErrNoRows {
			r = nil
			err = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
	}
	return
}

func (d *dao) GetPassword(ctx context.Context, mid int64) (password string, err error) {
	row := d.db.QueryRow(ctx, fmt.Sprintf(_selPasswordById, d.hit(mid)), mid)
	if err = row.Scan(&password); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
	}
	return
}

func (d *dao) AddInfo(ctx context.Context, mid int64, name, password string) (err error) {
	_, err = d.db.Exec(ctx, fmt.Sprintf(_insertInfo, d.hit(mid)), mid, name, password, name, password)
	if err != nil {
		log.Error("db.Exec error(%v)", err)
	}
	return
}

func (d *dao) SetInfo(ctx context.Context, mid int64, field, value string) (err error) {
	_, err = d.db.Exec(ctx, fmt.Sprintf(_updateInfo, d.hit(mid), field), value, mid)
	if err != nil {
		log.Error("db.Exec error(%v)", err)
	}
	return
}
