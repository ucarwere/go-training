package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}

	fmt.Println(start)
}
