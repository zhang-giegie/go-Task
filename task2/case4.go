package main

import (
	"fmt"
	"sync"
)

func task1() {

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("ch1:",num)
		}
	}()

	wg.Wait()

}


func task2(){

	ch2 := make(chan int,20)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()
		for num := range ch2 {
			fmt.Println("ch2:", num)
		}
	}()

	wg.Wait()
}

func main_case4() {
	task1()
	task2()
}
