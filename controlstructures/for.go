package controlstructures

import (
	"fmt"
)

/**
// Like a C for - for 循环基本结构
for init; condition; post { }

// Like a C while - 相当于 while 结构
for condition { }

// Like a C for(;;) - 无限循环结构
for { }

*/

// 简单 for 循环
func forSimple() {
	for i := 0; i < 10; i++ {
	}
}

// while
func likeWhile() {
	// while 表达式循环
	i := 1 > 0
	for i {
	}

	// 无限循环
	for {

	}
}

// 可以遍历：Array Slice String Map Channel

// 遍历 Array
// key: 0, value: 1
// key: 1, value: 2
// key: 2, value: 3
func forArray() {
	numbers := [3]int8{1, 2, 3}
	for key, value := range numbers {
		fmt.Printf("key: %d, value: %d\n", key, value)
	}
}

// 遍历 Slice
// index: 0, word: hello
// index: 1, word: effective
// index: 2, word: go
func forSlice() {
	slices := []string{"hello", "effective", "go"}
	for i, s := range slices {
		fmt.Printf("index: %d, word: %s \n", i, s)
	}
}

// 遍历 Map
// key: 1, value: 1
// key: 2, value: 2
// key: 3, value: 3
func forMap() {
	m := make(map[string]string)
	m["1"] = "1"
	m["2"] = "2"
	m["3"] = "3"

	for key, value := range m {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}

// 遍历 string
// index: 0, s as rune: U+0073, s as character: 's'
// index: 1, s as rune: U+0074, s as character: 't'
// index: 2, s as rune: U+0072, s as character: 'r'
// index: 3, s as rune: U+0069, s as character: 'i'
// index: 4, s as rune: U+006E, s as character: 'n'
// index: 5, s as rune: U+0067, s as character: 'g'
func forString() {
	for i, s := range "string" {
		// d 整数， U 码点， s 字符
		fmt.Printf("index: %d, s as rune: %U, s as character: '%c'\n", i, s, s)
	}
}

// Go 中可以多重赋值, 而 Java 中不能直接像这样转换变量： a[i], a[j] = a[j], a[i], 需要通过中间变量进行, string 零值为 ""
// edcba
func forMultiVariable() {
	a := make([]string, 10)
	a[0] = "a"
	a[1] = "b"
	a[2] = "c"
	a[3] = "d"
	a[4] = "e"

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for _, v := range a {
		fmt.Printf("%s", v)
	}
}

// int 零值 0
// 0000043210
func forMultiVariable2() {
	a := make([]int8, 10)
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for _, v := range a {
		fmt.Printf("%d", v)
	}
}
