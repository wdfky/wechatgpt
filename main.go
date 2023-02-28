package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/wechatgpt/wechatbot/bootstrap"
	"github.com/wechatgpt/wechatbot/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Warn("没有找到配置文件，尝试读取环境变量")
	}
	wechatEnv := config.GetWechat()
	telegramEnv := config.GetTelegram()
	qqEnv := config.GetQQ()
	if wechatEnv != nil && *wechatEnv == "true" {
		bootstrap.StartWebChat()
	} else if telegramEnv != nil {
		bootstrap.StartTelegramBot()
	} else if qqEnv != nil {
		bootstrap.StartQQChat()
	}
}
