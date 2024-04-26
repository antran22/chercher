package config

import (
	"chercher/search"
	"fmt"
	"github.com/spf13/viper"
	"path"
	"regexp"
)

type GoEnv string

const (
	Prod GoEnv = "prod"
	Dev  GoEnv = "dev"
)

type SearcherConfig struct {
	RootConfig *Config
	Name       string
	Type       search.SearcherType
	DataDir    string
	Config     map[string]interface{}
}

type Config struct {
	Env             GoEnv
	SearcherDataDir string
	SearcherConfigs []SearcherConfig
}

func (sc SearcherConfig) getDataDir() string {
	if sc.DataDir != "" {
		return sc.DataDir
	}
	normalizedName := regexp.MustCompile(`/[^\w ]+/g`).ReplaceAllString(sc.Name, "")
	return path.Join(sc.RootConfig.SearcherDataDir, normalizedName)
}

func parseConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	config := Config{}

	err := viper.Unmarshal(&config)

	if err != nil {
		panic(fmt.Errorf("fatal error parsing config file: %w", err))
	}

	for i := range config.SearcherConfigs {
		config.SearcherConfigs[i].RootConfig = &config
	}

	return config
}

//var AppConfig = parseConfig()
