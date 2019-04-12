package persist

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"nick.com/proxy/src/engine"
	"time"
)

func ItemSaver(index string) (chan engine.Proxy, error) {
	ch := make(chan engine.Proxy, 1024)
	go func() {
		itemCount := 0
		for item := range ch {
			itemCount++
			err := Save(item)
			log.Printf("Item Saver : Save error : %s,count:%v", item, itemCount)
			if err != nil {
				log.Printf("Item Saver : Save error : %s", err)
			}
		}
	}()
	return ch, nil
}

func Save(item engine.Proxy) error {
	log.Printf("Item Saver : Save error : %s.", item)
	if item.Port == "" {
		return errors.New("item not be arrowed null")
	}
	if proxyCheck(item) {
		log.Printf("proxy url [%s] is valid.", item)
	}
	return nil
}

const targetUrl = "http://icanhazip.com/"

func proxyCheck(proxy engine.Proxy) bool {
	proxyUrl := "http://" + proxy.Ip + ":" + proxy.Port
	urlProxy, err := url.Parse(proxyUrl)
	if err != nil {
		return false
	}
	c := http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyURL(urlProxy),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * time.Duration(5),
		},
		Timeout: time.Second * 10,
	}
	if resp, err := c.Get(targetUrl); err != nil {
		log.Printf("proxy url is invalid , the url :%s", proxyUrl)
		return false
	} else if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		return true
	} else {
		defer resp.Body.Close()
		return false
	}

}
