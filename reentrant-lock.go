package main

import (
	"fmt"
	"runtime"
	"sync"
)

type ReentrantLock struct {
	mu        sync.Mutex
	owner     int64 // 持有锁的 Goroutine ID
	recursion int   // 重入次数
}

// 获取当前 Goroutine ID
func getGoroutineID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := buf[10:n]
	var id int64
	for _, b := range idField {
		if b >= '0' && b <= '9' {
			id = id*10 + int64(b-'0')
		} else {
			break
		}
	}
	return id
}

// 加锁
func (r *ReentrantLock) Lock() {
	gid := getGoroutineID() // 获取当前 Goroutine ID

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.owner == gid { // 如果当前 Goroutine 已经持有锁，增加重入计数
		r.recursion++
		return
	}

	// 如果没有持有锁，设置持有锁的 Goroutine 并初始化重入计数
	for r.recursion != 0 { // 等待其他 Goroutine 完全释放锁
		r.mu.Unlock()
		r.mu.Lock()
	}
	r.owner = gid
	r.recursion = 1
}

// 解锁
func (r *ReentrantLock) Unlock() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.owner != getGoroutineID() { // 只有持有锁的 Goroutine 才能解锁
		panic("attempt to unlock a lock not held by this goroutine")
	}

	r.recursion-- // 减少重入计数
	if r.recursion == 0 {
		r.owner = 0 // 完全释放锁
	}
}

func ReentrantLockTest() {
	lock := &ReentrantLock{}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		lock.Lock()
		fmt.Println("First lock acquired")

		// 尝试再次加锁，重入锁允许这样做
		lock.Lock()
		fmt.Println("Second lock acquired")

		// 释放两次锁
		lock.Unlock()
		fmt.Println("First unlock done")

		lock.Unlock()
		fmt.Println("Second unlock done")
	}()

	wg.Wait()
}
