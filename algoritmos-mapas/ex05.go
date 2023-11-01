//Escreva uma função que receba uma string e retorne um mapa onde as chaves são os caracteres
//presentes na string e os valores são a frequência de cada caractere.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	charCountMap := countChar(str)

	fmt.Print(charCountMap)
}

func countChar(str string) map[string]int {
	charCountMap := make(map[string]int)

	for _, char := range str {
		charCountMap[string(char)]++
	}

	return charCountMap
}
