package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/test", Test)
	http.Handle("/custom", &Myhandler{greeting: "Anubhav"})
	http.Handle("/custom1", http.TimeoutHandler(&Myhandler{greeting: "Anubhav"}, time.Second, "hey there timeout"))
	fmt.Println("Listening at port 8080")
	http.ListenAndServe(":8080", nil)
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

//Myhandler test
type Myhandler struct {
	greeting string
}

func (mh *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v hello", mh.greeting)))
	time.Sleep(4 * time.Second)
}
