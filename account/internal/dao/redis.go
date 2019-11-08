package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/prometheus/common/log"
)

const (
	_prefixName = "name:%s"
)

func keyName(name string) string {
	return fmt.Sprintf(_prefixName, name)
}

func (d *dao) SetName(ctx context.Context, name string, mid int64) (err error) {
	key := keyName(name)
	con := d.redis.Get(ctx)
	defer con.Close()

	if err = con.Send("SET", key, mid); err != nil {
		log.Error("con.Send error(%v)", err)
		return
	}

	if err = con.Flush(); err != nil {
		log.Error("con.Flush error(%v)", err)
	}

	for i := 0; i < 2; i++ {
		if _, err = con.Receive(); err != nil {
			log.Error("con.Receive error(%v)", err)
		}
	}
	return
}

func (d *dao) GetMidByName(ctx context.Context, name string) (mid int64, err error) {
	key := keyName(name)
	con := d.redis.Get(ctx)
	defer con.Close()

	if mid, err = redis.Int64(con.Do("GET", key)); err != nil {
		if err == redis.ErrNil {
			err = nil
		} else {
			log.Error("con.Do(GET %s) error(%v)", key, err)
		}
	}
	return
}
