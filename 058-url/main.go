package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	fmt.Println(s)

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed: %#v\n", u)

	fmt.Println(u.Scheme)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	h, p, _ := net.SplitHostPort(u.Host)
	fmt.Println(h, p)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Printf("%+v", m)
	fmt.Println(m["k"][0])
}
