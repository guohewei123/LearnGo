package main

import (
	"fmt"
	"sync"
	"time"
)


func MapDemo1() {
	m := &sync.Map{}
	go func() {
		for {
			m.Store("aa", "bb")
			m.Store(1, 2)
		}
	}()

	go func() {
		for {
			fmt.Println(m.Load("aa"))
			fmt.Println(m.Load(1))

		}
	}()
	fmt.Println("结束")

	time.Sleep(time.Millisecond)
}


func MapDemo2() {
	m := &sync.Map{}
	m.Store("aa", "bb")
	fmt.Println(m.Load("aa"))
	m.Delete("aa")
	fmt.Println(m.Load("aa"))
}


func MapDemo3() {
	m := &sync.Map{}
	m.LoadOrStore("cc", "dd")
	fmt.Println(m.LoadAndDelete("cc"))
	fmt.Println(m.Load("cc"))
}

func MapDemo4() {
	m := &sync.Map{}
	m.Store("aa", "bb")
	m.Store("11", "22")
	m.Store("cc", "dd")
	m.Store(1, 2)

	fmt.Println("------------- Range return true -------------")
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		time.Sleep(time.Second)
		return true
	})

	fmt.Println("------------- Range return false -------------")
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		time.Sleep(time.Second)
		return false
	})
}

func main() {
	//MapDemo1()
	//MapDemo2()
	//MapDemo3()
	MapDemo4()
}
