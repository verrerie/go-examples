package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "poing")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, header := range req.Header {
		fmt.Fprintln(w, name, ":", header)
	}
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
