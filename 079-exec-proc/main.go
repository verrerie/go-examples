package main

import (
	"os"
	"os/exec"
	"syscall"
)

func ck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	bin, err := exec.LookPath("ls")
	ck(err)

	args := []string{"ls", "-l", "-a", "-t"}

	env := os.Environ()

	// not supported by windows
	err = syscall.Exec(bin, args, env)
	ck(err)
}
