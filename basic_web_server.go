package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	// We are not using r, but we need it
	// because HandleFunc requires it

	fmt.Fprintf(w, "Hey, Hey Again")
}

func main() {
	http.HandleFunc("/", index_handler)

	appengine.Main()
}
