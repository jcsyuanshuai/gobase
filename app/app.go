package app

import (
	"github.com/xx/gobase/util"
	"os"
	"os/signal"
	"sync"
)

type Runnable interface {
	Start() error
	Init(...Option) error
	Options() Options
	Run() error
	Stop() error
}

type app struct {
	opts Options
	once sync.Once
}

func NewApp(opts ...Option) *app {
	app := new(app)
	if len(opts) == 0 {
		app.opts = DefaultOptions()
	} else {
		app.opts = newOptions(opts...)
	}
	return app
}

func (a *app) Options() Options {
	return a.opts
}

func (a *app) Start() error {
	var err error
	for _, fn := range a.opts.BeforeStart {
		if err = fn(); err != nil {
			return err
		}
	}
	if err = startHttpEngine(a.opts); err != nil {
		return err
	}

	for _, fn := range a.opts.AfterStart {
		if err = fn(); err != nil {
			return err
		}
	}
	return nil
}

func (a *app) Init(opts ...Option) error {
	for _, o := range opts {
		o(&a.opts)
	}
	return nil
}

func (a *app) Stop() error {
	var err error
	for _, fn := range a.opts.AfterStop {
		if err = fn(); err != nil {
			return err
		}
	}

	if err = stopHttpEngine(a.opts); err != nil {
		return err
	}

	for _, fn := range a.opts.AfterStop {
		if err = fn(); err != nil {
			return err
		}
	}
	return nil
}

func (a *app) Run() error {
	ch := make(chan os.Signal, 1)
	if a.opts.Signal {
		signal.Notify(ch, util.Shutdown()...)
	}
	select {
	case <-ch:
	case <-a.opts.Context.Done():
	}
	return a.Stop()
}

var _ Runnable = &app{}

func startHttpEngine(opts Options) error {
	go func() {
		err := opts.Engine.Run(":8080")
		if err != nil {
			opts.ErrorCh <- err
		}
	}()

	return nil
}

func stopHttpEngine(opts Options) error {
	return nil
}
