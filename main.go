package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	fmt.Println("Running server!")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello2", helloHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
