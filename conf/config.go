package conf

import (
	"encoding/json"
	"fmt"
	"github.com/xx/gobase/conf/local"
	"github.com/xx/gobase/util"
)

var GlobalConfig *Config

type Config struct {
	*local.App
	*local.ConfigCenter
	*local.Registry
}

func (a *Config) String() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func init() {
	GlobalConfig = &Config{
		App:          new(local.App),
		ConfigCenter: new(local.ConfigCenter),
		Registry:     new(local.Registry),
	}

	params := []struct {
		path string
		out  interface{}
	}{
		{
			path: fmt.Sprintf(""),
			out:  GlobalConfig.App,
		},
	}

	for _, item := range params {
		err := util.YamlLoad(item.path, item.out)
		if err != nil {
			return
		}
	}
}

func GetConfigCenter() *local.ConfigCenter {
	return GlobalConfig.ConfigCenter
}

func GetRegistry() *local.Registry {
	return GlobalConfig.Registry
}
