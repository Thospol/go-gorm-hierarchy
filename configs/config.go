package config

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// CF -> for use configs model
	CF = &Configs{}
)

// SQLConfig sql config model
type SQLConfig struct {
	Host            string        `mapstructure:"HOST"`
	Port            int           `mapstructure:"PORT"`
	Username        string        `mapstructure:"USERNAME"`
	Password        string        `mapstructure:"PASSWORD"`
	DatabaseName    string        `mapstructure:"DATABASE_NAME"`
	DriverName      string        `mapstructure:"DRIVER_NAME"`
	Enable          bool          `mapstructure:"ENABLE"`
	MaxIdleConns    int           `mapstructure:"MAX_IDLE_CONNS"`
	MaxOpenConns    int           `mapstructure:"MAX_OPEN_CONNS"`
	ConnMaxLifetime time.Duration `mapstructure:"MAX_LIFE_TIME"`
}

// Configs config models
type Configs struct {
	Database SQLConfig `mapstructure:"DATABASE"`
}

// InitConfig init config
func InitConfig(configPath string) error {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.AutomaticEnv()
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		logrus.Error("read config file error:", err)
		return err
	}

	if err := bindingConfig(v, CF); err != nil {
		logrus.Error("binding config error:", err)
		return err
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := bindingConfig(v, CF); err != nil {
			logrus.Error("binding error:", err)
			return
		}
	})

	return nil
}

// bindingConfig binding config
func bindingConfig(vp *viper.Viper, cf *Configs) error {
	if err := vp.Unmarshal(&cf); err != nil {
		logrus.Error("unmarshal config error:", err)
		return err
	}

	return nil
}
