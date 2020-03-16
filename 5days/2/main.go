package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Sample struct {
	X int
	Y int
}

type Accnt struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Rank     int    `json:"rank"`
}

const Pi float32 = 3.141519

type Person struct {
	name string
	age  uint
}

func (p *Person) profile_1(header string) {
	p.name = header
}

func (p Person) profile_2(header string) {
	p.name = header
}

func main () {
	// boolean
	var sample bool
	sample = true
	fmt.Printf("%t\n", sample)

	// number 
	var a int64
	a = int64(300)
	fmt.Printf("%d\n", a)

	var f float32
	f = float32(0.007)
	fmt.Printf("%f\n", f)

	var f64 float64
	f64 = float64(0.007)
	fmt.Println(f64)

	var c complex64
	fmt.Printf("%f\n", c)

	var c128 complex128
	c128 = complex(1, 2)
	var c128_real float64 = real(c128)
	var c128_imag float64 = imag(c128)
	fmt.Println(c128_real, c128_imag)

	// string
	var str string
	str = "Hello World"
	fmt.Println(str, len(str))

	// type
	var s Sample
	fmt.Printf("(X, Y) = (%d, %d)\n", s.X, s.Y)

	// Pointer
	var p *int
	n := 10

	p = &n
	// address
	fmt.Println(p)
	// value
	fmt.Println(*p)
	n = 20

	// new
	var p_new *int = new(int)

	fmt.Println(p_new)
	fmt.Println(*p_new)

	// array
	var arr[3] int
	fmt.Println(arr, len(arr))

	num := [...]int{1, 2, 3, 4, 5}
	fmt.Println(num, len(num))

	// slice
	var arr_a[] int
	arr_b := [] string{"hoge", "fuga"}
	arr_c := [] int{1, 2}
	fmt.Println(arr_a, len(arr_a))
	fmt.Println(arr_b, len(arr_b))
	fmt.Println(arr_c, len(arr_c))

	// map
	a_map := map[string]int{"hello": 1, "world": 2}
	fmt.Println(a_map)

	// struct
	type strct struct {
		alpha string
		beta, gamma int
	}

	val := new(strct)
	fmt.Println(val.alpha, val.beta, val.gamma)
	val_accnt := &Accnt {
		Email: "hoge@fuga.com",
		Password: "password",
		Rank: 1,
	}

	// encode
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(val_accnt)
	fmt.Println(buf.String())

	// decode
	decoded := new(Accnt)
	json.NewDecoder(buf).Decode(decoded)
	fmt.Println(decoded)

	// method
	person_a := &Person{
		name: "vankichi",
		age: 29,
	}
	person_a.profile_1("kichivan")
	fmt.Println(person_a.name)
	person_b := &Person{
		name: "vankichi",
		age: 29,
	}
	person_b.profile_2("kichivan")
	fmt.Println(person_b.name)
}
