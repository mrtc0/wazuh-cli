package command

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

type BasicAuth struct {
	Username string
	Password string
}

type Config struct {
	APIURL    string
	BasicAuth BasicAuth
}

func GetConfig() (*Config, error) {
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(filepath.Join("$HOME", ".config", "wazuh-cli"))
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func ShowConfig(config Config) {
	fmt.Printf("%-20v%-20v\n", "API URL:", config.APIURL)
	fmt.Printf("%-20v%-20v\n", "BasicAuth User:", config.BasicAuth.Username)
	fmt.Printf("%-20v%-20v\n", "BasicAuth Pass:", config.BasicAuth.Password)
}
