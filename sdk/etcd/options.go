package etcd

type Options struct {
	Endpoint string
	Username string
	Password string
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
	return Options{
		Endpoint: "127.0.0.1:2379",
		Username: "",
		Password: "",
	}
}

func SetEndpoint(endpoint string) Option {
	return func(options *Options) {
		options.Endpoint = endpoint
	}
}

func SetUsername(username string) Option {
	return func(options *Options) {
		options.Username = username
	}
}

func SetPassword(password string) Option {
	return func(options *Options) {
		options.Password = password
	}
}
