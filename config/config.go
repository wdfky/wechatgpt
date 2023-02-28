package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	ChatGpt ChatGptConfig `json:"chatgpt" mapstructure:"chatgpt" yaml:"chatgpt"`
}

type ChatGptConfig struct {
	Token         string  `json:"token,omitempty"  mapstructure:"token,omitempty"  yaml:"token,omitempty"`
	Wechat        *string `json:"wechat,omitempty" mapstructure:"wechat,omitempty" yaml:"wechat,omitempty"`
	WechatKeyword *string `json:"wechat_keyword"   mapstructure:"wechat_keyword"   yaml:"wechat_keyword"`
	Telegram      *string `json:"telegram"         mapstructure:"telegram"         yaml:"telegram"`
	TgWhitelist   *string `json:"tg_whitelist"     mapstructure:"tg_whitelist"     yaml:"tg_whitelist"`
	TgKeyword     *string `json:"tg_keyword"       mapstructure:"tg_keyword"       yaml:"tg_keyword"`
	QQ            *string `json:"qq"               mapstructure:"qq"               yaml:"qq"`
	QQUin         *string `json:"qq_uin"           mapstructure:"qq_uin"           yaml:"qq_uin"`
	QQPassword    *string `json:"qq_password"      mapstructure:"qq_password"      yaml:"qq_password"`
	QQKeyword     *string `json:"qq_keyword"       mapstructure:"qq_keyword"       yaml:"qq_keyword"`
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./local")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}

func GetWechat() *string {
	wechat := getEnv("wechat")

	if wechat != nil {
		return wechat
	}
	if config == nil {
		return nil
	}
	if wechat == nil {
		wechat = config.ChatGpt.Wechat
	}
	return wechat
}

func GetWechatKeyword() *string {
	keyword := getEnv("wechat_keyword")

	if keyword != nil {
		return keyword
	}
	if config == nil {
		return nil
	}
	if keyword == nil {
		keyword = config.ChatGpt.WechatKeyword
	}
	return keyword
}

func GetTelegram() *string {
	tg := getEnv("telegram")
	if tg != nil {
		return tg
	}
	if config == nil {
		return nil
	}
	if tg == nil {
		tg = config.ChatGpt.Telegram
	}
	return tg
}

func GetTelegramKeyword() *string {
	tgKeyword := getEnv("tg_keyword")

	if tgKeyword != nil {
		return tgKeyword
	}
	if config == nil {
		return nil
	}
	if tgKeyword == nil {
		tgKeyword = config.ChatGpt.TgKeyword
	}
	return tgKeyword
}

func GetTelegramWhitelist() *string {
	tgWhitelist := getEnv("tg_whitelist")

	if tgWhitelist != nil {
		return tgWhitelist
	}
	if config == nil {
		return nil
	}
	if tgWhitelist == nil {
		tgWhitelist = config.ChatGpt.TgWhitelist
	}
	return tgWhitelist
}
func GetQQ() *string {
	qq := getEnv("qq")
	if qq != nil {
		return qq
	}
	if config == nil {
		return nil
	}
	if qq == nil {
		qq = config.ChatGpt.QQ
	}
	return qq
}

func GetQQUin() *string {
	qqUin := getEnv("qq_uin")
	if qqUin != nil {
		return qqUin
	}
	if config == nil {
		return nil
	}
	if qqUin == nil {
		qqUin = config.ChatGpt.QQUin
	}
	return qqUin
}

func GetQQPassword() *string {
	qqPassword := getEnv("qq_password")
	if qqPassword != nil {
		return qqPassword
	}
	if config == nil {
		return nil
	}
	if qqPassword == nil {
		qqPassword = config.ChatGpt.QQPassword
	}
	return qqPassword
}

func GetQQKeyword() *string {
	qqKeyword := getEnv("qq_keyword")
	if qqKeyword != nil {
		return qqKeyword
	}
	if config == nil {
		return nil
	}
	if qqKeyword == nil {
		qqKeyword = config.ChatGpt.QQKeyword
	}
	return qqKeyword
}

func GetOpenAiApiKey() *string {
	apiKey := getEnv("api_key")

	if apiKey != nil {
		return apiKey
	}

	if config == nil {
		return nil
	}
	if apiKey == nil {
		apiKey = &config.ChatGpt.Token
	}
	return apiKey
}

func getEnv(key string) *string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = os.Getenv(strings.ToUpper(key))
	}

	if len(value) > 0 {
		return &value
	}

	if config == nil {
		return nil
	}

	if len(value) > 0 {
		return &value
	} else if config.ChatGpt.WechatKeyword != nil {
		value = *config.ChatGpt.WechatKeyword
	}
	return nil
}
