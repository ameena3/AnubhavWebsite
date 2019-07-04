package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.ListenAndServe(":8100", nil)
}

//DefaultHandler ...
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hey there Anubhav \n")
}
