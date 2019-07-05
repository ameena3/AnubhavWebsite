package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/test", Test)
	fmt.Println("Listening at port 80")
	http.ListenAndServe(":80", nil)
}

//DefaultHandler ...
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there Anubhav \n")
	fmt.Println("Default endpoint hit")
}

//Test ...
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test endpoint hit")
}
