package concurrency

import (
	"fmt"
	"time"
)

// select 关键字，用于处理多个通道。阻塞读取任一已完成通道
func keySelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "消息1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "消息2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("收到", msg1)
	case msg2 := <-ch2:
		fmt.Println("收到", msg2)
		//case <-time.After(1 * time.Second):
		//	fmt.Println("超时")
		//default:
		//	// 非阻塞操作
		//	fmt.Println("没有可用的通道")
	}
}

// select 非阻塞，使用 default 关键字
func selectNonblock() {
	ch1 := make(chan int)

	go func() {
		ch1 <- 111
	}()

	select {
	case i := <-ch1:
		fmt.Println("收到", i)
	default:
		fmt.Println("无可用通道")
	}

}

// select 超时： <-time.After()
func selectTimeout() {
	chq := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chq <- "1111"
	}()

	select {
	case msg1 := <-chq:
		fmt.Println("收到", msg1)
	case <-time.After(1 * time.Second):
		fmt.Println("超时")
	}
}

// 系统阻塞
func selectEmptyBlock() {
	select {}
}
