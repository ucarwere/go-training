package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// for k, v := range r.Header {
	// 	fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// }
	// fmt.Fprintf(w, "Host = %q\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// if err := r.ParseForm(); err != nil {
	// 	log.Print(err)
	// }
	// for k, v := range r.Form {
	// 	fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	// }
	// mu.Lock()
	// count++
	// mu.Unlock()
	// fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
	cycles := 5
	res := 0.001
	size := 100
	nframes := 64
	delay := 8

	if err := r.ParseForm; err != nil {
		log.Print(err)
	}

	for param, values := range r.Form {
		value := values[0]

		// We are going to convert existing parameters only if the conversion
		// succeeds. Otherwise just ignore them.
		switch param {
		case "cycles":
			if i, err := strconv.Atoi(value); err == nil {
				cycles = i
			}
		case "res":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				res = f
			}
		case "size":
			if i, err := strconv.Atoi(value); err == nil {
				size = i
			}
		case "nframes":
			if i, err := strconv.Atoi(value); err == nil {
				nframes = i
			}
		case "delay":
			if i, err := strconv.Atoi(value); err == nil {
				delay = i
			}
		}
	}

	lissajous(w, cycles, res, size, nframes, delay)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "COUNT = %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(2*cycles)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
