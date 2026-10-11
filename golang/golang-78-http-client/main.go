package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func ck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	resp, err := http.Get("https://gobyexample.com")
	ck(err)
	defer resp.Body.Close()

	fmt.Println("status:", resp.Status)

	scan := bufio.NewScanner(resp.Body)
	for range 20 {
		scan.Scan()
		fmt.Println(scan.Text())
	}
	ck(scan.Err())
}
