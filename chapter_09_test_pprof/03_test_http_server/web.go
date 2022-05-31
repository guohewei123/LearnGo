package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"testHttpServer/handler"
)

// 定义 自己的error
type userError interface {
	error
	Message() string
}

// 定义 handler 函数类型
type appHandler func(response http.ResponseWriter, request *http.Request) error

// 定义 统一错误处理函数
// 输入: 自定义handler类型
// 输出: http 库要求的 handler类型
func errorHandler(handler appHandler) func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		defer func() { // recover 代码运行的错误
			if r := recover(); r != nil {
				log.Printf("Recover panic error: %v", r)
				http.Error(
					response,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
				return
			}
		}()
		err := handler(response, request)
		if err != nil {
			log.Printf("An error occur: %s", err.Error())

			// user error
			if err, ok := err.(userError); ok { // 处理自己返回的错误
				http.Error(response, err.Message(), http.StatusBadRequest)
				return
			}

			// system error
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
	http.HandleFunc("/", errorHandler(handler.FileHandler))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
