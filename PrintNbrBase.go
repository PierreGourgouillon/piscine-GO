package student

import (
	tools "../tools"
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if verifBase(base) == false {
		tools.PrintStr("NV")
		z01.PrintRune('\n')
		return
	}
	negatif := isNegative(nbr)
	baseString := len(base)
	tableau := NbrBase(nbr, baseString, base, negatif)
	printTableau(tableau, negatif)
}

func NbrBase(nbr, base int, String string, negatif bool) []byte {
	var modulo int
	var tableau []byte

	if negatif == true {
		nbr *= (-1)
	}

	for nbr > 0 {

		modulo = nbr % base
		nbr = nbr / base
		tableau = append(tableau, intInBase(modulo, String, negatif))

	}
	return tableau
}

func intInBase(nombre int, base string, negatif bool) byte {

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

func verifBase(base string) bool {
	var garage byte
	var compteur int
	var COMPTEUR int

	if len(base) < 2 {
		return false
	}

	for k := range base {
		if base[k] == '+' || base[k] == '-' {
			return false
		}
	}

	for i := range base {
		garage = base[i]
		for j := range base {
			COMPTEUR++
			if garage == base[j] {
				compteur++

			}

			if compteur >= 2 {
				return false
			}

			if COMPTEUR == len(base) {
				COMPTEUR = 0
				compteur = 0
			}
		}
	}
	return true
}

//FINI
