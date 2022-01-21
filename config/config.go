package config

import (
	"encoding/json"
	"fmt"
	"github.com/xx/gobase/env"
	"github.com/xx/gobase/util"
)

var GlobalConfig *Config

type Config struct {
	*env.Env
	*App
	*ConfigCenter
	*Registry
}

func (a *Config) String() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func Init() {
	GlobalConfig = &Config{
		Env:          new(env.Env),
		App:          new(App),
		ConfigCenter: new(ConfigCenter),
		Registry:     new(Registry),
	}

	env.LoadEnv(GlobalConfig.Env)

	checkAndLoad(GlobalConfig.Env, GlobalConfig)
}

func checkAndLoad(env *env.Env, config *Config) {
	params := []struct {
		path string
		out  interface{}
	}{
		{
			path: fmt.Sprintf("%s/%s/%s", env.ConfPath, env.Mode, "app.yaml"),
			out:  &config.App,
		},
		{
			path: fmt.Sprintf("%s/%s/%s", env.ConfPath, env.Mode, "config.yaml"),
			out:  &config.App,
		},
		{
			path: fmt.Sprintf("%s/%s/%s", env.ConfPath, env.Mode, "registry.yaml"),
			out:  &config.App,
		},
	}

	for _, param := range params {
		err := util.YamlLoad(param.path, param.out)
		if err != nil {
			return

		}
	}
}

func GetConfigCenter() *ConfigCenter {
	return GlobalConfig.ConfigCenter
}

func GetRegistry() *Registry {
	return GlobalConfig.Registry
}
