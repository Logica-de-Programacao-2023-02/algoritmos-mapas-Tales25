//Escreva uma função que receba um mapa com valores inteiros e retorne a soma de todos os valores.

package main

import "fmt"

func main() {
	mapFruits := map[string]int{
		"banana":  20,
		"uva":     100,
		"maçã":    15,
		"morango": 30,
	}

	fruitsQuantity := fruitsTotalQuantity(mapFruits)
	fmt.Println(mapFruits)
	fmt.Print("A quantidade total de frutas é: ", fruitsQuantity)
}

func fruitsTotalQuantity(mapFruits map[string]int) int {
	var fruitsTotalQuantity int

	for _, fruitsQuantity := range mapFruits {
		fruitsTotalQuantity += fruitsQuantity
	}

	return fruitsTotalQuantity
}
