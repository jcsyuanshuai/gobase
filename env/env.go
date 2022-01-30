package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xx/gobase/util"
)

type Env struct {
	Mode     string
	BasePath string
	ConfPath string
}

var defaultEnv *Env = &Env{
	Mode:     "prod",
	BasePath: util.CurrentAbsPath(),
	ConfPath: fmt.Sprintf("%s/%s/%s", util.CurrentAbsPath(), "conf", "dev"),
}

func DefaultEnv() *Env {
	return defaultEnv
}

func LoadEnv(env *Env) {
	env = defaultEnv

	if envPath := fmt.Sprintf("%s/.env", env.BasePath); util.FileExist(envPath) {
		err := godotenv.Load(envPath)
		if err != nil {
			fmt.Printf("cat't find .env file in path: %s", envPath)
		}
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
