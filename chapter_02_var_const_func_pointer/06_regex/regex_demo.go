package main

import (
	"fmt"
	"regexp"
)

func main() {
	testStr := `this is a email ggg@163.com
		email1 is qqq@qq.com
		email2 is test2@test.com.cn
		email3 is test3@gmail.com.top
		`
	//re := regexp.MustCompile(`ggg@163\.com`)
	//findString := re.FindString(testStr)

	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[.a-zA-Z0-9]+)`)
	//allString := re.FindAllString(testStr, -1)
	submatch := re.FindAllStringSubmatch(testStr, -1)
	for _, item := range submatch {
		fmt.Println(item)
	}
}
