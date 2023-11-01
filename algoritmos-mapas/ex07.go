//Escreva uma função que receba uma string contendo uma frase e retorne um mapa onde as chaves são as palavras
//encontradas na frase e os valores são mapas contendo a contagem de cada letra em cada palavra.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	wordsCharMap := createWordsCharMap(str)
	fmt.Print(wordsCharMap)
}

// retorna um mapa que tem uma string como chave e um mapa (com chave string e valor int) como valor
func createWordsCharMap(str string) map[string]map[string]int {
	//para cirar um slice onde cada palavra é uma string (como strings.Split(str, " "))
	sliceStr := strings.Fields(str)

	wordsCharMap := make(map[string]map[string]int)

	for _, words := range sliceStr {
		charMap := make(map[string]int)

		for _, valor := range words {
			charMap[string(valor)]++
			wordsCharMap[words] = charMap
		}
	}

	return wordsCharMap
}

//CRIAR UM MAPA ONDE AS CHAVES SÃO AS PALAVRAS DE UMA STR ------------------- OK
//OS VALORES DESSE MAPA DEVEM SER MAPAS QUE CONTAM CADA CHAR DA PALAVRA ----- OK
