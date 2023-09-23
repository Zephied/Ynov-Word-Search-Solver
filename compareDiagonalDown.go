package Correction

// pour chaque diagonale de bas en haut
func CompareDiagonalDown(field [10][10]rune, wordList []string) []string {
	wordFound := []string{}
	for _, word := range wordList {
		for row := 0; row < 10; row++ {
			for grid := 0; grid < 10; grid++ {
				if rune(word[0]) == field[row][grid] {
					char := 1
					for char < len(word) && row-char > 0 && grid+char < len(field[row]) && rune(word[char]) == field[row-char][grid+char] {
						char++
					}
					if char == len(word) {
						wordFound = append(wordFound, word)
					}
				}
			}
		}
	}
	return wordFound
}
