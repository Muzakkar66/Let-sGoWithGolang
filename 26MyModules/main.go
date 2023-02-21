package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to MOD in golang complete")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", servehome).Methods("Get")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hey there mode users")
}
func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Mode Users.</h1>"))
}
