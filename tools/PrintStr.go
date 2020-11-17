package tools

import "github.com/01-edu/z01"

func PrintStr(chaine string) {
	for i := range chaine {
		z01.PrintRune(rune(chaine[i]))
	}
}
