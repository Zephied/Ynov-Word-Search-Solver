package Correction

// pour chaque colonne

func CompareGrid(field [10][10]rune, wordList []string) []string { // fonction pour comparer les mots dans chaque colonne
	wordFound := []string{}         // liste des mots trouvés
	for _, word := range wordList { // parcours des mots
		for row := 0; row < 10; row++ { // parcours chaque ligne
			for grid := 0; grid < 10; grid++ { // parcours chaque colonne
				if rune(word[0]) == field[row][grid] { // si la première lettre du mot correspond à la première lettre du champ
					char := 1                                                                                         // nombre de lettres trouvées
					for char < len(word) && row+char < len(field[row]) && rune(word[char]) == field[row+char][grid] { // parcours les mots
						char++ // incrémentation du nombre de lettres trouvées
					}
					if char == len(word) { // si le nombre de lettres trouvées est égal au nombre de lettres du mot
						wordFound = append(wordFound, word) // ajout du mot dans la liste des mots trouvés
					}
				}
			}
		}
	}
	if len(wordFound) == 0 { // si aucun mot n'a été trouvé
		wordFound = append(wordFound, "Aucun mot de la liste n'a été trouvé") // ajout d'un message dans la liste des mots trouvés
	}
	return wordFound // retourne la liste des mots trouvés
}
