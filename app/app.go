package app

import (
	"github.com/xx/gobase/util"
	"os"
	"os/signal"
	"sync"
)

type Runnable interface {
	Start() error
	Init(...OptionFunc) error
	Option() Option
	Run() error
	Stop() error
}

type app struct {
	opts Option
	once sync.Once
}

var defaultApp *app

func NewApp(opts ...OptionFunc) *app {
	app := new(app)
	if len(opts) == 0 {
		app.opts = DefaultOptions()
	} else {
		app.opts = newOptions(opts...)
	}
	return app
}

func Start() error {
	return defaultApp.Start()
}

func Default(opts ...OptionFunc) {
	defaultApp = NewApp(opts...)
}

func Init(opts ...OptionFunc) error {
	err := defaultApp.Init(opts...)
	if err != nil {
		return err
	}
	return nil
}

func GetOption() Option {
	return defaultApp.Option()
}

func Stop() error {
	return defaultApp.Stop()
}

func (a *app) Option() Option {
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

func (a *app) Init(opts ...OptionFunc) error {
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

func startHttpEngine(opts Option) error {
	go func() {
		err := opts.Engine.Run(":8080")
		if err != nil {
			opts.ErrorCh <- err
		}
	}()

	return nil
}

func stopHttpEngine(opts Option) error {
	return nil
}
