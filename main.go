package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TEST, TEST, TEST!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", TestHandler)
	http.Handle("/", router)
	fmt.Println("Everything is set up!")
}
