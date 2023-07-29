package main

import (
	"net/http"
)

func main() {
	hh := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, This is GO-server!"))
	}

	http.HandleFunc("/hello", hh)
	http.ListenAndServe("", nil)
}
