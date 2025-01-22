package data

import "fmt"

// 数组初始化
func array() {
	// 数组：固定长度，在编译时就确定大小
	arr1 := [10]int{}         // 长度为 10 的数组
	arr2 := [...]int{1, 2, 3} // 根据初始化值确定长度的数组

	fmt.Printf("arr1: %v, arr2: %v", arr1, arr2)
}

// 切片初始化
func slices() {
	// 方式1：make 函数
	slice1 := make([]int, 10)     // 长度和容量都是 10
	slice2 := make([]int, 10, 20) // 长度 10，容量 20
	fmt.Printf("s1: %v, s2: %v", slice1, slice2)

	// 方式2：切片字面量
	slice3 := []int{1, 2, 3} // 长度和容量都是 3
	fmt.Printf("s3: %v", slice3)

	// 方式3：从数组创建
	arr := [5]int{1, 2, 3, 4, 5}
	slice4 := arr[1:3] // 长度 2，容量 4
	fmt.Printf("s4: %v", slice4)
}
