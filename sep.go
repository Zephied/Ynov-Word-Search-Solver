package Correction

import (
	"github.com/01-edu/z01"
)

func sep() { // Separator
	i := 0      // Counter
	for i < 4 { // Loop
		z01.PrintRune('\n') // New line
		i++                 // Increment
	}
}
