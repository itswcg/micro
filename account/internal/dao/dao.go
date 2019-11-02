package dao

import (
	"context"
	"github.com/itswcg/micro/account/api"
	"time"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	GetInfo(ctx context.Context, mid int64) (r *api.Info, err error)
	GetProfile(ctx context.Context, mid int64) (r *api.Profile, err error)
	GetPassword(ctx context.Context, mid int64) (password string, err error)
	AddInfo(ctx context.Context, mid int64, name, password string) (err error)
	SetInfo(ctx context.Context, mid int64, field, value string) (err error)
}

// dao dao.
type dao struct {
	db          *sql.DB
	redis       *redis.Pool
	redisExpire int32
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a dao and return.
func New() Dao {
	var (
		dc struct {
			Mysql *sql.Config
		}
		rc struct {
			Redis       *redis.Config
			RedisExpire xtime.Duration
		}
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	checkErr(paladin.Get("redis.toml").UnmarshalTOML(&rc))
	return &dao{
		// mysql
		db: sql.NewMySQL(dc.Mysql),
		// redis
		redis:       redis.NewPool(rc.Redis),
		redisExpire: int32(time.Duration(rc.RedisExpire) / time.Second),
	}
}

// Close close the resource.
func (d *dao) Close() {
	d.redis.Close()
	d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	if err = d.pingRedis(ctx); err != nil {
		return
	}
	return d.db.Ping(ctx)
}

func (d *dao) pingRedis(ctx context.Context) (err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}
