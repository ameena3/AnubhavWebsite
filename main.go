package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening at port 80")
	if err := http.ListenAndServe(":80", http.FileServer(http.Dir("Public"))); err != nil {
		fmt.Println(err.Error())
	}

}
