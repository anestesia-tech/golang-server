package main

import (
	"fmt"
	"net/http"
	"log"
)

func hello(w, http.ResponseWriter, r *httpRequest) {
	fmt.Fprintf(w, "Hello!")
}


func output() {
	//
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAdnServe: ", err)
	}
}

