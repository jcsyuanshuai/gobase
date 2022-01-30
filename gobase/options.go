package gobase

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xx/gobase/config"
	"gorm.io/gorm"
)

type Options struct {
	Engine *gin.Engine
	Config *config.Config
	DB     *gorm.DB

	ErrorCh chan error
	// Before and After Functions
	BeforeStart []func() error
	AfterStart  []func() error
	BeforeStop  []func() error
	AfterStop   []func() error

	// context
	Context context.Context

	Signal bool
}

type Option func(options *Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		Engine:  gin.New(),
		Context: context.Background(),
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func DefaultOptions() Options {
	opt := newOptions(
		// hooks
		InitConfig,
		InitMysql,
		InitLogger,
		InitCache,
		// routers
		//SetupSwaggerRouters,
		//SetupUserRouters,
		//SetupPermissionRouters,
		//SetupClientRouters,
		//SetupTokenRouters,
		// others
		AfterStart(func() error {
			fmt.Printf("after start\n")
			return nil
		}),
		AfterStop(func() error {
			fmt.Printf("after stop\n")
			return nil
		}),
		BeforeStart(func() error {
			fmt.Printf("before start\n")
			return nil
		}),
		BeforeStop(func() error {
			fmt.Printf("before stop\n")
			return nil
		}),
	)
	return opt
}

// BeforeStart Run functions before services start
func BeforeStart(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

// AfterStart Run functions after services start
func AfterStart(fn func() error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

// BeforeStop Run functions before services stop
func BeforeStop(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

// AfterStop Run functions after services stop
func AfterStop(fn func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
