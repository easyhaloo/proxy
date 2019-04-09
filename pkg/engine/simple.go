package engine

import "log"

type SimpleEngine struct {
}

func (e *SimpleEngine) Run(queue ...Request) {
	var count = 0
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		parseResult, err := Worker(r)
		if err != nil {
			// 抓取错误，记录日志
			log.Printf("crawl is error,the url : %s", r.Url)
			continue
		}

		if IsVisited(parseResult.NextRequest.Url) {
			// 抓取到重复的ＵＲＬ直接跳过
			continue
		}
		queue = append(queue, parseResult.NextRequest)

		for _, item := range parseResult.Proxys {
			count++
			log.Printf("Got proxy :$%d %v", count, item)
		}

	}
}
