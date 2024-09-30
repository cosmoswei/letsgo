package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

func Pprof() {
	for i := 0; i < 1000; i++ {
		go func() {
			select {}
		}()

		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
	}
	// go tool pprof -http=:1248 http://127.0.0.1:6060/debug/pprof/goroutine
	select {}
}

func CpuPProf() {
	// 创建文件保存 CPU 剖析数据
	f, err := os.Create("cpu.pprof")
	// go tool pprof cpu.pprof
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 开始 CPU 剖析
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// 模拟一些 CPU 任务
	for i := 0; i < 10000000; i++ {
	}
}
