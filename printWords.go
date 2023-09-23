package Correction

import (
	"github.com/01-edu/z01"
)

func PrintWords(words []string, voidString *int, ifFirstWord *int) {
	if words[0] != "Aucun mot de la liste n'a été trouvé" {
		if *ifFirstWord == 0 {
			PrintString("Mots trouvés =") // envoie le texte a ecrire dans la fonction printstring
			*ifFirstWord++
		}
		z01.PrintRune('[')
		for _, word := range words {
			for _, c := range word {
				z01.PrintRune(c)
			}
			if word != words[len(words)-1] {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')
	} else {
		*voidString++
	}
	if *voidString == 4 {
		for _, c := range words[0] {
			z01.PrintRune(c)
		}
	}
}
