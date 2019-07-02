package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.ListenAndServe(":8080", nil)
}

//DefaultHandler ...
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hey there Anubhav")
}
