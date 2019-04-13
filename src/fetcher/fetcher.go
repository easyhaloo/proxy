package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 限流
var rateLimiter = time.Tick(50 * time.Millisecond)

func Fetcher(url string) ([]byte, error) {
	// 阻塞，当５０毫秒后会想rateLimiter发送数据
	<-rateLimiter
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	loop := 0
	for resp == nil || (resp.StatusCode != http.StatusOK && loop < 10) {
		resp, err = client.Do(request)
		loop++
		if resp != nil {
			fmt.Errorf("抓取页面出错，返回码[%d],尝试重新抓取，重试次数:[%d]\n", resp.StatusCode, loop)
		}
	}

	loop = 0

	if resp.StatusCode != http.StatusOK {
		client = proxyClient()
		if client == nil {
			return nil, fmt.Errorf("无法获取代理客户端\n")
		} else {
			for resp.StatusCode != http.StatusOK && loop < 10 {
				log.Printf("使用代理客户端进行请求,请求地址：%v", url)
				resp, err = client.Do(request)
				loop++
				for resp == nil {
					client = proxyClient()
					resp, err = client.Do(request)
				}
				//log.Printf("使用代理客户端进行请求,返回状态吗：%v", resp.StatusCode)
				fmt.Errorf("代理客户端进行请求，抓取页面出错，返回码[%d],尝试重新抓取，重试次数:[%d]\n", resp.StatusCode, loop)
			}
		}
		if err != nil || resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("抓取页面出错，返回码[%d],已达到最大尝试次数\n", resp.StatusCode)
		}
	}

	if resp == nil {
		return nil, fmt.Errorf("请求地址为：%s,返回体为空", url)
	}

	//fmt.Printf("抓取页面url:[%s]\n", url)
	e := determineEncoding(resp.Body)
	// 转换类型
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	defer resp.Body.Close()
	return ioutil.ReadAll(utf8Reader)
}

// 自动推断编码类型
func determineEncoding(reader io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(reader).Peek(1024)
	// 不足１０２４时，发生异常
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
