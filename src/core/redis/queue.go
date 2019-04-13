package redis

import (
	"errors"
	"sync"
)
var (
	lock = &sync.Mutex{}
	redisKeys []string
)


func PushKey(key string){
	redisKeys = append(redisKeys,key)
}


func PopKey() (string,error){
	lock.Lock()
	defer lock.Unlock()
	if len(redisKeys) > 0 {
		key := redisKeys[0]
		redisKeys = redisKeys[1:]
		return key,nil
	}
	return "",errors.New("no redis key can pop");
}