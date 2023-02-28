package bootstrap

import (
	"os"
	"os/signal"

	qqbot "github.com/Logiase/MiraiGo-Template/bot"
	qconfig "github.com/Logiase/MiraiGo-Template/config"
	log "github.com/sirupsen/logrus"
	"github.com/wechatgpt/wechatbot/config"
	_ "github.com/wechatgpt/wechatbot/handler/qq"
)

func StartQQChat() {
	if _, err := os.Stat("device.json"); err != nil {
		if os.IsNotExist(err) {
			log.Warn("没有找到device.json,正在生成")
			qqbot.GenRandomDevice()
		} else {
			log.Fatal(err)
		}
	}
	//环境变量初始化配置
	// 	bot:
	//   login-method: "qrcode" # qrcode/common 二维码登录/账号密码登录
	//   account: ""
	//   password: ""
	qconfig.GlobalConfig.SetConfigType("yaml")
	account := config.GetQQUin()
	password := config.GetQQPassword()
	qconfig.GlobalConfig.SetDefault("bot.login-method", "qrcode")
	qconfig.GlobalConfig.SetDefault("bot.account", account)
	qconfig.GlobalConfig.SetDefault("bot.password", password)
	qqbot.StartService()
	qqbot.UseProtocol(qqbot.AndroidPhone)
	if err := qqbot.Login(); err != nil {
		panic(err)
	}
	qqbot.RefreshList()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	qqbot.Stop()
}
