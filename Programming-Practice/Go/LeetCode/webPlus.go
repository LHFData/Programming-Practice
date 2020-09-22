package main

import (
	"fmt"
	"net/http"
)

func index_handler_Plus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Plus</p>")
}
func main() {
	http.HandleFunc("/", index_handler_Plus)
	http.ListenAndServe(":8000", nil)
}
