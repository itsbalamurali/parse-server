package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"os"
	"fmt"
	"log"
)


const (
	// defaultConfigName is the name of the config file when no alternative is supplied.
	defaultConfigName = "config.toml"
)


// cfgFile is the location of the config file
var cfgFile string


func InitConfig()  {

	var err error
	cfgFile, err = findConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	viper.SetConfigName(defaultConfigName)
	viper.AddConfigPath(cfgFile)
	viper.AutomaticEnv()

	if _, err := os.Stat(cfgFile); err == nil {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln("reading config failed:", err)
		}
	}
}


func findConfig() (string, error) {
	if cfgFile != "" {
		return cfgFile, nil
	}

	ch := configHome()
	if err := os.MkdirAll(ch, 0755); err != nil {
		return "", err
	}

	return filepath.Join(ch, defaultConfigName), nil
}

func configPath() string {
	return fmt.Sprintf("%s/%s", configHome(), defaultConfigName)
}


func configHome() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = filepath.Join(homeDir(), ".config", "doctl")
	}

	return configHome
}


func homeDir() string {
	return os.Getenv("HOME")
}