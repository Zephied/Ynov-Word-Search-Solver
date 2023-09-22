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

func Compare(field [10][10]rune, wordList []string) []string { // fonction pour comparer les mots entre le tableau et le fichier
	wordFound := []string{}         // liste des mots trouvés
	for _, word := range wordList { // parcours des mots dans le fichier
		for row := 0; row < 10; row++ { // parcours les lignes du tableau
			for grid := 0; grid < 10; grid++ { // parcours les colonnes du tableau
				if rune(word[0]) == field[row][grid] { // si la première lettre du mot correspond à la première lettre du table
					char := 1                                                                                          // compte le nombre de caractères restants dans le mot
					for char < len(word) && grid+char < len(field[row]) && rune(word[char]) == field[row][grid+char] { // parcours les caractères restants dans le mot
						char++ // compte le nombre de caractères restants dans le mot
					}
					if char == len(word) { // si le nombre de caractères restants dans le mot correspond à la taille du mot
						wordFound = append(wordFound, word) // ajoute le mot trouvé dans la liste des mots trouvés
					}
				}
			}
		}
	}
	return wordFound // retourne la liste des mots trouvés
}

func CompareGrid(field [10][10]rune, wordList []string) []string { // fonction pour comparer les mots entre le tableau et le fichier
	wordFound := []string{}         // liste des mots trouvés
	for _, word := range wordList { // parcours des mots dans le fichier
		for grid := 0; grid < 10; grid++ { // parcours les lignes du tableau
			for row := 0; row < 10; row++ { // parcours les colonnes du tableau
				if rune(word[0]) == field[grid][row] { // si la première lettre du mot correspond à la première lettre du table
					char := 1                                                                                          // compte le nombre de caractères restants dans le mot
					for char < len(word) && row+char < len(field[grid]) && rune(word[char]) == field[grid][row+char] { // parcours les caractères restants dans le mot
						char++ // compte le nombre de caractères restants dans le mot
					}
					if char == len(word) { // si le nombre de caractères restants dans le mot correspond à la taille du mot
						wordFound = append(wordFound, word) // ajoute le mot trouvé dans la liste des mots trouvés
					}
				}
			}
		}
	}
	return wordFound // retourne la liste des mots trouvés
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
