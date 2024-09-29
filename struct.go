package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) IncrAge1() {
	p.Age = p.Age + 1
}

func (p Person) IncrAge2() {
	p.Age = p.Age + 1
}

func (p Person) getAge() int {
	return p.Age
}

func Struct() {
	person := Person{
		"huangxuwei", 25,
	}
	person.IncrAge1()
	fmt.Println(person.getAge())
	person2 := Person{
		"huangxuwei", 25,
	}
	person2.IncrAge2()
	fmt.Println(person2.getAge())
}
