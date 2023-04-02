package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// ReadConfig -- read the config file regard to the file type (file extension)
// Also, viper can watch changes on the file --> allow us to hot reload the application
func ReadConfig(fileName string, configPaths ...string) bool {
	viper.SetConfigName(fileName)
	if len(configPaths) < 1 {
		// look for current dir
		viper.AddConfigPath(".")
	} else {
		for _, configPath := range configPaths {
			viper.AddConfigPath(configPath)
		}
	}
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Cannot read config file. %v", err)
		return false
	}

	return true
}

// ReadConfigByFile -- read config file by file path
func ReadConfigByFile(file string) bool {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Cannot read config file. %v", err)
		return false
	}

	return true
}
