package main

import (
	"fmt"
	"net/http"
)

func isAlive(w http.ResponseWriter, req *http.Request) {

	fmt.Println("asking to check if alive")
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/isAlive", isAlive)
	http.HandleFunc("/headers", headers)

	fmt.Println("starting server")

	http.ListenAndServe(":30304", nil)
}
