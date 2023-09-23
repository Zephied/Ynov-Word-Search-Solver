package Correction

import (
	"github.com/01-edu/z01"
)

func Compare(field [10][10]rune, wordlist []string) { // fonction pour comparer les mots trouvés avec les mots de la liste
	PrintString("Mots trouvés =")                    // envoie le texte a ecrire dans la fonction printstring
	PrintWords(CompareRow(field, wordlist))          // envoie dans la fonction d'affichage les mots trouver dans CompareRow
	PrintWords(CompareGrid(field, wordlist))         // envoie dans la fonction d'affichage les mots trouver dans CompareGrid
	PrintWords(CompareDiagonalUp(field, wordlist))   // envoie dans la fonction d'affichage les mots trouver dans CompareDiagonalUp
	PrintWords(CompareDiagonalDown(field, wordlist)) // envoie dans la fonction d'affichage les mots trouver dans CompareDiagonalDown
	z01.PrintRune('\n')
}
