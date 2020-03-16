package main

import "fmt"

func main() {
	i := make([]int, 5)
	f := make([]float64, 9)
	s := make([]string, 4)

	fmt.Println(len(i), len(f), len(s))
}
