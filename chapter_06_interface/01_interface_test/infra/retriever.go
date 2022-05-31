package infra

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
