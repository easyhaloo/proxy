package main

import (
	"nick.com/proxy/src/core/parser/89proxy"
	"nick.com/proxy/src/core/parser/qyproxy"
	"nick.com/proxy/src/core/parser/sixsixproxy"
	"nick.com/proxy/src/core/parser/xici"
	"nick.com/proxy/src/core/redis"
	"nick.com/proxy/src/engine"
	"nick.com/proxy/src/persist"
	"nick.com/proxy/src/schduler"
)

func main() {

	seed := []engine.Request{
		{
			Url:   "https://www.xicidaili.com/nn",
			Parse: engine.NewFuncParse(xici.ProxyHostParse, "proxy"),
		},
		{
			Url:   "http://www.66ip.cn/",
			Parse: engine.NewFuncParse(sixsixproxy.ProxyHostParse, "proxy"),
		},
		{
			Url:   "http://www.89ip.cn/",
			Parse: engine.NewFuncParse(_9proxy.ProxyHostParse, "proxy"),
		},
		{
			Url:   "http://www.qydaili.com/free/",
			Parse: engine.NewFuncParse(qyproxy.ProxyHostParse, "proxy"),
		},
	}
	factory := redis.RedisFactory{}
	factory.Initialization()
	saver, _ := persist.ItemSaver("xxx", factory)
	e := engine.ConcurrentEngine{
		Scheduler:      &schduler.QueueScheduler{},
		MaxWorkerCount: 100,
		ItemChan:       saver,
		RequestWorker:  engine.Worker,
	}

	e.Run(seed...)
}
