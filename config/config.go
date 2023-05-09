package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig `yaml:"app"`
	TG  TgConfig  `yaml:"tg"`
}

type TgConfig struct {
	WebhookMethod     string `yaml:"webhook_method"`
	SendMessageMethod string `yaml:"send_message_method"`
}

type AppConfig struct {
	WebhookURL              string `yaml:"webhook_url"`
	PasswordRetentionMinute int    `yaml:"password_retention_minute"`
	Port                    int    `yaml:"port"`
}

var C Config

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	C.App.WebhookURL = viper.GetStringMap("app")["webhook_url"].(string)
	C.App.Port = viper.GetStringMap("app")["port"].(int)
	C.TG.WebhookMethod = viper.GetStringMap("tg")["webhook_method"].(string)
	C.TG.SendMessageMethod = viper.GetStringMap("tg")["send_message_method"].(string)

	return nil
}
