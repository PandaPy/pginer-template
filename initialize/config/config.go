package config

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var AppConfig Config

// LoadConfig loads the configuration for the entire application
func Init(env string) {
	color.Green("初始化配置文件")

	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		zap.L().Fatal("fatal error config file", zap.Error(err))
	}

	// 将配置内容解析到结构体
	if err := viper.Unmarshal(&AppConfig); err != nil {
		zap.L().Fatal("[unable to decode into struct]", zap.Error(err))
	}

	// 参数存在性校验
	validateConfig()
}

func validateConfig() {
	requiredKeys := []string{
		"SERVER.LISTEN",
		"SERVER.MODE",
		"SERVER.ALLOWED_HOSTS",
		"SERVER.ROUTER_PREFIX",
		"SERVER.SECRET_KEY",

		"DATABASES.default.NAME",
		"DATABASES.default.USER",
		"DATABASES.default.PASSWORD",
		"DATABASES.default.HOST",
		"DATABASES.default.PORT",

		"REDIS.HOST",
		"REDIS.USER",
		"REDIS.PASSWORD",
	}

	// 检查每一个必需的键是否在配置文件中设置
	for _, key := range requiredKeys {
		if !viper.IsSet(key) {
			zap.L().Fatal("config validation error", zap.Error(fmt.Errorf("missing required configuration key: %s", key)))
		}
	}
}
