package main

import Correction "scrabble"

func main() {
	field := [10][10]rune{
		{'s', 'l', 'c', 'h', 'e', 'v', 'r', 'e', 's', 'r'},
		{'a', 'u', 'c', 'e', 'h', 'm', 'd', 'o', 'i', 'e'},
		{'n', 'n', 'l', 'f', 'i', 'e', 'h', 'r', 'n', 'n'},
		{'d', 'e', 'ô', 't', 'h', 'n', 'r', 'm', 'h', 'c'},
		{'a', 't', 't', 't', 'p', 'u', 't', 'n', 's', 'l'},
		{'l', 't', 'u', 'p', 'o', 'i', 't', 'n', 's', 'l'},
		{'e', 'e', 'r', 'n', 'w', 'f', 'q', 'h', 'r', 's'},
		{'s', 's', 'e', 'e', 'c', 't', 'y', 'e', 'y', 'e'},
		{'c', 'a', 'p', 't', 'i', 'v', 'i', 't', 'é', 'i'},
		{'c', 'a', 's', 'q', 'u', 'e', 't', 't', 'e', 'r'},
	}
	Correction.Solver(field)
}
