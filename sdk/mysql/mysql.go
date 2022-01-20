package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Engine struct {
	opts Options
	*gorm.DB
}

func (o *Engine) SetDatabase(opts Options) {
	err := o.opts.Validate()
	if err != nil {
		return
	}

	db, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		return
	}

	err = sqlDb.Ping()
	if err != nil {
		return
	}

	sqlDb.SetMaxIdleConns(o.opts.MaxIdleConn)
	sqlDb.SetMaxOpenConns(o.opts.MaxOpenConn)
	sqlDb.SetConnMaxIdleTime(o.opts.MaxIdleTime)
	sqlDb.SetConnMaxLifetime(o.opts.MaxLifetime)
}
