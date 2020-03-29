package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world, %s", req.URL.Path[1:])
}

func main () {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

