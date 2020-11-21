package student

import (
	tools "../tools"

	"github.com/01-edu/z01"
)

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := AtoiBase(nbr, baseFrom)
	tableauByte := PrintNbrBase(n, baseTo)
	var NewTableau []byte

	for i := len(tableauByte) - 1; i >= 0; i-- {
		NewTableau = append(NewTableau, tableauByte[i])
	}
	return string(NewTableau)
}

func AtoiBase(s string, base string) int { // Genre de main
	if verifBase(base) == false {
		return 0
	}
	Number := len(base)
	tableauBase := createTableBase(s)
	numberFinal := calculBase(tableauBase, s, base, Number)

	return numberFinal
}

func calculBase(tableau []rune, s, base string, tailleBase int) int { // Convertisseur de mon string en mon int
	compteur := 0
	nombre := 0

	for i := len(tableau) - 1; i >= 0; i-- {

		lettreNumber := checkNumberInBase(rune(tableau[i]), base) //marche correctement
		nombre += lettreNumber * tools.RecursivePower(tailleBase, compteur)
		compteur++
	}

	return nombre
}

func checkNumberInBase(lettre rune, base string) int { //Je check dans string l'index, GOOD
	compteur := 0

	for i := range base {
		if rune(base[i]) == lettre {
			return compteur
		}
		compteur++
	}

	return 100
}

func createTableBase(s string) []rune { //Je crée un tableau pour mon S, GOOD

	tableauBase := make([]rune, len(s))

	for i := range s {
		tableauBase[i] = rune(s[i])
	}

	return tableauBase
}

func PrintNbrBase(nbr int, base string) []byte { //genre de fonction main
	var TABLEAU = []byte{0}

	if verifBase(base) == false {
		tools.PrintStr("NV")
		z01.PrintRune('\n')
		return TABLEAU
	}
	baseString := len(base)
	tableau := NbrBase(nbr, baseString, base)
	return tableau
}

func NbrBase(nbr, base int, String string) []byte {
	var modulo int
	var tableau []byte

	for nbr > 0 {

		modulo = nbr % base
		nbr = nbr / base
		tableau = append(tableau, intInBase(modulo, String))

	}
	return tableau
}

func intInBase(nombre int, base string) byte {

	for i := len(base) - 1; i >= 0; i-- {
		if i == nombre {
			return base[i]
		}
	}

	return 'Z'
}

func verifBase(base string) bool { //Verification de la base
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
