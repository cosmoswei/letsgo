package main

import (
	"fmt"
	"golang.org/x/net/context"
	"sync"
	"time"
)

var contextWg sync.WaitGroup

func doTask(n int) {
	time.Sleep(time.Duration(n) * time.Second)
	contextWg.Done()
}

func Context() {
	for i := 0; i < 3; i++ {
		contextWg.Add(1)
		go doTask(i + 1)
	}

	contextWg.Wait()
	fmt.Println("done")

	//go reqTask("worker1")
	//time.Sleep(time.Second * 3)
	//stop <- true
	//time.Sleep(time.Second * 3)

	ctx, cancelFunc := context.WithCancel(context.Background())
	go reqTaskContext(ctx, "worker2")
	time.Sleep(time.Duration(3) * time.Second)
	cancelFunc()
	time.Sleep(time.Duration(3) * time.Second)

	fmt.Println("===========================")

	ctx2, cancelFunc2 := context.WithCancel(context.Background())

	go reqTaskContext(ctx2, "worker3")
	go reqTaskContext(ctx2, "worker4")

	time.Sleep(time.Duration(3) * time.Second)
	cancelFunc2()
	time.Sleep(time.Duration(3) * time.Second)

	fmt.Println("===========================")

	ctx3, cancel3 := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx3, "options", &Options{1})

	go reqTaskContext2(vCtx, "worker5")
	go reqTaskContext2(vCtx, "worker6")

	time.Sleep(3 * time.Second)
	cancel3()
	time.Sleep(3 * time.Second)

	fmt.Println("===========================")

	ctx4, cancel4 := context.WithTimeout(context.Background(), 2*time.Second)
	go reqTaskContext2(ctx4, "worker7")
	go reqTaskContext2(ctx4, "worker8")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel4()
	time.Sleep(3 * time.Second)

	fmt.Println("===========================")

	ctx5, cancel5 := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	go reqTaskContext3(ctx5, "worker9")
	go reqTaskContext3(ctx5, "worker10")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel5()
	time.Sleep(3 * time.Second)

}

var stop chan bool = make(chan bool)

func reqTask(name string) {
	for {
		select {
		case <-stop:
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func reqTaskContext(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

type Options struct{ Interval time.Duration }

func reqTaskContext2(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
		}
	}
}

func reqTaskContext3(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name, ctx.Err())
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}
