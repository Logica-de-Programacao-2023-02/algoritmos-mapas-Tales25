package main

import "fmt"

func main() {
	//var newSliceInt []int
	mapSliceInt := make(map[int]int)
	sliceInt := []int{1, 1, 2, 3, 5, 6, 6, 7, 8, 8, 9, 9, 10, 10, 11}

	for _, num /*recebe cada numero do slice, 1 por vez*/ := range sliceInt {
		mapSliceInt[num]++ //para a chave [num] adicione 1 ao valor (inteiro)...
		//...(pela quantidade de vezes que cada chave aparece no slice)
		fmt.Println(mapSliceInt[num])
	}

	fmt.Println(mapSliceInt)
}
