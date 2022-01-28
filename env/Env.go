package env

import (
	"fmt"
	"github.com/xx/gobase/util"
)

type Env struct {
	Mode     string
	BasePath string
	ConfPath string
}

var defaultEnv *Env = &Env{
	Mode:     "dev",
	BasePath: util.CurrentAbsPath(),
	ConfPath: fmt.Sprintf("%s/%s/%s", util.CurrentAbsPath(), "conf", "dev"),
}

func DefaultEnv() *Env {
	return defaultEnv
}

func LoadEnv(env *Env) {
	env = defaultEnv

	if util.FileExist(
		fmt.Sprintf("%s/.env", env.BasePath)) {
		//todo
	}

	mode := util.GetEnv("MODE")
	confPath := util.GetEnv("CONF_PATH")

	if confPath != "" {
		env.ConfPath = confPath
	}

	if mode != "" {
		env.Mode = mode
	}

}
