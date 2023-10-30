//Escreva uma função que receba um slice de inteiros e retorne um novo slice contendo
//apenas os valores únicos, ou seja, sem duplicatas. Utilize um mapa para auxiliar na
//remoção das duplicatas.

package main

import "fmt"

func main() {
	sliceInt := []int{1, 1, 2, 3, 5, 6, 6, 7, 8, 8, 9, 9, 10, 10, 11}
	fmt.Println(sliceInt)

	newSliceInt := removeDuplicates(sliceInt)
	fmt.Println(newSliceInt)
}

func removeDuplicates(sliceInt []int) []int {
	var newSliceInt []int
	mapSliceInt := make(map[int]int)

	for _, valor := range sliceInt {
		mapSliceInt[valor]++
	}

	//por algum motivo eu não itero no mapa, e sim no próprio slice
	for _, valor := range sliceInt {
		//se o valor da chave "num" for igual a 1
		if mapSliceInt[valor] == 1 {
			//adicione o num ao slice
			newSliceInt = append(newSliceInt, valor)
		}
	}
	//OBS.: O NUMERO É O MESMO QUE A CHAVE NUM
	return newSliceInt
}
