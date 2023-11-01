//Escreva uma função que gere a sequência de Fibonacci até um determinado número n e retorne um mapa onde as
//chaves são os índices da sequência e os valores são os números correspondentes.

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	fibonacciMap := generateFibonacciMap(num)
	fmt.Print(fibonacciMap)
}

func generateFibonacciMap(num int) map[int]int {
	fibonacciMap := make(map[int]int)
	x, y := 0, 1
	chave := 0

	//enquanto o número inicial (número principal a ser alterado) for menor ou igual ao num
	for x <= num {
		//o valor da chave (inicialmente 0) recebe x
		fibonacciMap[chave] = x

		//x recebe y (seu próximo) e y recebe x+y (proximo depois do y inicial)
		x, y = y, x+y

		//adicionar 1 a chave
		chave++
	}

	return fibonacciMap
}
