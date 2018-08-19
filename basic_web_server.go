package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"google.golang.org/appengine"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	// We are not using r, but we need it
	// because HandleFunc requires it

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	http.HandleFunc("/", index_handler)

	appengine.Main()
}
