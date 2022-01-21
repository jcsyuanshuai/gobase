package config

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
