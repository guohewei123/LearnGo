package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "You弄啥了!"
	fmt.Println(s, "len=", len(s))   // 输出: You弄啥了! len= 13
	fmt.Println(s, "Rune count=", utf8.RuneCountInString(s))  // 输出: You弄啥了! Rune count= 7
	fmt.Println("--------------分割线----------------")
	fmt.Println("[]byte(s) = ", []byte(s))
	fmt.Println("[]rune(s) = ", []rune(s))
	fmt.Println()
	fmt.Println("--------------分割线----------------")
	for i, ch := range s {      // ch is a rune
		//fmt.Printf("(%d, %d, %c) ", i, ch, ch)
		fmt.Printf("(%d, %x) ", i, ch)          // ((0, 59) (1, 6f) (2, 75) (3, 5f04) (6, 5565) (9, 4e86) (12, 21)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		//fmt.Printf("(%d, %x, %c) ", i, ch, ch)
		fmt.Printf("(%d, %x) ", i, ch)          // (0, 59) (1, 6f) (2, 75) (3, 5f04) (4, 5565) (5, 4e86) (6, 21)
	}
	fmt.Println()
	fmt.Println("--------------分割线----------------")
	s = "You弄啥了!"
	myBytes := []byte(s)
	for len(myBytes) > 0 {
		ch, size := utf8.DecodeRune(myBytes)
		myBytes = myBytes[size:]
		fmt.Printf("%c ", ch)   // Y o u 弄 啥 了 !
	}
}
