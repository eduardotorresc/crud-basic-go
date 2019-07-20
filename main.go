package main

import (
	"fmt"

	"./models"
)

func main() {
	products, erro := models.GetAll()
	if erro != nil {
		fmt.Println("Ha ocurrido un error")
	}

	fmt.Println("Lista de los productos:\n")
	fmt.Println(products)
}
