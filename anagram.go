package main

import (
	"fmt"
)

type AnagramData struct {
	Chars          map[rune]bool
	IndexFromArray int
}

func anagram(params []string) [][]string {
	output := [][]string{}
	uniqueWords := []string{}

	mapData := make(map[string]AnagramData)

	for _, value := range params {
		isAnagramFound := false

		for key, anagramData := range mapData {
			if len(key) != len(value) {
				continue
			}

			isAnagramFound = true
			for _, runeVal := range value {
				if _, exist := anagramData.Chars[runeVal]; !exist {
					isAnagramFound = false
					break
				}
			}

			if isAnagramFound {
				words := output[anagramData.IndexFromArray]
				words = append(words, value)

				output[anagramData.IndexFromArray] = words
				break
			}
		}

		if !isAnagramFound {
			uniqueWords = append(uniqueWords, value)

			anagramData := AnagramData{}

			charsData := make(map[rune]bool)
			for _, runeVal := range value {
				charsData[runeVal] = true
			}
			anagramData.Chars = charsData
			anagramData.IndexFromArray = len(uniqueWords) - 1
			output = append(output, []string{value})

			mapData[value] = anagramData
		}
	}

	return output
}

func main() {
	fmt.Println("Hello, playground")
	stringParams := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	// fmt.Printf("%s", findBracket("adasdadasdda()"))
	fmt.Println(anagram(stringParams))
}
