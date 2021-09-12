package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	classify string
}

type Cat struct {
	Animal
	Age  uint
	Name string
}

func (c *Cat) Print(like string) {
	fmt.Printf("Cat like: %s, Name: %s \n", like, c.Name)
}

/* ValueOf() 和 TypeOf() 到底输出什么
animal := Animal{classify: "哺乳动物"}
cat := Cat{
	Animal: animal,
	Age:    1,
	Name:   "凯特",
}
reflectValueOf(cat)
输出：
{{哺乳动物} 1 凯特}
main.Cat
*/
func reflectValueOfTypeOf(in interface{}) {
	v := reflect.ValueOf(in)
	t := reflect.TypeOf(in)
	fmt.Println(v)
	fmt.Println(t)
}

/* NumField() 和 Field(int) 的使用
animal := Animal{classify: "哺乳动物"}
cat := Cat{
	Animal: animal,
	Age:    1,
	Name:   "凯特",
}
reflectNumFieldField(cat)
输出：
{哺乳动物}
1
凯特
*/
func reflectNumFieldField(in interface{}) {
	v := reflect.ValueOf(in)
	t := reflect.TypeOf(in)
	n := t.NumField()
	for i := 0; i < n; i++ {
		fmt.Println(v.Field(i))
	}
}

/* FieldByName("key") 和 FieldByIndex([]int{0,1}) 的使用
animal := Animal{classify: "哺乳动物"}
cat := Cat{
	Animal: animal,
	Age:    1,
	Name:   "凯特",
}
reflectFieldByNameAndFieldByIndex(cat)
输出：
input:  {{哺乳动物} 1 凯特}
第1个元素 {哺乳动物}
第2个元素 1
第3个元素 凯特
第0个元素内部的第0个元素 哺乳动物
================================
1
凯特
{哺乳动物}
*/
func reflectFieldByNameAndFieldByIndex(in interface{}) {
	v := reflect.ValueOf(in)
	//t := reflect.TypeOf(in)
	fmt.Println("input: ", in)                                // 输出 {{哺乳动物} 1 凯特}
	fmt.Println("第1个元素", v.FieldByIndex([]int{0}))            // 第1个元素 {哺乳动物}
	fmt.Println("第2个元素", v.FieldByIndex([]int{1}))            // 第2个元素 1
	fmt.Println("第3个元素", v.FieldByIndex([]int{2}))            // 第3个元素 凯特
	fmt.Println("第0个元素内部的第0个元素", v.FieldByIndex([]int{0, 0})) // 第0个元素内部的第0个元素 哺乳动物
	fmt.Println("================================")
	fmt.Println(v.FieldByName("Age"))
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Animal"))

}

/* Kind() 的使用
animal := Animal{classify: "哺乳动物"}
cat := Cat{
	Animal: animal,
	Age:    1,
	Name:   "凯特",
}
reflectKind(cat)
输出：
input is a Struct
*/
func reflectKind(in interface{}) {
	t := reflect.TypeOf(in)
	tk := t.Kind()
	if tk == reflect.Struct {
		fmt.Println("input is a Struct")
	}
	if tk == reflect.String {
		fmt.Println("input is a String")
	}
	if tk == reflect.Int {
		fmt.Println("input is a Int")
	}
}

/* Elem() 的使用
// 注意：1. 入参必须是地址  2. 修改的属性数Public类型（首字母大写）
animal := Animal{classify: "哺乳动物"}
cat := Cat{
	Animal: animal,
	Age:    1,
	Name:   "凯特",
}
reflectElem(&cat)
输出：
修改前 &{{哺乳动物} 1 凯特}
e.FieldByName("Name") =  凯特
修改后 &{{哺乳动物} 1 凯特媳妇}
*/
func reflectElem(in interface{}) {
	fmt.Println("修改前", in)
	v := reflect.ValueOf(in)
	e := v.Elem()
	fmt.Println(`e.FieldByName("Name") = `, e.FieldByName("Name"))
	e.FieldByName("Name").SetString("凯特媳妇")
	fmt.Println("修改后", in)
}

/* Method() 的使用
// 注意：1. 入参必须是地址  2. 修改的属性数Public类型（首字母大写）
animal := Animal{classify: "哺乳动物"}
cat := Cat{
	Animal: animal,
	Age:    1,
	Name:   "凯特",
}
cat.Print("大米")
reflectMethod(&cat)
输出：
Cat like: 大米, Name: 凯特
v.MethodByName("Print") =  0x10ad8a0
Cat like: 小米, Name: 凯特
v.Method(0) =  0x10ad8a0
Cat like: 中米, Name: 凯特
*/
func reflectMethod(in interface{}) {
	v := reflect.ValueOf(in)
	m := v.MethodByName("Print")
	fmt.Println(`v.MethodByName("Print") = `, v.MethodByName("Print"))
	m.Call([]reflect.Value{reflect.ValueOf("小米")})

	m1 := v.Method(0)
	fmt.Println(`v.Method(0) = `, v.Method(0))
	m1.Call([]reflect.Value{reflect.ValueOf("中米")})
}

func main() {

	animal := Animal{classify: "哺乳动物"}
	cat := Cat{
		Age:    1,
		Name:   "凯特",
		Animal: animal,
	}
	cat.Print("大米")
	reflectMethod(&cat)
}
