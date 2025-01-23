package data

import (
	"fmt"
)

// 数组初始化
func array() {
	// 数组：固定长度，在编译时就确定大小
	arr1 := [10]int{}         // 长度为 10 的数组
	arr2 := [...]int{1, 2, 3} // 根据初始化值确定长度的数组

	fmt.Printf("arr1: %v, arr2: %v", arr1, arr2)
}
