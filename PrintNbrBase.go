package student

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if verifBase1(base) == false {
		PrintStr("NV")
		z01.PrintRune('\n')
		return
	}
	negatif := isNegative(nbr)
	baseString := len(base)
	tableau := NbrBase1(nbr, baseString, base, negatif)
	printTableau(tableau, negatif)
}

func PrintStr(chaine string) {
	for i := range chaine {
		z01.PrintRune(rune(chaine[i]))
	}
}

func NbrBase1(nbr, base int, String string, negatif bool) []byte {
	var modulo int
	var tableau []byte

	if negatif == true {
		nbr *= (-1)
	}

	for nbr > 0 {

		modulo = nbr % base
		nbr = nbr / base
		tableau = append(tableau, intInBase1(modulo, String, negatif))

	}
	return tableau
}

func intInBase1(nombre int, base string, negatif bool) byte {

	if negatif == false {
		for i := range base {

			if i == nombre {
				return base[i]
			}

		}
	} else {
		for i := len(base) - 1; i >= 0; i-- {
			if i == nombre {
				return base[i]
			}
		}
	}

	return 'Z'
}

func printTableau(tableau []byte, negatif bool) {

	if negatif == true {
		z01.PrintRune('-')
		for i := len(tableau) - 1; i >= 0; i-- {
			z01.PrintRune(rune(tableau[i]))
		}
	} else {
		for i := len(tableau) - 1; i >= 0; i-- {
			z01.PrintRune(rune(tableau[i]))
		}
	}

	z01.PrintRune('\n')
}

func isNegative(base int) bool {

	if base < 0 {
		return true
	}
	return false
}

//FINI
