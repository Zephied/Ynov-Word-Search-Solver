package Correction

// pour chaque ligne
func CompareRow(field [10][10]rune, wordList []string) []string { // fonction pour comparer les mots entre le tableau et le fichier
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
