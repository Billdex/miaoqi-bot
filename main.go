package main

import (
	"fmt"
	"log"
	"miaoqi-bot/config"
	"miaoqi-bot/models"
	"miaoqi-bot/services"
	"miaoqi-bot/utils/logger"
)

func main() {
	// 初始化配置文件
	err := config.InitConfig()
	if err != nil {
		log.Println("读取配置文件出错！", err)
		return
	}
	fmt.Println("配置文件加载完毕")

	// 初始化logger
	err = logger.InitLog(config.AppConfig.Log.Style, config.AppConfig.Log.File, config.AppConfig.Log.Level)
	if err != nil {
		log.Println("初始化logger出错！", err)
		return
	}
	defer logger.Sync()
	logger.Info("初始化logger完毕")

	// 初始化数据库引擎
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local",
		config.AppConfig.DB.User,
		config.AppConfig.DB.Password,
		config.AppConfig.DB.Host,
		config.AppConfig.DB.Port,
		config.AppConfig.DB.Database,
	)
	err = models.SetUp(connStr)
	if err != nil {
		logger.Error("数据库连接出错!", err)
		return
	}
	logger.Info("初始化数据库引擎完毕")

	s := services.SetUp()

	err = s.Run(fmt.Sprintf(":%d", config.AppConfig.Server.Port))
	if err != nil {
		logger.Error("Bot启动失败")
	}
}
