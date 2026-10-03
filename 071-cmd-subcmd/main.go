package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", true, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 1, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected foo or bar sub command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand", os.Args[1])
		fmt.Println("enable", *fooEnable)
		fmt.Println("name", *fooName)
		fmt.Println("tail", os.Args[2:])
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand", os.Args[1])
		fmt.Println("level", *barLevel)
		fmt.Println("tail", os.Args[2:])
	default:
		fmt.Println("exptected foo or bar sub command")
		os.Exit(1)
	}
}
