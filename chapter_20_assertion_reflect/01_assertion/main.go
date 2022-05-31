package main

import (
	"fmt"
)

type Animal struct {
	classify string
}

type Cat struct {
	Animal
	age  uint
	name string
}

type Dog struct {
	Animal
	age  uint
	name string
}

func (c *Cat) print() {
	fmt.Printf("Cat name: %s \n", c.name)
}

func checkInputEg(v interface{}) {
	switch v.(type) {
	case Animal:
		fmt.Println("v.(Animal).classify=", v.(Animal).classify)
	case Cat:
		cat := v.(Cat)
		cat.print()
	case Dog:
		fmt.Println("v.(Dog).name=", v.(Dog).name)
	}
}

func main() {
	cat := Cat{
		Animal: Animal{classify: "哺乳动物"},
		age:    1,
		name:   "凯特",
	}
	checkInputEg(cat)

	dog := Dog{
		Animal: Animal{},
		age:    2,
		name:   "大卫",
	}
	checkInputEg(dog)

	base := Animal{classify: "爬行动物"}
	checkInputEg(base)
}
