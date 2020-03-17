package main

import (
	"fmt"
)

type Human interface {
	greeting()
	selfintroduction()
}

type Person struct {
	name string
	age  int
}

func (ps Person) greeting() {
	fmt.Printf("my name is %s\n", ps.name)
}

func (ps Person) selfintroduction() {
	fmt.Printf("my name is %s, I'm %d old\n", ps.name, ps.age)
}

type Calc interface {
	Calc(a int, b int) int
}

type Add struct {
}

func (n Add) Calc(a, b int) int {
	return a + b
}

type Sub struct {
}

func (n Sub) Calc(a, b int) int {
	return a - b 
}

func main() {
	self := Person{ "vankichi", 20 }
	var h Human = self

	h.greeting()
	h.selfintroduction()

	var num Calc
	var add Add
	var sub Sub

	num = add
	fmt.Println(num.Calc(10, 20))
	num = sub
	fmt.Println(num.Calc(20, 10))
}
