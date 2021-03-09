package main

import "fmt"

func main() {
	// map 的定义
	fmt.Println("------------ map 的定义 ----------")
	m := map[string]string{
		"course": "golang",
		"site":   "imooc",
	}
	m1 := make(map[string]int)
	var m2 map[string]int
	fmt.Println("m=", m)
	fmt.Println("m1=", m1)
	fmt.Println("m2=", m2)

	// map 的遍历
	fmt.Println("------------ map 的遍历 ----------")
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 获取 map 中的值
	fmt.Println("------------ 修改和删除map中的值 ----------")
	fmt.Println("len(m) = ", len(m))
	fmt.Println("m['course'] = ", m["course"])
	// 如果获取的key不存在，就会返回默认值
	fmt.Println("m['name'] = ", m["name"]) // "golang true"
	val, f := m["course"]                  // " false"
	fmt.Println(val, f)
	val, f = m["name"]
	fmt.Println(val, f)
	// 判断 m 中是否有 key = "name" 的元素
	if v, f := m["name"]; f {
		fmt.Println("m['name'] = ", v)
	} else {
		fmt.Printf("The element[%s] dose not exist.\n", "name")
	}

	// 修改和删除 map 中的值
	fmt.Println("------------ 修改和删除map中的值 ----------")
	delete(m, "course")
	fmt.Println("m=", m)

}
