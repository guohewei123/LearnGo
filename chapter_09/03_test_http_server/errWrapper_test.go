package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// 1. 定义 测试 panic 异常 Handle
func PanicHandle(response http.ResponseWriter, request *http.Request) error {
	panic(123)
}

// 2. 定义 测试自定义错误 Handle
type testUserError string

func (e testUserError) Error() string {
	return e.Message()
}

func (e testUserError) Message() string {
	return string(e)
}

func UserErrorHandle(response http.ResponseWriter, request *http.Request) error {
	return testUserError("user custom define error")
}

// 3. 定义 测试 NotExist 异常 Handle
func NotExistHandle(response http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

// 4. 定义 测试 Permission 异常 Handle
func PermissionHandle(response http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

// 5. 定义 测试 Others 异常 Handle
func OthersErrHandle(response http.ResponseWriter, request *http.Request) error {
	return errors.New("others error")
}

// 6. 定义 测试没有异常 Handle
func NoErrorHandle(response http.ResponseWriter, request *http.Request) error {
	//fmt.Fprint(response, "no error")
	response.Write([]byte("no error"))
	return nil
}

// 测试集
var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{PanicHandle, 500, "Internal Server Error"},
	{UserErrorHandle, 400, "user custom define error"},
	{NotExistHandle, 404, "Not Found"},
	{PermissionHandle, 403, "Forbidden"},
	{OthersErrHandle, 500, "Internal Server Error"},
	{NoErrorHandle, 200, "no error"},
}

func verifyResponse(response *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)                                   // 获取返回body
	body := strings.Trim(string(b), "\n")                                   // 将获取的body 转换成 string 类型
	if body != expectedMsg || response.StatusCode != expectedCode {
		t.Errorf("expected (%s, %d), got (%s, %d)", expectedMsg, expectedCode, body, response.StatusCode)
	}
}

// 只测试 ErrWrapper 不启动服务
func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errorHandler(tt.h)                                                 // 调用errorHandler 获取返回函数
		response := httptest.NewRecorder()                                      // 手动创建测试 response
		request := httptest.NewRequest(http.MethodGet, "http://baidu.com", nil) // 手动创建 request
		f(response, request)                                                    // 调用返回函数
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

// 启动服务测试 ErrWrapper
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errorHandler(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

