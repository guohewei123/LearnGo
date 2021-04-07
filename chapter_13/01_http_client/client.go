package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func httpClientGet() {
	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}


func httpClientGetRedirect() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	//resp, err := http.DefaultClient.Do(request)
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect req to: ", req.URL)  // 存放 upcoming request 即将到来的请求
			for i, v := range via {
				fmt.Printf("Via request [%d] to [%s]: \n", i, v.URL)  // via 存放 requests made already 已经发出的请求列表
			}
			return nil  // 返回 nil 表示允许重定向

	}}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Printf("%s\n", s)
}


func main() {
	//httpClientGet()
	httpClientGetRedirect()
}
