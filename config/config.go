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
	defaultConfigName = "config.yml"
)


// cfgFile is the location of the config file
var cfgFile string
var config *viper.Viper


func InitConfig()  {

	var err error
	cfgFile, err = findConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	config = viper.New()
	config.SetConfigName(defaultConfigName)
	config.AddConfigPath(cfgFile)
	if _, err := os.Stat(cfgFile); err == nil {
		if err := config.ReadInConfig(); err != nil {
			log.Fatalln("reading config failed:", err)
		}
	}
}


func findConfig() (string, error) {
	if cfgFile != "" {
		return cfgFile, nil
	}
	return filepath.Join(configPath(), defaultConfigName), nil
}

func configPath() string {
	wd,err := os.Getwd()
	if err != nil {
		log.Fatalln("can't determine config path", err)
	}
	return fmt.Sprintf("%s/%s", wd, defaultConfigName)
}


// Get fetches a generic value from viper config
func Get(key string) interface{} {

	// parse the config file if it's not already
	if config == nil {
		InitConfig()
	}

	return config.Get(key)
}

// Viper returns the viper config object
func Viper() *viper.Viper {

	// parse the config file if it's not already
	if config == nil {
		 InitConfig()
	}

	return config
}
