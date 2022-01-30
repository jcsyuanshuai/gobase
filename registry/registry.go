package registry

import (
	"github.com/xx/gobase/config"
)

const DefaultRegistryType = "etcd"

var defaultClient Client

var clientMap map[string]Client

func Init() {
	registryType := config.GlobalConfig.Registry.Type

	if registryType == "etcd" {
		defaultClient = clientMap[registryType]
	}
}

func SetClient(name string, cli Client) {
	clientMap[name] = cli
}
