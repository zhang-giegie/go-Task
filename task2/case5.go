package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func mutexCounter() {
	var counter int64
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("使用sync.Mutex保护的计数器值:", counter)

}

func atomicCounter() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("使用atomic保护的计数器值:", counter)
}

func main() {
	mutexCounter()
	atomicCounter()
}
