package main

import "fmt"

func main() {
	// map 的定义
	fmt.Println("------------ map 的定义 ----------")
	m := map[string]string{
		"course": "golang",
		"site":   "imooc",
	}
	m["123"] = "hao"
	m1 := make(map[string]int)
	var m2 map[string]int
	fmt.Println("m=", m)
	fmt.Println("m1=", m1)
	fmt.Println("m2=", m2)

	// 定义 map 中嵌套 map
	mm := map[string]map[string]string{
		"name": {
			"first":  "Wang",
			"second": "er"},
		"age": {"last_year": "18"}}
	for k, v := range mm {
		fmt.Println(k, v)
	}

	// map 的遍历
	fmt.Println("------------ map 的遍历 ----------")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("只遍历 map 中的 key")
	for k := range m { // 只遍历 map 中的 key
		fmt.Println(k)
	}
	fmt.Println("只遍历 map 中的 val")
	for _, v := range m { // 只遍历 map 中的 val
		fmt.Println(v)
	}

	// 获取 map 中的值
	fmt.Println("------------ 修改和删除map中的值 ----------")
	fmt.Println("len(m) = ", len(m))
	fmt.Println("m['course'] = ", m["course"])
	// 如果获取的key不存在，就会返回默认值
	fmt.Println("m['name'] = ", m["name"]) // "golang true"
	val, ok := m["course"]                 // " false"
	fmt.Println(val, ok)
	val, ok = m["name"]
	fmt.Println(val, ok)
	// 判断 m 中是否有 key = "name" 的元素
	if v, ok := m["name"]; ok {
		fmt.Println("m['name'] = ", v)
	} else {
		fmt.Println("The element[name] dose not exist.")
	}

	// 修改和删除 map 中的值
	fmt.Println("------------ 修改和删除map中的值 ----------")
	delete(m, "course")
	fmt.Println("m=", m)

}
