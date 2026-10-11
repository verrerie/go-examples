package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p.*h", "peach")
	fmt.Println("match:", match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println("match:", r.MatchString("peach"))
	fmt.Println("found:", r.FindString("peach punch"))
	fmt.Println("find indexes:", r.FindStringIndex("peach punch"))
	fmt.Println("find sub:", r.FindStringSubmatch("peach punch"))
	fmt.Println("find sub:", r.FindStringSubmatchIndex("peach punch"))
	fmt.Println("all:", r.FindAllString("peach punch pinch", -1))
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println("all:", r.FindAllString("peach punch pinch", 2))
	fmt.Println("match:", r.Match([]byte("peach")))
	fmt.Println("repalced:", r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("repaced:", string(out))
}
