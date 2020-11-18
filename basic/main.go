package main

import (
	"net/http",
	"log"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		// for logging to console
		log.Println("Hello World")
	})

	http.ListenAndServe(":8000", nil)
}