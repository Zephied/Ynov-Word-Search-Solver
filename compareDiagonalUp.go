package Correction

// pour chaque diagonale de haut en bas

func CompareDiagonalUp(field [10][10]rune, wordList []string) []string { // fonction pour comparer les diagonales de haut en bas
	wordFound := []string{}         // liste des mots trouvés
	for _, word := range wordList { // parcours de tous les mots
		for row := 0; row < 10; row++ { // parcours de chaque ligne
			for grid := 0; grid < 10; grid++ { // parcours de chaque colonne
				if rune(word[0]) == field[row][grid] { // si la première lettre du mot correspond à la première lettre du champ
					char := 1                                                                                                                             // compte le nombre de lettres restantes
					for char < len(word) && row+char < len(field[row]) && grid+char < len(field[row]) && rune(word[char]) == field[row+char][grid+char] { // parcours de tous les mots restants
						char++ // parcours de tous les mots restants
					}
					if char == len(word) { // si tous les mots restants ont été trouvés
						wordFound = append(wordFound, word) // ajout du mot dans la liste des mots trouvés
					}
				}
			}
		}
	}
	return wordFound // retour de la liste des mots trouvés
}
