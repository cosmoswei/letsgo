package main

import (
	"fmt"
	"log"
	"sync"
	"time"
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

func ConcurrencyMapNotSafe() {
	s := make(map[int]int)
	for i := 0; i < 100000; i++ {
		go func(i int) {
			s[i] = i
		}(i)
	}
	for i := 0; i < 100000; i++ {
		go func(i int) {
			fmt.Printf("map 的第 %d 个元素是 %d", i, s[i])
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func ConcurrencyMapSafe() {
	s := make(map[int]int)
	var lock sync.RWMutex
	for i := 0; i < 100000; i++ {
		go func(i int) {
			lock.Lock()
			s[i] = i
			lock.Unlock()
		}(i)
	}
	for i := 0; i < 100000; i++ {
		go func(i int) {
			lock.RLock()
			fmt.Printf("map 的第 %d 个元素是 %d \n", i, s[i])
			lock.RUnlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func ConcurrencyMapSafe2() {
	var m sync.Map
	for i := 0; i < 100000; i++ {
		go func(i int) {
			m.Store(i, i)
		}(i)
	}
	for i := 0; i < 100000; i++ {
		go func(i int) {
			value, ok := m.Load(i)
			fmt.Printf("map 的第 %d 个存在？%v，元素是 %d \n", i, ok, value)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
