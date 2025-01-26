package interfaces

import "fmt"

type Stringer interface {
	String() string
}

func testString(v interface{}) string {

	switch str := v.(type) { // str 是 v 的值
	case string: // 当符合条件时，str 的值转换成为 string 类型
		return str + " 222"
	case Stringer:
		return str.String()
	}

	return "default"
}

// default
// 111 222
func test() {
	i := 111
	s := testString(i)
	fmt.Println(s)

	i1 := "111"
	s = testString(i1)
	fmt.Println(s)
}

// 判断单个值的类型
func judgeSingleType() {
	var v interface{}
	v = "string"

	//s := v.(int) // 会报错：panic: interface conversion: interface {} is string, not int
	//fmt.Println(s)

	s2, ok := v.(string)
	if ok {
		fmt.Printf("string value is: %q\n", s2)
	} else {
		fmt.Printf("value is not a string\n")
	}

	if str, ok := v.(string); ok {
		fmt.Printf("string value is: %q\n", str)
	} else if str, ok := v.(Stringer); ok {
		fmt.Printf("Stringer value is: %q\n", str)
	}
}
