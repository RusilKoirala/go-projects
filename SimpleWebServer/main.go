package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is running at http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: v%", err)
		return
	}

	fmt.Fprintf(w, "POST request succesful")
	name, email, age, address := r.FormValue("name"), r.FormValue("email"), r.FormValue("age"), r.FormValue("address")
	fmt.Printf("Name: %s\nEmail: %s\nAge: %s\nAddress: %s\n\n", name, email, age, address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		WriteError(w, "404 Not Found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		WriteError(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Hello User"))
}

func WriteError(w http.ResponseWriter, msg string, code int) {
	http.Error(w, msg, code)
}
