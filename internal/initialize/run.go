package initialize

import (
	"fmt"

	"example.com/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "Success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}