//Escreva uma função que receba uma lista de mapas, onde cada mapa contém a contagem de palavras de um
//texto(uma frase, uma string), e retorne um único mapa contendo a soma de todas as contagens.

package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Essa é a primeira frase"
	str2 := "Essa é a segunda frase"
	str3 := "Essa é a terceira frase"

	countWordsMap1 := countWords(str1)
	countWordsMap2 := countWords(str2)
	countWordsMap3 := countWords(str3)

	sliceMaps := []map[string]int{countWordsMap1, countWordsMap2, countWordsMap3}

	totalSumMap := sumWordsMap(sliceMaps)
	fmt.Print(totalSumMap)
}

func sumWordsMap(sliceMaps []map[string]int) map[string]int {
	totalWords := make(map[string]int)

	for _, maps := range sliceMaps {
		for chave, _ := range maps {
			totalWords[chave]++
		}
	}

	return totalWords
}

func countWords(str string) map[string]int {
	countWordsMap := make(map[string]int)

	//criar um slice de palavras divididas por um espaço
	words := strings.Fields(str)

	for _, word := range words {
		countWordsMap[word]++
	}

	return countWordsMap
}

//CRIAR UMA FUNÇÃO QUE FAÇA O MAPA E CONTE AS PALAVRAS -------------- OK
//CRIAR UM MAPA E IR ARMAZENANDO A SOMA DE TODAS A PALAVRAS NELE ----
