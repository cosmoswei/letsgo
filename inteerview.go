package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
https://zhuanlan.zhihu.com/p/360306642
*/
func InterView() {
	numsPtr := new(int)
	fmt.Println(*numsPtr)
	*numsPtr = 10
	fmt.Println(*numsPtr)

	p2person := new(Person)
	p2person.Age = 1
	p2person.Name = "huangxuwei"
	fmt.Println(p2person)
}

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eat")
}

func Reflect() {
	animal := Animal{}
	reflVal := reflect.ValueOf(&animal)
	f := reflVal.MethodByName("Eat")
	f.Call([]reflect.Value{})
}

func ArrAndSince() {
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 4, 5, 6, 7, 8}
	for i, v := range slice {
		fmt.Println(i, v)
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	var a uint8 = 255
	var b uint8 = 1

	fmt.Println(a + b)

}

func modifySlice(s []int) {
	// 修改 slice 中的值
	s[0] = 100
	fmt.Println("Inside function:", s) // [100 2 3]
}

func StaticLen() {
	slice := []int{1, 2, 3}                     // 创建一个 slice
	fmt.Println("Before function call:", slice) // [1 2 3]

	modifySlice(slice) // 传入 slice

	fmt.Println("After function call:", slice) // [100 2 3]
}

func modifySlice2(s []int) {
	// 触发扩容，增加一个元素
	s = append(s, 4)
	s[0] = 100
	fmt.Println("Inside function after append:", s) // [100 2 3 4]
}

func extendLen() {
	slice := []int{1, 2, 3}                     // 创建一个 slice，长度为 3
	fmt.Println("Before function call:", slice) // [1 2 3]

	modifySlice2(slice) // 传入 slice

	fmt.Println("After function call:", slice) // [1 2 3]
}

func Select() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		{
			fmt.Println(msg1)
		}
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default:
		fmt.Println("No message received")
	}
}
