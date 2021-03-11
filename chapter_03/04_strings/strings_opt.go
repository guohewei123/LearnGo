package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("------- 字符串测试 -----")
	fmt.Println("字符串转化")
	//获取程序运行的操作系统平台下 int 类型所占的位数，如：strconv.IntSize。
	fmt.Println(strconv.IntSize) // 64

	fmt.Println("------------------ 1. 将字符串转换为 int 型 -----------------------")
	var s string = "100"
	res, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res) // 100
	}

	fmt.Println("------------------ 2. 将字符串转换为 float64 型 -----------------------")
	var str01 string = "100.55"
	float01, err := strconv.ParseFloat(str01, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(float01) // 100.55
	}
	str02 := strconv.Itoa(99)           // "99"
	fmt.Printf("int 转字符: %s \n", str02) // int 转字符: 99

	fmt.Println("------------------ 3. 字符串比较 -----------------------")
	//com01 := strings.Compare(str01, str02)
	com01 := strings.Compare(str02, str02)
	if com01 == 0 {
		fmt.Println("相等", com01) // 相等 0
	} else {
		fmt.Println("不相等 ", com01) // 不相等  -1

	}

	fmt.Println("------------------ 4. 包含 -----------------------")
	str01 = "hello world"
	isCon := strings.Contains(str01, "hello")
	fmt.Println(isCon) // true

	fmt.Println("------------------ 5. 查找位置 -----------------------")
	str01 = "hello, world"
	theIndex := strings.Index(str01, ",")
	fmt.Println(theIndex)                     // 5
	fmt.Println(strings.Index(str01, "haha")) //不存在返回 -1
	lastIndex := strings.LastIndex(str01, "o")
	fmt.Println("在字符串中最后出现位置的索引 ", strconv.Itoa(lastIndex)) // 在字符串中最后出现位置的索引  8

	fmt.Println("------------------ 6. 统计给定子串sep的出现次数 -----------------------")
	//统计给定子串sep的出现次数, sep为空时, 返回: 1 + 字符串的长度
	fmt.Println(strings.Count("cheeseeee", "ee")) // 3
	fmt.Println(strings.Count("five", ""))        // 5  sep为空时, 返回 len("five") + 1

	fmt.Println("------------------ 7. 重复s字符串count次 -----------------------")
	// 重复s字符串count次, 最后返回新生成的重复的字符串
	fmt.Println("hello " + strings.Repeat("world ", 5)) // hello world world world world world

	fmt.Println("------------------ 8. 在s字符串中, 把old字符串替换为new字符串 -----------------------")
	// 在s字符串中, 把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	var str03 string = "/Users//Documents/GOPatch/src/MyGO/config/TestString/"
	str04 := strings.Replace(str03, "/", "**", -1)
	str05 := strings.Replace(str03, "/", "**", 4)
	fmt.Println(str04) // **Users****Documents**GOPatch**src**MyGO**config**TestString**
	fmt.Println(str05) // **Users****Documents**GOPatch/src/MyGO/config/TestString/

	fmt.Println("删除字符串的开头和尾部")
	fmt.Println("------------------ 9. 删除字符串的开头和尾部 -----------------------")
	str03 = "/Users/Documents/GOPatch/src/TestString/"
	fmt.Println("删除两头的 / = " + strings.Trim(str03, "/"))       // 删除两头的 / = Users/Documents/GOPatch/src/TestString
	fmt.Println("删除左边的 / = " + strings.TrimLeft(str03, "/"))   // 删除左边的 / =  Users/Documents/GOPatch/src/TestString/
	fmt.Println("删除右边边的 / = " + strings.TrimRight(str03, "/")) // 删除右边边的 / = /Users/Documents/GOPatch/src/TestString
	str06 := strings.TrimSpace(" hello hao hao hao ")
	fmt.Printf("删除开头末尾的空格:%s已删除\n", str06) // 删除开头末尾的空格:hello hao hao hao已删除

	fmt.Println("------------------ 10. 大小写 -----------------------")
	fmt.Println(strings.Title("hello hao hao hao"))   // Hello Hao Hao Hao
	fmt.Println(strings.ToLower("Hello Hao Hao Hao")) // hello hao hao hao
	fmt.Println(strings.ToUpper("hello hao hao hao")) // HELLO HAO HAO HAO

	fmt.Println("------------------ 11. 前缀 后缀 -----------------------")
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasSuffix("Amigo", "go"))  // true

	fmt.Println("------------------ 12. 字符串分割 -----------------------")
	fieldsStr := "  hello   it's  a  nice day today    "
	//根据空白符分割,不限定中间间隔几个空白符
	fieldsSlice := strings.Fields(fieldsStr)
	fmt.Println(fieldsSlice) // [hello it's a nice day today]

	for i, v := range fieldsSlice {
		fmt.Printf("下标 %d 对应值 = %s \n", i, v)
	}
	/*下标 0 对应值 = hello
	下标 1 对应值 = it's
	下标 2 对应值 = a
	下标 3 对应值 = nice
	下标 4 对应值 = day
	下标 5 对应值 = today*/
	for i := 0; i < len(fieldsSlice); i++ {
		fmt.Printf("%s ", fieldsSlice[i]) // hello it's a nice day today
	}
	fmt.Println()

	//根据特定字符分割
	slice01 := strings.Split("q,w,e,r,t,y,", ",") // slice01 = [q w e r t y ]
	fmt.Println(slice01)                          // [q w e r t y ]
	fmt.Println(cap(slice01))                     //7  最后多个空""
	for i, v := range slice01 {
		fmt.Printf("下标 %d 对应值 = %s \n", i, v)
	}
	/*下标 0 对应值 = q
	下标 1 对应值 = w
	下标 2 对应值 = e
	下标 3 对应值 = r
	下标 4 对应值 = t
	下标 5 对应值 = y
	下标 6 对应值 =  */

	fmt.Println("------------------ 13. 字符串拼接 -----------------------")
	//Join 用于将元素类型为 string 的 slice, 使用分割符号来拼接组成一个字符串：
	fieldsSlice = []string{"hello", "it's", "a", "nice", "day", "today"}
	var str08 string = strings.Join(fieldsSlice, ",")
	fmt.Println("Join拼接结果=" + str08) // Join拼接结果=hello,it's,a,nice,day,today

	fmt.Println("------------ 14. 对比字符串拼接效率----------------")

	// buffer 拼接
	var buffer bytes.Buffer
	start := time.Now()
	for i := 0; i < 100000; i++ {
		buffer.WriteString("test is here")
	}
	buffer.String() // 输出拼接结果
	end := time.Now()
	fmt.Println("Buffer time is: ", end.Sub(start).Seconds()) // Buffer time is:  0.003907722

	// += 拼接
	start = time.Now()
	str := ""
	for i := 0; i < 100000; i++ {
		str += "test is here"
	}
	end = time.Now()
	fmt.Println("'+=' time is: ", end.Sub(start).Seconds()) // '+=' time is:  4.642418455

	// Join 拼接
	start = time.Now()
	var sl []string
	for i := 0; i < 100000; i++ {
		sl = append(sl, "test is here")
	}
	strings.Join(sl, "")
	end = time.Now()
	fmt.Println("Join time is: ", end.Sub(start).Seconds()) // Join time is:  0.015088876
	/*
		Buffer time is:  0.003907722
		'+=' time is:  4.642418455
		Join time is:  0.015088876
	*/
}
