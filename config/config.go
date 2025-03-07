package config

import (
	"errors"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Log      LogConfig
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type JWTConfig struct {
	Secret string
	Expire int
}

var GlobalConfig Config

func (c *Config) Validate() error {
	if c.Server.Port == "" {
		return errors.New("服务器端口不能为空")
	}
	if c.Database.Host == "" || c.Database.Port == "" {
		return errors.New("数据库配置不完整")
	}
	if c.JWT.Secret == "" {
		return errors.New("JWT密钥不能为空")
	}
	return nil
}

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		return err
	}

	return GlobalConfig.Validate()
}

func WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&GlobalConfig); err != nil {
		}
	})
}
