package Correction

import "github.com/01-edu/z01"

func Solver(field [10][10]rune) {
	PrintScramble(field)
}

func PrintScramble(field [10][10]rune) {
	row := 0
	grid := 0
	sep()
	for row < 10 {
		grid = 0
		for grid < 10 {
			if grid == 0 {
				z01.PrintRune('[')
			}
			z01.PrintRune(field[row][grid])
			if grid < 9 {
				z01.PrintRune(' ')
			}
			grid++
		}
		row++
		z01.PrintRune(']')
		z01.PrintRune('\n')
	}
	sep()
}

func sep() {
	i := 0
	for i < 4 {
		z01.PrintRune('\n')
		i++
	}
}
