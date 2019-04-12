package config

import (
	"fmt"
	"testing"
)

func TestReaderConfig(t *testing.T) {

	config := InitConfig()
	fmt.Println(config)
}
