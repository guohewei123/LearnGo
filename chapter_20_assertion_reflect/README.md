## 20-1 断言assertion使用
```go
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
```
输出
```shell script
Cat name: 凯特 
v.(Dog).name= 大卫
v.(Animal).classify= 爬行动物
```

## 20-2 反射reflect使用
通过反射可以知道数据的原始类型，数据内容，方法等，并可以进行一定的操作
### 常用操作：
- reflect.ValueOf(in)                    获取输入参数接口中的数据的值
- reflect.TypeOf(in)                     获取输入参数接口中的值的类型
- reflect.TypeOf(in).NumField()          获取输入参数接口中的值的个数
- reflect.ValueOf(in).Field(int)         用来获取结构体第几个属性的值
- reflect.ValueOf(in).FieldByName("key") 用来获取结构体key属性的值
- reflect.ValueOf(in)..FieldByIndex([]int{0,1}) 按照层级取值 xxx([]int{0,1}) 第0元素内部的第1个元素； xxx([]int{1,2}) 第1个元素内部的第2个元素
- reflect.TypeOf(in).Kind()                用来判断类型
    ```go
    tk := reflect.TypeOf(input).Kind()
    if tk == reflect.Struct{
        fmt.Println("input is a Struct")
    }
    if tk == reflect.String{
        fmt.Println("input is a String")
    }
    if tk == reflect.Int{
        fmt.Println("input is a Int")
    }
    ```
- reflect.ValueOf(in).Elem()          获取原始数据并操作
  - 修改字段的值 `reflect.ValueOf(in).Elem().FieldByName("Name").SetString("凯特媳妇")`
  - **注意：1. 入参必须是地址  2. 修改的属性数Public类型（首字母大写）**
    ```go
    v := reflect.ValueOf(input) 
    e := v.Elem()
    e.FiledByName("key").SetString("修改后")
    fmt.Println(input)
    ```
- reflect.ValueOf(in).Method(0)           获取input第0个方法
- reflect.ValueOf(in).NumMethod()         获取input的方法总数
- reflect.ValueOf(in).MethodByName("Key") 获取input的方法名为Key的方法
    ```go
    v := reflect.ValueOf(input) 
    m := v.Method(0)
    m.Call([]reflect.Value{reflect.ValueOf("入参测试")})
    ```

### 使用示例：
- 普通反射
- struct反射
- 匿名或切入字段反射
- 判断传入的类型是否是我们想要的类型
- 通过反射修改内容
- 通过反射调用方法

### 获取golang中的函数的名字
- 如何获取golang中的函数的名字, 这里需要用到反射. 可以看代码: "chapter_20/03_reflect_FuncForPC/main.go".
- 重点使用 runtime.FuncForPC 这个函数获取函数名.
- 使用strings.FieldsFunc 对得到的带 路径名和包名的 函数名进行必要的处理.

### 优秀博文
- Go reflection 三定律与最佳实践: https://halfrost.com/go_reflection/