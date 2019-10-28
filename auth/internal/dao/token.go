package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/prometheus/common/log"
)

const (
	_prefixToken = "account:token:%s"
	_tokenExpire = 60 * 60 * 24 * 15
)

func keyToken(token string) string {
	return fmt.Sprintf(_prefixToken, token)
}

func (d *dao) SetToken(ctx context.Context, token string, mid int64) (err error) {
	key := keyToken(token)
	con := d.redis.Get(ctx)
	defer con.Close()

	if err = con.Send("SET", key, mid); err != nil {
		log.Error("con.Send error(%v)", err)
		return
	}

	if err = con.Send("EXPIRE", key, _tokenExpire); err != nil {
		log.Error("con.Send error(%v)", err)
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

func (d *dao) GetMidByToken(ctx context.Context, token string) (mid int64, err error) {
	key := keyToken(token)
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
