package Correction

import (
	"github.com/01-edu/z01"
)

func PrintString(s string) { // fonction pour afficher un string
	for _, v := range s { // boucle sur les caractères du string
		z01.PrintRune(v) // affiche un caractère
	}
	z01.PrintRune('\n') // affiche un retour de ligne
}
