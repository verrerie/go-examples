package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Resp1 struct {
	Page   int
	Fruits []string
}

type Resp2 struct {
	Page   int      `json:"Page"`
	Fruits []string `json:"Fruits"`
}

func main() {
	boolB, _ := json.Marshal(true)
	fmt.Println("bool:", string(boolB))

	intB, _ := json.Marshal(100)
	fmt.Println("int:", string(intB))
	floatB, _ := json.Marshal(10.1)
	fmt.Println("float:", string(floatB))
	strB, _ := json.Marshal("hello")
	fmt.Println("str:", string(strB))
	sliceB, _ := json.Marshal([]string{"go", "python"})
	fmt.Println("slice:", string(sliceB))
	mapB, _ := json.Marshal(map[string]int{"go": 100, "python": 101})
	fmt.Println("map:", string(mapB))

	r1 := Resp1{
		1,
		[]string{"apple", "orange"},
	}
	s1, _ := json.Marshal(r1)
	fmt.Println("s1:", string(s1))
	r2 := Resp2{
		2,
		[]string{"kiwi", "clementine"},
	}
	s2, _ := json.Marshal(r2)
	fmt.Println("s2:", string(s2))

	jsonB := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var data map[string]any
	if err := json.Unmarshal(jsonB, &data); err != nil {
		panic(err)
	}
	fmt.Printf("data:%+v\n", data)
	fmt.Println("num:", data["num"].(float64))
	fmt.Println("strs:", data["strs"].([]any))
	fmt.Println("strs[0]:", data["strs"].([]any)[0])

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Resp2{}
	_ = json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits)

	d := map[string]int{"fruits": 1, "page": 4}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(d)
	fmt.Println("data:", buf.String())

	res2 := Resp2{}
	_ = json.NewDecoder(strings.NewReader(str)).Decode(&res2)
	fmt.Printf("res2:%+v\n", res2)
}
