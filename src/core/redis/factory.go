package redis

import (
	"github.com/go-redis/redis"
	"log"
	"nick.com/proxy/src/core/config"
)

type RedisFactory struct {
	connections []*redis.Client
	sign        chan bool
}

// redis 连接工厂
func (r *RedisFactory) NewConnection() *redis.Client {
	var connection *redis.Client
	if len(r.connections) > 0 {
		log.Printf("conrrect requried connection")
		connection = r.connections[0]
		r.connections = r.connections[1:]
		return connection
	} else {
		log.Printf("waiting back connection")
		select {
		case <-r.sign:
			connection = r.connections[0]
			r.connections = r.connections[1:]
			return connection
		}
	}

}

func (r *RedisFactory) Backup(redis *redis.Client) {
	log.Printf("back connection")
	r.connections = append(r.connections, redis)
	r.sign <- true
}

//initialized connection
func (r *RedisFactory) Initialization() {
	// 初始化锁
	r.sign = make(chan bool, 10)
	redisInfo := config.GetRedisInfo()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisInfo.Address,
		Password: redisInfo.Password,
		DB:       redisInfo.DB,
	})
	for i := 0; i < redisInfo.Count; i++ {
		r.connections = append(r.connections, redisClient)
	}
}
