package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func panicHandler(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from: %v\n", r)
			debug.PrintStack()

			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "internal server error\n")
		}
	}()

	panic("oh no")

	io.WriteString(w, "everything is fine 🔥")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/panic", panicHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
