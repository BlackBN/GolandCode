package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {

	funcMap := template.FuncMap{
		"strupper": upper,
	}
	t := template.New("test1")
	tmpl, err := t.Funcs(funcMap).Parse(`{{strupper .}}`)
	if err != nil {
		panic(err)
	}
	_ = tmpl.Execute(os.Stdout, "go programming")
}

func upper(str string) string {
	return strings.ToUpper(str)
}
