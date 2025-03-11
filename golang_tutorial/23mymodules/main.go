package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r:=mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":4000",r))
}
func greeter() {
	fmt.Println("hello greeting you\n")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to the golang stuff by Sahil </h1>"))
}
