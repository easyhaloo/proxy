package fetcher

import (
	"net/http"
	"net/url"
	"nick.com/proxy/src/core/redis"
	"time"
)

func proxyClient() *http.Client {
	proxyAddr, err := redis.PopKey()
	for err != nil {
		if err != nil {
			proxyAddr, err = redis.PopKey()
			continue
		}
		proxy, err := url.Parse(proxyAddr)
		if err != nil {
			proxyAddr, err = redis.PopKey()
			continue
		}
		netTransport := &http.Transport{
			Proxy:                 http.ProxyURL(proxy),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * time.Duration(5),
		}

		httpClient := &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		}
		//log.Printf("循环获取值,%v",netTransport)
		return httpClient
	}
	proxy, err := url.Parse(proxyAddr)
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	if err != nil {
		return nil
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	return httpClient
}
