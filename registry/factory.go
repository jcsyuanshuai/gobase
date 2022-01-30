package registry

type Factory interface {
	Create() Client
}
