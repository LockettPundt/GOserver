package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/slayer/autorestart"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	autorestart.StartWatcher()
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	//http.HandleFunc("/hello", routes.helloHandler)

	port := ":5055"
	fmt.Printf(" ✔ Server starting on PORT %v ✔  \n", port)
	if err := http.ListenAndServe(":5055", nil); err != nil {
		log.Fatal(err)
	}
}
