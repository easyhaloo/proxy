package src

import (
	eight "nick.com/proxy/src/core/parser/89proxy/parse"
	six "nick.com/proxy/src/core/parser/sixsixproxy/parse"
	xici "nick.com/proxy/src/core/parser/xici/parse"
	"nick.com/proxy/src/engine"
	"nick.com/proxy/src/persist"
	"nick.com/proxy/src/schduler"
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
