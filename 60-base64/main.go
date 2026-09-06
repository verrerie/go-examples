package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	res := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(res)

	decoded, _ := b64.StdEncoding.DecodeString(res)
	fmt.Println(string(decoded))

	res = b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(res)

	decoded, _ = b64.URLEncoding.DecodeString(res)
	fmt.Println(string(decoded))
}
