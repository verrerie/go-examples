package main

import (
	"fmt"
	"net/http"
	"time"
)

func serve(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("handler started")
	defer req.Body.Close()
	defer fmt.Println("handler stopped")

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("hello")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}

func main() {
	http.HandleFunc("/serve", serve)
	http.ListenAndServe(":8090", nil)
}
