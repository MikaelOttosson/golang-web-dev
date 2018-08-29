package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	myFood := map[string]string{
		"MCD":        "Burgers",
		"BurgerKing": "Fries",
		"PizzaHut":   "Pizza",
		"Subway":     "Sandwiches",
	}

	err := tpl.Execute(os.Stdout, myFood)
	if err != nil {
		log.Fatalln(err)
	}
}
