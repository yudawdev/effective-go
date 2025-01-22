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
