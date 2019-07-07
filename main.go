package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening at port 8080")
	if err := http.ListenAndServe(":8080", http.FileServer(http.Dir("Public"))); err != nil {
		fmt.Println(err.Error())
	}

}
