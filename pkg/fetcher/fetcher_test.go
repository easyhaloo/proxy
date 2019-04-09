package fetcher

import (
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFetcher(t *testing.T) {
	url := "https://www.kuaidaili.com/free/inha/"
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	i := 0
	for i < 10 {
		resp, err := client.Do(request)
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}
		fmt.Printf("抓取页状态码:[%d]\n", resp.StatusCode)
		e := determineEncoding(resp.Body)
		// 转换类型
		utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
		bytes, _ := ioutil.ReadAll(utf8Reader)
		fmt.Printf("read %v\n", len(bytes))
		i++
	}

}
