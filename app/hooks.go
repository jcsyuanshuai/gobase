package app

import (
	"fmt"
	"github.com/xx/gobase/config"
)

func InitConfig(opt *Option) {
	config.Init()

	fmt.Printf("%v\n", config.GlobalConfig)

	//config_center.Init()
	////
	//registry.Init()
}

func InitMysql(opt *Option) {
	//config_center.Get(context.TODO(), "")
}

func InitMongo(opt *Option) {
	//config_center.Get(context.TODO(), "")
}

func InitRedis(opt *Option) {
	//config_center.Get(context.TODO(), "")

}

func InitLogger(opt *Option) {
	//config_center.Get(context.TODO(), "")
}

func InitCache(opt *Option) {
	//config_center.Get(context.TODO(), "")
}
