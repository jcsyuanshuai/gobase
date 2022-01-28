package app

import (
	"fmt"
)

func InitConfig(opts *Options) {
	fmt.Printf("init config\n")

}

func InitMysql(opts *Options) {
	fmt.Printf("init mysql\n")
}

func InitDatabase(opts *Options) {
	fmt.Printf("init database\n")

}

func InitRedis(opts *Options) {
	fmt.Printf("init redis\n")

}

func InitLogger(opts *Options) {
	fmt.Printf("init logger\n")
}

func InitCache(opts *Options) {
	fmt.Printf("init cache\n")
}
