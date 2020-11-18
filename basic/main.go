package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	// have deafult route behaviour
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// for logging to console
		log.Println("Hello World!")
		d, _ := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "error", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/api", func(http.ResponseWriter, *http.Request) {
		log.Println("api is hit")
	})

	http.ListenAndServe(":8000", nil)
}