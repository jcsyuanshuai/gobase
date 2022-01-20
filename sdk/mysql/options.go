package mysql

import (
	"time"
)

type Options struct {
	Driver       string
	Username     string
	Password     string
	Host         string
	Port         string
	Database     string
	MaxIdleConn  int
	MaxOpenConn  int
	MaxLifetime  time.Duration
	MaxIdleTime  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

func (o *Options) Validate() error {
	return nil
}

type Option func(options *Options)

func NewOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func DefaultOptions() Options {
	return Options{}
}
