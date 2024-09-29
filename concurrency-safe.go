package main

import (
	"log"
	"sync"
)

func ConcurrencyNotSafe() {
	a := make([]int, 0)

	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			a = append(a, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Print(len(a))
}

func ConcurrencySafe() {
	a := make([]int, 0)
	var mu sync.Mutex // 创建互斥锁

	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()        // 加锁，确保并发安全
			a = append(a, i) // 修改切片
			mu.Unlock()      // 解锁
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Print(len(a)) // 打印切片长度
}
