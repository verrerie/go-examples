package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	server, err := net.Listen("tcp", ":8090")
	ck(err)
	defer server.Close()

	for {
		conn, err := server.Accept()
		ck(err)
		go handleTCP(conn)
	}
}

func handleTCP(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, err := reader.ReadString(byte('\n'))
	ck(err)

	reply := strings.ToUpper(message)
	response := fmt.Sprintf("ACK:%v", reply)
	_, err = conn.Write([]byte(response))
	ck(err)
}

func ck(e error) {
	if e != nil {
		fmt.Println("server error")
		panic(e)
	}
}
