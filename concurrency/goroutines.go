package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func f(msg string) {
	for i := 1; i < 3; i++ {
		fmt.Printf("cycle index:%d, mode: %s\n", i, msg)
	}
}

// 简单测试：直接调用 vs 协程调用
func simpleTest() {
	f("direct")

	// 使用关键字 `go` 启动一个轻量级协程
	go f("goroutines")

	time.Sleep(1 * time.Second)
}

// 利用同步锁
func testSync() {
	var wg sync.WaitGroup
	wg.Add(1) // 增加一个等待计数

	go func() {
		defer wg.Done() // 在goroutine完成时减去一个等待计数
		for i := 1; i < 3; i++ {
			fmt.Printf("cycle index:%d, mode: %s\n", i, "sync")
		}
	}()

	wg.Wait() // 等待所有的等待计数清零，即所有goroutine完成
}
