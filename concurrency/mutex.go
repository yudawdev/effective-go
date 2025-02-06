package concurrency

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func testMultiGoroutinesHandle() {
	wg := sync.WaitGroup{}

	// 初始化方式一
	//c := Counter{
	//	mu:    sync.Mutex{},
	//	count: 0,
	//}

	// 初始化方式二（推荐）
	//c := Counter{count: 0}

	// 初始化方式三
	var c Counter

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("goroutines %d\n", id)
			c.Increment()
		}(i)
	}

	wg.Wait()

	fmt.Println("count: ", c.count)
}
