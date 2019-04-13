package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const fileName = "/home/haloo/go/src/nick.com/proxy/config/config.yml"

type Config struct {
	// redis 配置信息
	Redis redis
}

// redis 连接信息结构体
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"database"`
	Count    int    `yaml:count`
}

func GetRedisInfo() redis {
	InitConfig()
	return config.Redis
}

var config *Config

func InitConfig() *Config {
	if config != nil {
		log.Printf("config already initialized")
		return config
	}
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
