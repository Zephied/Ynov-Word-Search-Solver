package Correction

import (
	"bufio"
	"os"

	"github.com/01-edu/z01"
)

func Solver(field [10][10]rune) {
	PrintScramble(field)     // fonction pour afficher le tableau
	wordList := ReadFile()   // fonction pour lire le fichier
	Compare(field, wordList) // fonction pour comparer les mots entre le tableau et le fichier
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

func Compare(field [10][10]rune, wordList []string) bool {
	row := 0
	grid := 0
	elems := 0
	char := 0
	for row < 10 {
		for grid < 10 {
			for {
				if wordList[elems][char] == byte(field[row][grid]) {

				}
			}
		}
	}
	return false
}

func ReadFile() []string {
	file, err := os.Open("words.txt")
	if err != nil {
		error("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() { // lis chaque ligne du fichier et les ajoutent comme élément dans le tableau
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil { //verifie si il n'y a pas d'erreur lors du lecture du fichier
		error("failed reading file")
	}
	return lines
}

func error(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
