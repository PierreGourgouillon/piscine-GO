package student

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
		nombre += lettreNumber * RecursivePower(tailleBase, compteur)
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

func verifBase(base string) bool { //Je vérif si ma base est bonne, GOOD
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

// FINI
