package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	// We are not using r, but we need it
	// because HandleFunc requires it

	fmt.Fprintf(w, "Hey, Hey Again")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe("0.0.0.0:8088", nil)
	fmt.Println("Running the server")
}
