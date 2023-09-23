package Correction

import (
	"github.com/01-edu/z01"
)

func PrintString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
