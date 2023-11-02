//Escreva uma função que receba um slice de inteiros e retorne um mapa onde as chaves são os pares de números
//encontrados no slice e os valores são as frequências de cada par.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	sliceInt := []int{1, 2, 2, 3, 4, 4}

	pairsMap := createPairsMap(sliceInt)
	fmt.Print(pairsMap)
}

func createPairsMap(sliceInt []int) map[string]int {
	var (
		x, y              int
		xStr, yStr, chave string
	)
	pairsMap := make(map[string]int)

	for i, num := range sliceInt {
		for j, pair := range sliceInt {
			if i != j {
				x, y = num, pair

				xStr = strconv.Itoa(x)
				yStr = strconv.Itoa(y)

				chave = "(" + xStr + "," + yStr + ")"
				pairsMap[chave]++
			}
		}
	}

	return pairsMap
}
