package main

import (
	"fmt"
	"sync"
	"time"
)

func Channel() {
	ch := make(chan int) // 创建一个不带缓冲区的 channel
	var wg sync.WaitGroup

	wg.Add(2)

	// Goroutine 1：发送数据
	go func() {
		defer wg.Done()
		ch <- 42 // 发送数据
	}()

	// Goroutine 2：接收数据
	go func() {
		defer wg.Done()
		value := <-ch // 接收数据
		fmt.Println("Received:", value)
	}()

	wg.Wait()
}

func MutliChannel() {
	// 创建一个带缓冲区的通道，缓冲区大小为 100，避免发送和接收阻塞
	ch := make(chan int, 200)

	// 等待组，用于等待所有 Goroutine 完成
	var wg sync.WaitGroup

	// 启动 10 个 Goroutine 进行数据发送
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				ch <- id*10 + j // 发送数据
				fmt.Printf("Producer %d sent: %d\n", id, id*10+j)
			}
		}(i)
	}

	// 启动 5 个 Goroutine 进行数据消费
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				value := <-ch // 接收数据
				fmt.Printf("Consumer %d received: %d\n", id, value)
			}
		}(i)
	}

	wg.Wait() // 等待所有 Goroutine 完成
	close(ch) // 关闭通道
}

var wg sync.WaitGroup

func OrderPrint() {
	ch1 := make(chan struct{}, 3)
	ch2 := make(chan struct{}, 3)
	ch3 := make(chan struct{}, 3)

	ch1 <- struct{}{}

	wg.Add(3)

	start := time.Now().Unix()

	go printThis("goroutine 1", ch1, ch2)
	go printThis("goroutine 2", ch2, ch3)
	go printThis("goroutine 3", ch3, ch1)

	wg.Wait()

	end := time.Now().Unix()

	fmt.Println(end - start)

}

func printThis(goStr string, inputchan chan struct{}, outchan chan struct{}) {
	time.Sleep(2 * time.Second)
	select {
	case <-inputchan:
		fmt.Println(goStr)
		outchan <- struct{}{}
	}
	wg.Done()

}

func deadlock() {

}
