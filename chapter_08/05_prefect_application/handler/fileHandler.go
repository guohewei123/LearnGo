package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// 自定义错误
type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix string = "/list/"

// handler 中有错误就会返回，不需要处理
func FileHandler(response http.ResponseWriter, request *http.Request) error {
	if !strings.HasPrefix(request.URL.Path, "/list") {
		return userError("Path must start with " + prefix)  // 返回自定义错误
	}
	path := request.URL.Path[len(prefix):]
	fmt.Println("request url: ", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	response.Write(all)
	return nil
}
