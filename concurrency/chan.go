package concurrency

import (
	"fmt"
	"time"
)

// chan 通道
func keyChan() {

	// 双向通道
	ch := make(chan int)

	ch <- 10

	// 只读通道
	var readCh <-chan int

	// 只写通道
	var wirteCh chan<- int

	// 带缓冲通道
	bufCh := make(chan int, 10)
}
