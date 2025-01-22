package controlstructures

import "fmt"

// match g
// 把 switch 当做 if else if else 使用
func switchAsIfElseIfElse() {
	i := "g"

	switch {
	case i == "e":
		fmt.Println("match e")
	case i == "t":
		fmt.Println("match t")
	case i == "f":
		fmt.Println("match f")
	case i == "g":
		fmt.Println("match g")
	}
}

// switch 后不必是常量或者整数
// match 111
// match 999
func switchAny() {
	s := "111"
	switch s {
	case "1":
		fmt.Println("match 1")
	case "11":
		fmt.Println("match 11")
	case "111":
		fmt.Println("match 111")
	}

	i := 999
	switch i {
	case 9:
		fmt.Println("match 9")
	case 99:

		fmt.Println("match 99")
	case 999:
		fmt.Println("match 999")
	}
}

// case 后可以跟多个 case
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

// string this is
// integer 99
// pointer to integer 8
func switchType() {
	var t interface{}
	t = "this is"
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case string:
		fmt.Printf("string %s\n", t) // t has type string
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	case *string:
		fmt.Printf("pointer to string %s\n", *t) // t has type *int
	}

	var y any
	y = 99
	switch y := y.(type) {
	default:
		fmt.Printf("unexpected type %T\n", y) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", y) // t has type bool
	case int:
		fmt.Printf("integer %d\n", y) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *y) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *y) // t has type *int
	}

	p := 8
	var i interface{} = &p // 首先将指针赋值给接口类型
	switch ptr := i.(type) {
	default:
		fmt.Printf("unexpected type %T\n", ptr) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", ptr) // t has type bool
	case int:
		fmt.Printf("integer %d\n", ptr) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *ptr) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *ptr) // t has type *int
	}
}
