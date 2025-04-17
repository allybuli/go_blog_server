package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	initialize.OtherInit()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.Es = initialize.ConnectEs()

	initialize.InitCron()

	core.RunServer()
}
