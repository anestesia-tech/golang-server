package main

import (
	"fmt"
	"net/http"
	"log"
)

func hello(w, http.ResponseWriter, r *httpRequest) {
	fmt.Fprintf(w, "Hello!")
}

func goodbye(w, http.ResponseWriter, r *httpRequest) {
	fmt.Fprintf(w, "Bye!")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAdnServe: ", err)
	}
}

