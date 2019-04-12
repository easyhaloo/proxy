package engine

import (
	"log"
	"nick.com/proxy/src/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("请求[%s]失败：%s", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parse.Parse(body, r.Url), nil
}
