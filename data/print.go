package data

import (
	"fmt"
	"os"
)

// 打印初识
// Hello 23
// Hello 23
// Hello 23
// Hello 23
func print() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
}

// 打印 string 用不同的 v% 占位符形式
/*
%v 可打印任意值，包括复合类型(数组、切片、结构体、map)
%+v 显示结构体字段名
%#v 输出 Go 语法格式
%T 打印值的类型
*/
// i: Hello
// i show filed: Hello
// i go: "Hello"
// i type is: string
func printV() {
	i := "Hello"
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i show filed: %+v\n", i)
	fmt.Printf("i go: %#v\n", i)
	fmt.Printf("i type is: %T\n", i)
}

// 打印 struct
// u: {you 15}
// u show filed: {name:you age:15}
// u go: main.User{name:"you", age:15}
// u type is: main.User
func printVForStruct() {
	u := User{
		name: "you",
		age:  15,
	}

	fmt.Printf("u: %v\n", u)
	fmt.Printf("u show filed: %+v\n", u)
	fmt.Printf("u go: %#v\n", u)
	fmt.Printf("u type is: %T\n", u)
}

type User struct {
	name string
	age  int
}

func (u User) String() string {
	return fmt.Sprintf("User name: %s, age: %d", u.name, u.age)
}
