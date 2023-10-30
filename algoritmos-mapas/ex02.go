//Escreva uma função que receba dois mapas e retorne um novo mapa contendo todos os elementos dos mapas de entrada.
//Em caso de chaves duplicadas, o valor do segundo mapa deve prevalecer.

package main

import "fmt"

func main() {
	map1 := map[string]int{
		"A": 10,
		"X": 10,
		"B": 10,
	}

	map2 := map[string]int{
		"C": 20,
		"X": 20,
		"D": 20,
	}

	concatenatedMaps := concatenateMaps(map1, map2)
	for chave, valor := range concatenatedMaps {
		fmt.Printf("%s: %d\n", chave, valor)
	}
}

func concatenateMaps(map1, map2 map[string]int) map[string]int {
	concatenatedMaps := make(map[string]int)

	for chave, valor := range map1 {
		concatenatedMaps[chave] = valor
	}

	for chave, valor := range map2 {
		//assim, caso a chave do segundo mapa for igual a do primeiro, irá substituir o valor pelo o do segundo
		concatenatedMaps[chave] = valor
	}

	return concatenatedMaps
}
