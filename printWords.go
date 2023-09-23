package Correction

import (
	"github.com/01-edu/z01"
)

func PrintWords(words []string) {
	z01.PrintRune('[')
	for _, word := range words {
		for _, c := range word {
			z01.PrintRune(c)
		}
		if word != words[len(words)-1] {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
}
