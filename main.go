package main

import (
	"bookingApp/controller"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /user", controller.CreateUser)
	mux.HandleFunc("GET /user", controller.FindAllUsers)
	
	http.ListenAndServe(":8080", mux)
}

func handleRoot (
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintf(w, "Hello World!!")
}
