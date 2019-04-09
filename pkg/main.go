package main

import (
	eight "nick.com/proxy/pkg/89proxy/parse"
	"nick.com/proxy/pkg/engine"
	"nick.com/proxy/pkg/persist"
	"nick.com/proxy/pkg/schduler"
	six "nick.com/proxy/pkg/sixsixproxy/parse"
	xici "nick.com/proxy/pkg/xici/parse"
)

func main() {

	saver, _ := persist.ItemSaver("xxx")
	seed := []engine.Request{
		{
			Url:   "https://www.xicidaili.com/nn",
			Parse: engine.NewFuncParse(xici.ProxyHostParse, "proxy"),
		},
		{
			Url:   "http://www.66ip.cn/",
			Parse: engine.NewFuncParse(six.ProxyHostParse, "proxy"),
		},
		{
			Url:   "http://www.89ip.cn/",
			Parse: engine.NewFuncParse(eight.ProxyHostParse, "proxy"),
		},
	}

	e := engine.ConcurrentEngine{
		Scheduler:      &schduler.QueueScheduler{},
		MaxWorkerCount: 100,
		ItemChan:       saver,
		RequestWorker:  engine.Worker,
	}

	e.Run(seed...)
}
