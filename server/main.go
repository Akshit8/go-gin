package main

import (
	"fmt"
	"log"
	"text/template"
	"io"
	"net/http"
)

func main() {

	// http.HandleFunc("/", defaultHandler)
	// http.HandleFunc("/dog/", dogHandler)
	// http.HandleFunc("/me/", meHandler)

	http.Handle("/", http.HandlerFunc(defaultHandler))
	http.Handle("/dog/", http.HandlerFunc(dogHandler))
	http.Handle("/me/", http.HandlerFunc(meHandler))

	http.ListenAndServe(":8000", nil)
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "default handler ran")
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "dog handler ran")
}

func meHandler(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("./me.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "me.gohtml", "Akshit")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
