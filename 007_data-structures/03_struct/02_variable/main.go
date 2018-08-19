package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type resturant struct {
	Place string
	Food  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	r1 := resturant{
		Place: "BurgerKing",
		Food:  "Fries",
	}

	err := tpl.Execute(os.Stdout, r1)
	if err != nil {
		log.Fatalln(err)
	}
}
