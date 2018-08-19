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
	mylist := []string{"Burgers", "Pizza", "Fries", "Meatballs"}

	err := tpl.Execute(os.Stdout, mylist)
	if err != nil {
		log.Fatalln(err)
	}

}
