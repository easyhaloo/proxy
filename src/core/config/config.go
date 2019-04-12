package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const fileName = "../../config/config.yml"

type Config struct {
	// redis 配置信息
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		DB       int    `yaml:"database"`
	}
}

func InitConfig() Config {
	config := Config{}
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		log.Printf("读取配置文件出错．．．\n")
		panic(err)
	} else {
		err = yaml.Unmarshal(contents, &config)
		if err != nil {
			log.Printf("配置文件解析出错．．．\n")
			panic(err)
		}
	}
	return config
}
