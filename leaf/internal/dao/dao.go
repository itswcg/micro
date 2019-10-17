package dao

import (
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
}

// dao dao.
type dao struct {
	db *sql.DB
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
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))

	return &dao{
		// mysql
		db: sql.NewMySQL(dc.Mysql),
	}
}

// Close close the resource.
func (d *dao) Close() {
	d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return d.db.Ping(ctx)
}
