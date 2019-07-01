package config

import (
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	VERSION string `yaml:"VERSION"`
	DB_HOST string `yaml:"DB_HOST"`
	DB_USER string `yaml:"DB_USER"`
	DB_PWD string `yaml:"DB_PWD"`
	DB_NAME string `yaml:"DB_NAME"`
}

func GetConfig() Config {
	// 设置文件路径
	dir, err := os.Getwd()
	filePath := dir + "/config.yaml"
	// 读取并解析文件
	buffer, err := ioutil.ReadFile(filePath)
	config := Config{}
	err = yaml.Unmarshal(buffer, &config)
  if err != nil {
		fmt.Println(err.Error())
	}
	return config
}

