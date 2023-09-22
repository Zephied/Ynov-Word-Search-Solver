package Correction

import "github.com/01-edu/z01"

func Solver(field [10][10]rune) {
	PrintScramble(field)
}

func PrintScramble(field [10][10]rune) {
	row := 0
	grid := 0
	for row < 10 {
		grid = 0
		for grid < 10 {
			z01.PrintRune(field[row][grid])
		}
	}
}
