package concurrency

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// 消费者
// 消费一个需要花费 1s
func consumer(i <-chan int, name string) {
	fmt.Printf("consuemr %s initial...\n", name)
	for val := range i {
		fmt.Printf("%s consumer | consume %d\n", name, val)
		time.Sleep(time.Second) // 模拟处理时间
	}
}

// 生产者
// 每隔 1s 中生产一个
func producer(ch chan<- int, name string, done <-chan bool) {
	fmt.Printf("producer %s initial...\n", name)

	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			i := rand.IntN(100)
			fmt.Printf("%s producer | produce %d\n", name, i)
			ch <- i
		case <-done:
			fmt.Printf("producer %s is done\n", name)
			return
		}
	}
}

func main() {
	ch := make(chan int, 10) // 创建双向通道 | 为什么不创建缓冲通道呢？
	done := make(chan bool)

	// 创建三个生产者
	for i := 0; i < 3; i++ {
		go producer(ch, strconv.Itoa(i), done)
	}

	// 创建三个消费者
	for i := 0; i < 3; i++ {
		go consumer(ch, strconv.Itoa(i))
	}

	// 信号处理
	sigChan := make(chan os.Signal, 1)
	// 注册要监听的信号：SIGINT(Ctrl+C)和SIGTERM(终止信号)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞等待信号
	s := <-sigChan
	fmt.Println("终止信号：", s)

	// 优雅退出
	close(done)
	close(ch)
}
