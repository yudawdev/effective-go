package concurrency

import (
	"fmt"
	"sort"
)

// 通过协程进行并发
func testListSort() {
	list := []int{10, 9, 11, 55, 33, 11, 43, 2, 50, 623, 23}

	c := make(chan int)

	// 并发做一些事情
	go func() {
		sort.Ints(list)
		c <- 1
		fmt.Println("sort done!")
	}()

	func() {
		fmt.Println("do something....")
		for i := 0; i < 100; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println("do something done!")
	}()

	// 在执行后续逻辑时，必须等 sort 完成
	<-c
	fmt.Println("等 sort 完成后。。。")
}

// 延伸问题
// 1. fmt.Println("sort done!") 这行代码会打印吗？
// 2. 如果想在程序结束前，必须打印出 fmt.Println("sort done!") 该如何编排代码？
