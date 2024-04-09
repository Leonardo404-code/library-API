package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	Load()
}

func GetOrDefault(envName, defaultValue string) string {
	value := viper.GetString(envName)

	if !viper.IsSet(envName) {
		return defaultValue
	}

	return value
}

func GetString(envName string) string {
	if !viper.IsSet(envName) {
		panic(fmt.Sprintf("env variable %s is not set", envName))
	}

	return viper.GetString(envName)
}

func Load() {
	cwd, _ := os.Getwd()
	var pathTransversal string

	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
			resultPath := filepath.Join(cwd) + "/config/env"
			viper.AddConfigPath(resultPath)
			break
		}
		cwd = filepath.Dir(cwd)
		pathTransversal = filepath.Join(pathTransversal, "..")
		if cwd == "/" || cwd == "\\" {
			log.Printf("Couldn`t find go.mod file")
			break
		}
	}

	viper.AddConfigPath("./config/env")
	viper.SetConfigName("envs.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Couldn't read config file. Error: %v\n", err.Error())
	}

	viper.AutomaticEnv()
}
