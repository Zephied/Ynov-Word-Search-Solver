package Correction

// pour chaque diagonale de bas en haut

func CompareDiagonalDown(field [10][10]rune, wordList []string) []string { // fonction pour comparer les diagonales de bas en haut
	wordFound := []string{}         // liste des mots trouvés
	for _, word := range wordList { // parcours de tous les mots
		for row := 0; row < 10; row++ { // parcours de toutes les lignes
			for grid := 0; grid < 10; grid++ { // parcours de toutes les colonnes
				if rune(word[0]) == field[row][grid] { // si la première lettre du mot correspond à la première lettre du champ
					char := 1                                                                                                               // nombre de lettres trouvées
					for char < len(word) && row-char > 0 && grid+char < len(field[row]) && rune(word[char]) == field[row-char][grid+char] { // parcours de toutes les lettres du mot
						char++ // incrémentation du nombre de lettres trouvées
					}
					if char == len(word) { // si le nombre de lettres trouvées est égal au nombre de lettres du mot
						wordFound = append(wordFound, word) // ajout du mot trouvé dans la liste des mots trouvés
					}
				}
			}
		}
	}
	if len(wordFound) == 0 { // si aucun mot n'a été trouvé
		wordFound = append(wordFound, "Aucun mot de la liste n'a été trouvé") // ajout d'un message dans la liste des mots trouvés
	}
	return wordFound // retour de la liste des mots trouvés
}
