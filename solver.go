package Correction

func Solver(field [10][10]rune) {
	PrintScramble(field)     // fonction pour afficher le tableau
	wordList := ReadFile()   // fonction pour lire le fichier
	Compare(field, wordList) // fonction pour comparer les mots entre le tableau et le fichier
}
