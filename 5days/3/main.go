package main

import (
	"fmt"
)

func main () {
	val := 5
	if val == 20 {
		fmt.Printf("answer = %d\n", val)
	} else {
		fmt.Printf("answer = %T\n", val)
	}

	val = 100
	switch {
	case val == 1:
		fmt.Println(val)
	case val == 10:
		fmt.Println(val)
	case val == 100:
		fmt.Println(val)
		fallthrough
	default:
		fmt.Println("end")
	}

	// loop
	for i := 0 ; i < 1 ; i++ {
		fmt.Println(i)
	}

	for i, t := range [] int {5, 6, 7} {
		fmt.Println(i, t)
	}

	// func
	var addFunc func(int, int) int
	addFunc = add
	ans := addFunc(1, 2)
	fmt.Println(ans)

	clFunc := closer()
	// res will be 1
	res := clFunc()
	fmt.Println(res)
	// res will be 2
	res = clFunc()
	fmt.Println(res)

	defFunc := def
	defFunc()

	// panic
	pnc := callPanic
	pnc()
}

func add(a int, b int) int {
	return a + b
}

func closer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func def() {
	defer fmt.Println("exist w/success")
	fmt.Println("called func def")
}

func callPanic() int {
	fmt.Println("call panic")
	panic("panic is called")
}
