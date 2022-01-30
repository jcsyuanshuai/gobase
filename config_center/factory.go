package config_center

type Factory interface {
	Create() Client
}
