package main

import (
	"fmt"
)

func findBracket(str string) string {
	if len(str) > 0 {
		var runeBracketOpen, runeBracketClose rune = '(', ')'
		var idxOpen, idxClose int = -1, -1
		for idxData, runeVal := range str {
			if runeVal == runeBracketOpen {
				idxOpen = idxData
			}
			if runeVal == runeBracketClose {
				idxClose = idxData

				if idxOpen != -1 {
					runes := []rune(str)
					if idxClose-idxOpen == 1 {
						break
					}
					return string(runes[idxOpen+1 : idxClose-1])
				} else {
					break
				}
			}
		}
	}
	return ""
}

func main() {
	fmt.Printf("%s", findBracket("adasdadasdda()"))
}
