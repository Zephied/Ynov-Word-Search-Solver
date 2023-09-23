package Correction

import (
	"github.com/01-edu/z01"
)

func PrintScramble(field [10][10]rune) { // fonction to print scrambled field
	row := 0       // initialise row à 0
	grid := 0      // initialise grid à 0
	sep()          // print separator
	for row < 10 { // pour chaque ligne
		grid = 0        // initialise grid à 0
		for grid < 10 { // pour chaque colonne
			if grid == 0 { // si la première colonne = 0
				z01.PrintRune('[') // on affiche [
			}
			z01.PrintRune(field[row][grid]) // on affiche le caractère
			if grid < 9 {                   // si la dernière colonne = 9
				z01.PrintRune(' ') // on affiche un espace
			}
			grid++ // on affiche
		}
		row++
		z01.PrintRune(']')  // on affiche ]
		z01.PrintRune('\n') // on affiche
	}
	sep() // print separator
}
