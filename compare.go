package Correction

import (
	"github.com/01-edu/z01"
)

func Compare(field [10][10]rune, wordlist []string) { // fonction pour comparer les mots trouvés avec les mots de la liste
	PrintString("Mots trouvés =")                    // affiche les mots trouvés
	PrintWords(CompareRow(field, wordlist))          // affiche les mots trouvés
	PrintWords(CompareGrid(field, wordlist))         // affiche les mots trouvés
	PrintWords(CompareDiagonalUp(field, wordlist))   // affiche les mots trouvés
	PrintWords(CompareDiagonalDown(field, wordlist)) // affiche les mots trouvés
	z01.PrintRune('\n')
}
