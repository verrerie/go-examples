package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	ck(err)
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("failed executing:", e)
		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode := e.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		} else {
			ck(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep\nhhhhhhelloooooo\nhell oh no\n"))
	grepIn.Close()
	grepBytes, err := io.ReadAll(grepOut)
	ck(err)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	ck(err)
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

func ck(err error) {
	if err != nil {
		panic(err)
	}
}
