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

// 切片 - 长度、容量以及添加元素方式
// ints ----------------------------------
// index: 0, value: 0
// index: 1, value: 0
// index: 2, value: 0
// index: 3, value: 0
// index: 4, value: 0
// index: 5, value: 0
// index: 6, value: 0
// index: 7, value: 0
// index: 8, value: 0
// index: 9, value: 0
// index: 10, value: 1
// index: 11, value: 2
// index: 12, value: 3
// index: 13, value: 4
// index: 14, value: 5
// index: 15, value: 6
// index: 16, value: 7
// index: 17, value: 8
// index: 18, value: 9
// index: 19, value: 10
// ints2 ----------------------------------
// index: 0, value: 1
// index: 1, value: 2
// index: 2, value: 3
func sliceLenCapAppend() {
	ints := make([]int, 10) // 长度和容量都是 10
	ints = append(ints, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("ints ---------------------------------- \n")
	for index, i := range ints {
		fmt.Printf("index: %d, value: %d\n", index, i)
	}

	fmt.Printf("ints2 ---------------------------------- \n")
	ints2 := make([]int, 0, 10) // 长度是0，容量10
	ints2 = append(ints2, 1, 2, 3)
	for index, i := range ints2 {
		fmt.Printf("index: %d, value: %d\n", index, i)
	}
}

// slices[a:b] --- 获取 slices 切片中，底层索引 [a, b) 间的元素
// i1 := ints[:]----------------------------------
// index: 0, value: 1
// index: 1, value: 2
// index: 2, value: 3
// index: 3, value: 4
// index: 4, value: 5
// i2 := ints[1:]----------------------------------
// index: 0, value: 2
// index: 1, value: 3
// index: 2, value: 4
// index: 3, value: 5
// i3 := ints[:2]----------------------------------
// index: 0, value: 1
// index: 1, value: 2
// i4 := ints[1:2]----------------------------------
// index: 0, value: 2
// index: 1, value: 3
// index: 2, value: 4
// index: 3, value: 5
func sliceIndex() {
	ints := make([]int, 0, 32)
	ints = append(ints, 1, 2, 3, 4, 5)

	fmt.Printf("i1 := ints[:]---------------------------------- \n")
	i1 := ints[:]
	for index, i := range i1 {
		fmt.Printf("index: %d, value: %d\n", index, i)
	}

	fmt.Printf("i2 := ints[1:]---------------------------------- \n")
	i2 := ints[1:]
	for index, i := range i2 {
		fmt.Printf("index: %d, value: %d\n", index, i)
	}

	fmt.Printf("i3 := ints[:2]---------------------------------- \n")
	i3 := ints[:2]
	for index, i := range i3 {
		fmt.Printf("index: %d, value: %d\n", index, i)
	}

	fmt.Printf("i4 := ints[1:2]---------------------------------- \n")
	i4 := ints[1:5]
	for index, i := range i4 {
		fmt.Printf("index: %d, value: %d\n", index, i)
	}
}

// 切片 - 去除元素 - 正向遍历，删除元素后，进行数组移动
// i: 1
// i: 2
// i: 3
// i: 4
// i: 6
// i: 7
// i: 8
// i: 9
// i: 10
func sliceRemoveElement() {
	s := make([]int, 0, 10)
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for i := 0; i < len(s); i++ {
		if meetCondition(s[i]) {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}

	for _, i := range s {
		fmt.Printf("i: %d\n", i)
	}

}

// 切片 - 去除元素 - 逆向遍历，删除元素后，同样进行数组移动
// i: 1
// i: 2
// i: 3
// i: 4
// i: 6
// i: 7
// i: 8
// i: 9
// i: 10
func sliceRemoveElement2() {
	s := make([]int, 0, 10)
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	for i := len(s) - 1; i >= 0; i-- {
		if meetCondition(s[i]) {
			s = append(s[:i], s[i+1:]...)
		}
	}

	for _, i := range s {
		fmt.Printf("i: %d\n", i)
	}

}

// 切片 - 去除元素 - 直接创建新的动态数组
// i: 1
// i: 2
// i: 3
// i: 4
// i: 6
// i: 7
// i: 8
// i: 9
// i: 10
func sliceRemoveBestPractice() {
	s := make([]int, 0, 10)
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	nSlice := make([]int, 0, len(s))

	for i := 0; i < len(s); i++ {
		if !meetCondition(s[i]) {
			nSlice = append(nSlice, s[i])
		}
	}

	for _, i := range nSlice {
		fmt.Printf("i: %d\n", i)
	}
}

func meetCondition(i int) bool {
	return i == 5
}
