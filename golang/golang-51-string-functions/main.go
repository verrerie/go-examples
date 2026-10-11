package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	p(strings.Contains("test", "t"))
	p(strings.Count("test", "t"))
	p(strings.Clone("test"))
	p(strings.Compare("test", "tese"))
	p(strings.ContainsAny("test", "abcdefg"))
	p(strings.HasPrefix("test", "te"))
	p(strings.HasSuffix("test", "te"))
	p(strings.Index("test", "s"))
	p(strings.Join([]string{"t", "st"}, "e"))
	p(strings.Repeat("hello", 5))
	p(strings.Replace("test", "s", "8", 2))
	p(strings.ReplaceAll("test", "t", "7"))
	p(strings.Split("test", "e"))
	p(strings.ToUpper("Test"))
	p(strings.ToLower("TEST"))
}
