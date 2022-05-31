package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// handler 中有错误就会返回，不需要处理
func FileHandler(response http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
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
