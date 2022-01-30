package registry

type Client interface {
	Discover() error

	Register(metas ...Metadata) error

	Unregister(meta Metadata) error
}

type Metadata interface {
	Metadata()
}

type GrpcMeta struct {
}
