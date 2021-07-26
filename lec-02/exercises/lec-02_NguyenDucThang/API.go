package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Input struct {
	OP int
	X  int
	Y  int
}

func main() {
	fmt.Println("Start!")
	defer func() {
		fmt.Println("End!")
	}()
	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/mul", mul)
	http.HandleFunc("/div", div)
	http.ListenAndServe(":8080", nil)
}
func add(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Query().Get("x")
	y := r.URL.Query().Get("y")
	var a = 0
	var b = 0
	a, _ = strconv.Atoi(x)
	b, _ = strconv.Atoi(y)
	math := &Input{OP: a + b, X: a, Y: b}
	re, err := json.Marshal(math)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v", string(re))

}

func sub(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Query().Get("x")
	y := r.URL.Query().Get("y")
	var a = 0
	var b = 0
	a, _ = strconv.Atoi(x)
	b, _ = strconv.Atoi(y)
	math := &Input{OP: a - b, X: a, Y: b}
	re, err := json.Marshal(math)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v", string(re))
}
func mul(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Query().Get("x")
	y := r.URL.Query().Get("y")
	var a = 0
	var b = 0
	a, _ = strconv.Atoi(x)
	b, _ = strconv.Atoi(y)
	math := &Input{OP: a * b, X: a, Y: b}
	re, err := json.Marshal(math)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v", string(re))
}
func div(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Query().Get("x")
	y := r.URL.Query().Get("y")
	var a = 0
	var b = 0
	a, _ = strconv.Atoi(x)
	b, _ = strconv.Atoi(y)
	if b != 0 {
		math := &Input{OP: a / b, X: a, Y: b}
		re, err := json.Marshal(math)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", string(re))
	} else {
		fmt.Println("Error")
	}
}
