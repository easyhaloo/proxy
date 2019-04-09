package parse

import (
	"fmt"
	"nick.com/proxy/pkg/engine"
	"regexp"
)

const (
	site       = "https://www.xicidaili.com/nn/"
	addressReg = `<td>(\w+\.+\w+\.\w+\.+\w+)</td>`
	portReg    = `<td>(\d+)</td>`

	nextReg = `<a class="next_page" rel="next" href="/nn/([\d]+)"[^>]*>([^<]+)</a>`
)

func ProxyHostParse(content []byte, _ string) engine.ParseResult {
	compile := regexp.MustCompile(addressReg)
	matches := compile.FindAllSubmatch(content, -1)

	compile2 := regexp.MustCompile(portReg)
	matches2 := compile2.FindAllSubmatch(content, -1)

	compile3 := regexp.MustCompile(nextReg)

	proxys := make([]engine.Proxy, len(matches))
	for i, v := range matches {
		proxys[i].Ip = string(v[1])
	}
	for i, v := range matches2 {
		proxys[i].Port = string(v[1])
	}

	rs := engine.ParseResult{}
	all := compile3.FindAllSubmatch(content, -1)
	next := ""
	for _, v := range all {
		fmt.Println(string(v[1]))
		next = string(v[1])
	}
	nextUrl := site + next

	// 下一页还用此解析器
	rs.Proxys = proxys
	rs.NextRequest = engine.Request{
		Url:   nextUrl,
		Parse: engine.NewFuncParse(ProxyHostParse, "proxy"),
	}
	return rs
}
