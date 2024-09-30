package main

import (
	"log"
	"os"
	"runtime/trace"
)

func Gc() {

	//  GODEBUG='gctrace1' go run main.go
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 1<<20)
	}
}

func GcTrace() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 启动 trace
	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	// 模拟程序运行
	for i := 0; i < 1000000; i++ {
		_ = make([]byte, 1<<20)
	}
}
