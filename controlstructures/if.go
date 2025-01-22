package controlstructures

import "fmt"

// true
func ifSimple() {
	if true {
		fmt.Println(true)
	}
}

// 1
func ifInitial() {
	if i := 1; i > 0 {
		fmt.Println(i)
	}
}

// i > 0
func ifReturn() {
	i := 1
	if i < 0 {
		fmt.Println("i < 0")
		return
	}

	fmt.Println("i > 0")
}
