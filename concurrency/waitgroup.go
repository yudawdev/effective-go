package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func testSyncGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			if id == 3 {
				time.Sleep(2 * time.Second)

			}
			fmt.Printf("Worker %d 完成\n", id)
		}(i)
	}

	wg.Wait()

	fmt.Printf("All Worker ")
}
