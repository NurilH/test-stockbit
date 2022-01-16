package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 3 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 2 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound])
			}
			return ""
		}
		return ""
	}
	return ""
}

func main() {
	fmt.Println(findFirstStringInBracket(")()coba ya itu hasilnya"))
}
