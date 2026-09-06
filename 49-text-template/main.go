package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("this is value: {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 150)
	t1.Execute(os.Stdout, []string{"go", "python", "c++"})

	Create := func(name, txt string) *template.Template {
		return template.Must(template.New(name).Parse(txt))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct{ Name string }{"John"})

	t2.Execute(os.Stdout, map[string]string{"Name": "Alice"})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")

	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"go", "python", "typescript"})
}
