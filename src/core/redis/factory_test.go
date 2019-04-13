package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestNewConnection(t *testing.T) {

	factory := RedisFactory{}
	factory.Initialization()

	for i:= 0 ; i < 100 ;i ++ {
		connection := factory.NewConnection()
		go func(client *redis.Client) {
			pong, err := connection.Ping().Result()
			if err != nil {
				panic(err)
			}
			fmt.Println(pong)
			factory.Backup(connection)
		}(connection)

		//time.Sleep(time.Duration(2)*time.Second)
		//factory.Backup(connection)
	}
	time.Sleep(time.Duration(20)*time.Second)

}
