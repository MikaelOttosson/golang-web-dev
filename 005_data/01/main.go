package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// pass in data
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "golang")
	if err != nil {
		log.Fatalln(err)
	}
}
