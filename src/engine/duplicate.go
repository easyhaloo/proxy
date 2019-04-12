package engine

import "log"

// 去重
var visitedUrls = make(map[string]bool)

func IsVisited(url string) bool {
	if visitedUrls[url] {
		log.Printf("抓取到重复的URL: %s", url)
		return true
	}
	visitedUrls[url] = true
	return false
}
