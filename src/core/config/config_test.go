package config

import (
	"fmt"
	"testing"
)

func TestReaderConfig(t *testing.T) {

	config := InitConfig()
	info := GetRedisInfo()
	fmt.Println(config,info)
}
