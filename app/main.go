package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get PORT from environment variable
	port := os.Getenv("SERVER_PORT")

	// Define handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("pong")
		writer.Write([]byte("it works!"))
	})

	fmt.Printf("Listening on port %s... \n", port)
	http.ListenAndServe(port, mux)
}
