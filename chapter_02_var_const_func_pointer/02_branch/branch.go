package main

import (
	"fmt"
	"io/ioutil"
)

// if 语句1
func readFile1() {
	const filename = "learnGo/chapter_02/02_branch/abc.txt" // abc.txt = aaaa\nbbbbb\ncccc
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n", contents)
	}
}

// if 语句1
func readFile2() {
	const filename = "learnGo/chapter_02/02_branch/abc.txt" // abc.txt = aaaa\nbbbbb\ncccc
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}
	//fmt.Println("%s\n", contents)   # contents 是if中定义的变量，出了if后，变量的生命周期结束了
}

// switch 语句
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
		fallthrough  // fallthrough 就不会自动break了
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	readFile1()
	fmt.Println()
	readFile2()

	fmt.Println(grade(30))
	fmt.Println(grade(60))
	fmt.Println(grade(70))
	fmt.Println(grade(80))
	fmt.Println(grade(90))
	fmt.Println(grade(100))
	//fmt.Println(grade(-1))
	//fmt.Println(grade(101))

}
