package util

import "os"

var envs = []string{
	"CONF_PATH",
	"MODE",
}

func GetEnv(name string) string {
	return os.Getenv(name)
}
