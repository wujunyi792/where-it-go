package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var configPath = "./config/"
var configName = "config"
var configSuffix = ".json"
var serveConfig *GlobalConfig

func init() {
	fmt.Println("start init conf")
	if !exist(configPath + configName + configSuffix) {
		data, err := json.Marshal(&GlobalConfig{
			MODE: "debug",
		})
		var str bytes.Buffer
		err = json.Indent(&str, data, "", "  ")
		err = ioutil.WriteFile(configPath+configName+configSuffix, str.Bytes(), 0644)
		if err != nil {
			panic("Try to generate config.json fileHandle failed!")
		}
		fmt.Println("config.json generate success")
		os.Exit(0)
	}
	serveConfig = new(GlobalConfig)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic("Config Read failed: " + err.Error())
	}
	err = viper.Unmarshal(serveConfig)
	if err != nil {
		panic("Config Unmarshal failed: " + err.Error())
	}
	Fresh()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config fileHandle changed: ", e.Name)
		_ = viper.ReadInConfig()
		err = viper.Unmarshal(serveConfig)
		if err != nil {
			fmt.Println("New Config fileHandle Parse Failed: ", e.Name)
			return
		}
	})
}

func exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetConfig() *GlobalConfig {
	return serveConfig
}

func Fresh() {
	data, _ := json.Marshal(serveConfig)
	var str bytes.Buffer
	_ = json.Indent(&str, data, "", "  ")
	_ = ioutil.WriteFile(configPath+configName+configSuffix, str.Bytes(), 0644)
}
