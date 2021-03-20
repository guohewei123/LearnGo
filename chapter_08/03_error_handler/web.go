package main

import (
	"errorHandler/handler"
	"log"
	"net/http"
	"os"
)

// 定义 handler 函数类型
type appHandler func(response http.ResponseWriter, request *http.Request) error

// 定义 统一错误处理函数
// 输入: 自定义handler类型
// 输出: http 库要求的 handler类型
func errorHandler(handler appHandler) func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		err := handler(response, request)
		if err != nil {
			log.Printf("An error occur: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(response, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errorHandler(handler.FileHandler))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
