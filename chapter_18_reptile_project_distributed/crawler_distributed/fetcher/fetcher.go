package fetcher

import (
	"bufio"
	"crawler_distributied/config"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 获取数据 出入url 输出结果
var rateLimiter = time.Tick(time.Second / config.Qps)
func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	log.Printf("Fetching url %s", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// 设置header
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted) {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)         // 获取html的编码方式
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())  // 按照html的编码方式 解析html为 utf-8 编码
	return ioutil.ReadAll(utf8Reader)
}

// 获取html的编码方式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err :=r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}