package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello JUG!")
	http.HandleFunc("/hi", hello)
	http.ListenAndServe(":8888", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
