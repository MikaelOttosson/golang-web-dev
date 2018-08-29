package main

import "fmt"

func main() {
	name := "Mikael Ottosson"

	tp1 := `
  <!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello Google!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
  `
	fmt.Println(tp1)

}
