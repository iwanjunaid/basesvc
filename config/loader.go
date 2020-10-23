package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/iwanjunaid/basesvc/config/db"
	"github.com/iwanjunaid/basesvc/config/logging"
	"github.com/iwanjunaid/basesvc/config/server"
	"github.com/iwanjunaid/basesvc/shared/logger"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Server   server.ServerList
	Database db.DatabaseList
	Logger   logging.LoggerConfig
}

var cfg Config

func init() {
	// Load configuration
	config := LoadConfig()

	// hot reload on config change...
	go func() {
		config.WatchConfig()
		config.OnConfigChange(func(e fsnotify.Event) {
			logger.WithFields(logger.Fields{"component": "configuration"}).Infof("config file changed %v", e.Name)
		})
	}()
}

func LoadConfig() *viper.Viper{
	config := viper.New()
	dir, err := os.Getwd()

	if err != nil {
		panic(fmt.Errorf("failed get directory config: %v", err))
	}

	// load server configuration
	config.AddConfigPath(dir + "/config/server")
	config.SetConfigType("yaml")
	config.SetConfigName("rest.yml")
	err = config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load server config: %v", err))
	}

	// load logging configuration
	config.AddConfigPath(dir + "/config/logging")
	config.SetConfigType("yaml")
	config.SetConfigName("logger.yml")
	err = config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load server config: %v", err))
	}

	// load database configuration
	config.AddConfigPath(dir + "/config/db")
	config.SetConfigName("mysql.yml")
	err = config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load database config: %v", err))
	}

	config.AutomaticEnv()
	config.Unmarshal(&cfg)

	return config
}

// GetConfig value
func GetConfig() *Config {
	return &cfg
}
