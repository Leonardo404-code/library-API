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

func GetOrDefaultInt(envName string, defaultValue int) int {
	value := viper.GetInt(envName)

	if !viper.IsSet(envName) {
		return defaultValue
	}

	return value
}

func GetOrDefaultInt64(envName string, defaultValue int64) int64 {
	value := viper.GetInt64(envName)

	if !viper.IsSet(envName) {
		return defaultValue
	}

	return value
}

func GetOrDefaultFloat64(envName string, defaultValue float64) float64 {
	value := viper.GetFloat64(envName)

	if !viper.IsSet(envName) {
		return defaultValue
	}

	return value
}

func GetStringSlice(envName string) []string {
	if !viper.IsSet(envName) {
		panic(fmt.Sprintf("env variable %s is not set", envName))
	}

	return viper.GetStringSlice(envName)
}

func GetString(envName string) string {
	if !viper.IsSet(envName) {
		panic(fmt.Sprintf("env variable %s is not set", envName))
	}

	return viper.GetString(envName)
}

func GetStringMap(envName string) map[string]string {
	if !viper.IsSet(envName) {
		panic(fmt.Sprintf("env variable %s is not set", envName))
	}

	return viper.GetStringMapString(envName)
}

func GetInt(envName string) int {
	if !viper.IsSet(envName) {
		panic(fmt.Sprintf("env variable %s is not set", envName))
	}

	return viper.GetInt(envName)
}

func GetFloat64(envName string) float64 {
	if !viper.IsSet(envName) {
		panic(fmt.Sprintf("env variable %s is not set", envName))
	}

	return viper.GetFloat64(envName)
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
