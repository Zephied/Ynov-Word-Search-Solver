package Correction

import (
	"bufio"
	"os"
)

func ReadFile() []string {
	file, err := os.Open("words.txt")
	if err != nil {
		PrintString("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() { // lis chaque ligne du fichier et les ajoutent comme élément dans le tableau
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil { //verifie si il n'y a pas d'erreur lors du lecture du fichier
		PrintString("failed reading file")
	}
	return lines
}
