//Escreva uma função que conte a ocorrência de cada palavra em uma string e retorne um mapa onde as chaves são as
//palavras encontradas e os valores são o número de ocorrências de cada palavra.

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

	mapStr := charQuantity(str)
	fmt.Print(mapStr)
}

func charQuantity(str string) map[string]int {
	sliceStr := strings.Split(str, "")
	mapStr := make(map[string]int)

	//loop que itera sobre um slice (normal)
	for _, char := range sliceStr {
		//utilizando o caractere como chave, adicionamos um ao valor cada vez que um char se repete
		mapStr[char]++
	}

	return mapStr
}
