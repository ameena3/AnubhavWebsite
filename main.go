package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Listening at port 80")
	templates := loadtemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		if requestedFile == "" {
			requestedFile = "index"
		}
		t := templates.Lookup(requestedFile + ".html")

		if t != nil {
			err := t.Execute(w, nil)
			catcherror(err)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/images/", http.FileServer(http.Dir("Public")))
	http.Handle("/css/", http.FileServer(http.Dir("Public")))
	http.Handle("/fonts/", http.FileServer(http.Dir("Public")))
	http.Handle("/js/", http.FileServer(http.Dir("Public")))
	http.HandleFunc("/control", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the control.")
	})
	err := http.ListenAndServe(":8080", nil)
	catcherror(err)

}

func loadtemplates() *template.Template {
	result := template.New("templates")
	template.Must(result.ParseGlob("templates/*.html"))
	return result
}

func catcherror(err error) {
	if err != nil {
		fmt.Println(err)

	}
}
