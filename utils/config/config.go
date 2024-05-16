package config

import (
	"encoding/json"
	"flag"
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
	ID      string `json:"id,omitempty" mapstructure:"id" structs:"id"`
	Type    string `json:"type,omitempty"`
	DataDir string `json:"data_dir,omitempty"`
	// Todo: use json.RawMessage instead. Only unmarshal after determining type
	Config map[string]interface{} `json:"config,omitempty"`
}

type Config struct {
	Mode            ApplicationMode  `json:"mode" mapstructure:"app_mode" structs:"app_mode" env:"APP_MODE"`
	ListenUrl       string           `json:"listen_url" mapstructure:"listen_url" structs:"listen_url" env:"LISTEN_URL"`
	SearcherDataDir string           `json:"searcher_data_dir" mapstructure:"searcher_data_dir" structs:"searcher_data_dir" env:"SEARCHER_DATA_DIR"`
	SearcherConfigs []SearcherConfig `json:"searcher_configs" mapstructure:"searcher_configs" structs:"searcher_configs"`
}

func (c *Config) GetSearcherConfig(id string) (*SearcherConfig, error) {
	for _, sc := range c.SearcherConfigs {
		if sc.ID == id {
			return &sc, nil
		}
	}
	return nil, fmt.Errorf("searcher config not found: %s", id)
}

var defaultConfig = Config{
	Mode:            Dev,
	ListenUrl:       "127.0.0.1:8080",
	SearcherDataDir: "data",
	SearcherConfigs: []SearcherConfig{
		{
			ID:   "ddg",
			Type: "DuckDuckGo",
		},
	},
}

func (sc SearcherConfig) GetDataDir() string {
	if sc.DataDir != "" {
		return sc.DataDir
	}
	normalizedName := regexp.MustCompile(`/[^\w ]+/g`).ReplaceAllString(sc.ID, "")
	return path.Join(AppConfig.SearcherDataDir, normalizedName)
}

func parseConfig() Config {
	configFilePath := flag.String("config", "", "path to config file")

	defaultsAsMap := structs.Map(defaultConfig)
	for key, value := range defaultsAsMap {
		viper.SetDefault(key, value)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")
	if (*configFilePath) != "" {
		viper.SetConfigFile(*configFilePath)
	}
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Println("Error reading config file:", err)
	}
	config := Config{}

	err := viper.Unmarshal(&config)

	if err != nil {
		panic(fmt.Errorf("fatal error parsing config file: %w", err))
	}

	if content, err := json.MarshalIndent(config, "", "  "); err == nil {
		log.Printf("parsed config: \n%s\n", string(content))
	}

	return config
}

var AppConfig = parseConfig()
