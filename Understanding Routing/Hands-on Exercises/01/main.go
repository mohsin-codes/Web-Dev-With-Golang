package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func pic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", pic)
	http.ListenAndServe(":8080", nil)
}
