package Correction

import (
	"github.com/01-edu/z01"
)

func Compare(field [10][10]rune, wordlist []string) { // fonction pour comparer les mots trouvés avec les mots de la liste
	voidString := 0                                                             // variable pour savoir si la liste est vide
	ifFirstWord := 0                                                            // variable pour savoir si c'est le premier mot
	PrintWords(CompareRow(field, wordlist), &voidString, &ifFirstWord)          // envoie dans la fonction d'affichage les mots trouver dans CompareRow
	PrintWords(CompareGrid(field, wordlist), &voidString, &ifFirstWord)         // envoie dans la fonction d'affichage les mots trouver dans CompareGrid
	PrintWords(CompareDiagonalUp(field, wordlist), &voidString, &ifFirstWord)   // envoie dans la fonction d'affichage les mots trouver dans CompareDiagonalUp
	PrintWords(CompareDiagonalDown(field, wordlist), &voidString, &ifFirstWord) // envoie dans la fonction d'affichage les mots trouver dans CompareDiagonalDown
	z01.PrintRune('\n')
}
