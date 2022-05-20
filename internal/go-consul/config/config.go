package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"strings"
	"time"
)

type Configuration struct {
	App         App `yaml:"app"`
	Consul      Consul
	ConsulExtra ConsulExtra `yaml:"consul_extra"`
}

type App struct {
	Name    string `yaml:"name"`
	Profile string `yaml:"profile"`
	Timeout int    `yaml:"timeout"`
	Port    int    `yaml:"port"`
}

type Consul struct {
	Enabled bool `yaml:"enabled"`
}

type ConsulExtra struct {
	Env1 string `yaml:"env1"`
	Env2 int    `yaml:"env2"`
}

func LoadConfig() (config *Configuration, e error) {
	initial := time.Now()

	logrus.Infof("[Env Config] Initializing env variable configurations")

	viper.SetConfigName("configs/application")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Load Config from File
	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			logrus.Warnf("No Config file found, loaded config from Environment - Default path ./conf")
		default:
			return nil, errors.Wrap(err, "config.LoadConfig")
		}
	}

	keys2 := viper.AllSettings()
	fmt.Println(keys2)

	// Load Config from Consul
	if viper.GetBool("consul.enabled") {
		viper.AddRemoteProvider(viper.GetString("consul.provider"), viper.GetString("consul.address"), viper.GetString("consul.keys_prefix"))
		viper.SetConfigType("yaml")

		if err := viper.ReadRemoteConfig(); err != nil {
			logrus.Error("[Env Config] Error serializing config", err)
			return nil, errors.Wrap(err, "[Env Config] Error serializing config from external provider")
		}
	}

	keys := viper.AllSettings()
	fmt.Println(keys)

	err := viper.Unmarshal(&config)

	if err != nil {
		logrus.Error("[Env Config] Error serializing config", err)
		return nil, errors.Wrap(err, "[Env Config] Error serializing config")
	}

	delta := time.Since(initial).Milliseconds()
	logrus.Infof(fmt.Sprintf("[Env Config] Finalized env variable configurations in %dus", delta))

	return config, nil
}
