package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "useage: ./fetchall\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	urls := os.Args[2:]

	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		// start goroutine
		go fetch(url, ch)
	}
	for range urls {
		// recieve ch channel
		fmt.Println(<-ch)
	}
	fmt.Fprintf(out, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// send err to channel
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
