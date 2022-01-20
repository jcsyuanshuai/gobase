package local

import "encoding/json"

type App struct {
	BasePath  string
	ConfPath  string
	Namespace string
	Name      string
	Version   string
	Host      string
	Ip        string
	HttpPort  int
	GrpcPort  int
}

func (a *App) String() string {
	ret, _ := json.Marshal(a)
	return string(ret)
}
