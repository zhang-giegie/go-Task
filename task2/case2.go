package main

import (
	"fmt"
	"sync"
	"time"
)

// 打印1到10的奇数
func printOddNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数:", i)
	}
}

// 打印2到10的偶数
func printEvenNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数:", i)
	}
}

// 并发执行任务并统计执行时间
func task(name string, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	fmt.Printf("%s 开始\n", name)
	time.Sleep(duration)
	fmt.Printf("%s 完成，用时 %v\n", name, time.Since(start))
}

func main_case2() {
	var wg sync.WaitGroup

	// 启动打印奇数和偶数的协程
	wg.Add(2)
	go printOddNumbers(&wg)
	go printEvenNumbers(&wg)
	wg.Wait()

	// 定义任务列表
	tasks := []struct {
		name     string
		duration time.Duration
	}{
		{"任务1", 2 * time.Second},
		{"任务2", 1 * time.Second},
		{"任务3", 3 * time.Second},
	}

	wg.Add(len(tasks))
	for _, t := range tasks {
		go task(t.name, t.duration, &wg)
	}
	wg.Wait()
	fmt.Println("所有任务已完成")
}
