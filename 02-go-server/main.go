package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//check if the route is not /hello
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" { //method is not GET
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post Request successful")

	user_name := r.FormValue("user_name")
	user_email := r.FormValue("user_email")
	user_message := r.FormValue("user_message")

	fmt.Fprintf(w, "Name = %s\n", user_name)
	fmt.Fprintf(w, "Email = %s\n", user_email)
	fmt.Fprintf(w, "Message = %s\n", user_message)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	//route handling
	http.Handle("/", fileServer) //serves the index.html file
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server starting on port 8080\n") //prinf the place the serve is starting

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Couldn't connect to server")
	}
}
