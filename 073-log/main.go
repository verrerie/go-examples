package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("microsecond log")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("shortfile log")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("sunday:")
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf-", log.LstdFlags)
	buflog.Println("from buflog")

	fmt.Println("buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hello world")

	myslog.Info("hello again !", "key", "val", "age", 21, "error", errors.New("some error"))
}
