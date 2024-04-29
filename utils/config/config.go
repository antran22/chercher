package config

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/spf13/viper"
	"log"
	"path"
	"regexp"
)

type ApplicationMode string

const (
	Prod ApplicationMode = "prod"
	Dev  ApplicationMode = "dev"
)

type SearcherConfig struct {
	RootConfig *Config                `json:"-"`
	Name       string                 `json:"name,omitempty" mapstructure:"name" structs:"name"`
	Type       string                 `json:"type,omitempty"`
	DataDir    string                 `json:"data_dir,omitempty"`
	Config     map[string]interface{} `json:"config,omitempty"`
}

type Config struct {
	Mode            ApplicationMode  `json:"mode" mapstructure:"app_mode" structs:"app_mode" env:"APP_MODE"`
	ListenUrl       string           `json:"listen_url" mapstructure:"listen_url" structs:"listen_url" env:"LISTEN_URL"`
	SearcherDataDir string           `json:"searcher_data_dir" mapstructure:"searcher_data_dir" structs:"searcher_data_dir" env:"SEARCHER_DATA_DIR"`
	SearcherConfigs []SearcherConfig `json:"searcher_configs" mapstructure:"searcher_configs" structs:"searcher_configs"`
}

var defaultConfig = Config{
	Mode:            Dev,
	ListenUrl:       "127.0.0.1:8080",
	SearcherDataDir: "data",
	SearcherConfigs: []SearcherConfig{
		{
			Name: "ddg",
			Type: "DuckDuckGo",
		},
	},
}

func (sc SearcherConfig) getDataDir() string {
	if sc.DataDir != "" {
		return sc.DataDir
	}
	normalizedName := regexp.MustCompile(`/[^\w ]+/g`).ReplaceAllString(sc.Name, "")
	return path.Join(sc.RootConfig.SearcherDataDir, normalizedName)
}

func parseConfig() Config {

	defaultsAsMap := structs.Map(defaultConfig)
	for key, value := range defaultsAsMap {
		viper.SetDefault(key, value)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
	config := Config{}

	err := viper.Unmarshal(&config)

	if err != nil {
		panic(fmt.Errorf("fatal error parsing config file: %w", err))
	}

	for i := range config.SearcherConfigs {
		config.SearcherConfigs[i].RootConfig = &config
	}

	if content, err := json.MarshalIndent(config, "", "  "); err == nil {
		log.Printf("parsed config: \n%s\n", string(content))
	}

	return config
}

var AppConfig = parseConfig()
