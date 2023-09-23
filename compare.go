package Correction

import (
	"github.com/01-edu/z01"
)

func Compare(field [10][10]rune, wordlist []string) {
	PrintString("Mots trouvés =")
	PrintWords(CompareRow(field, wordlist))
	PrintWords(CompareGrid(field, wordlist))
	PrintWords(CompareDiagonalUp(field, wordlist))
	PrintWords(CompareDiagonalDown(field, wordlist))
	z01.PrintRune('\n')
}
