package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p *Plant) String() string {
	return fmt.Sprintf("Plant: (id=%v, name=%s, origin=%v)", p.ID, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{ID: 123, Name: "coffee", Origin: []string{"Bresil", "Java"}}
	fmt.Println(coffee)

	out, _ := xml.MarshalIndent(coffee, " ", "")
	fmt.Println(string(out))
	fmt.Println(xml.Header, string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(&p)

	tomato := &Plant{ID: 234, Name: "tomato", Origin: []string{"China", "Europe"}}
	fmt.Println(tomato)

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(xml.Header, string(out))
}
