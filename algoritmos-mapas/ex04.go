//Escreva uma função que receba um slice de palavras e retorne um mapa onde as chaves são palavras que
//são anagramas entre si e os valores são os grupos de palavras anagramas.

package main

import "fmt"

func main() {
	sliceStr := []string{"amor", "roma", "Hello,", "mora", "abc", "World!", "cab", "bca"}

	mapAnagram := findAnagrams(sliceStr)
	fmt.Print(mapAnagram)
}

func findAnagrams(sliceStr []string) map[string][]string {
	anagramMap := make(map[string][]string)

	for i, word1 := range sliceStr {
		for j, word2 := range sliceStr {
			//se forem palavras diferentes (indices diferentes) e areAnagramas == true
			if i != j && areAnagrams(word1, word2) {
				anagramMap[word1] = append(anagramMap[word1], word2)
			}
		}
	}

	return anagramMap
}

func areAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	//mapas
	charCount1Map := countChars(str1)
	charCount2Map := countChars(str2)

	//conferindo se são anagramas
	for chave, _ := range charCount1Map {
		if charCount1Map[chave] != charCount2Map[chave] {
			return false
		}
	}

	return true
}

// com essa função só preciso de um mapa e passa-los para strings diferentes
func countChars(str string) map[string]int {
	charCount := make(map[string]int)

	for _, char := range str {
		charCount[string(char)]++
	}

	return charCount
}
