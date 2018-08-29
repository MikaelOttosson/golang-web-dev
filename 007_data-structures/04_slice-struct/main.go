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

	r2 := resturant{
		Place: "MCD",
		Food:  "Burgers",
	}

	r3 := resturant{
		Place: "PizzaHut",
		Food:  "Pizza",
	}

	resturants := []resturant{r1, r2, r3}
	err := tpl.Execute(os.Stdout, resturants)
	if err != nil {
		log.Fatalln(err)
	}
}
