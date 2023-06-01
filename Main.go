package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "This methood is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func formHandlaar(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "Post Request Sucessful")
	first_name := r.FormValue("fname")
	last_name := r.FormValue("lname")
	fmt.Fprintf(w, "Name = %s \n", first_name)
	fmt.Fprintf(w, "Last name = %s \n", last_name)

}

func main() {
	fileserver := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandlaar)
	http.HandleFunc("/hello", hellohandler)

	fmt.Printf("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
