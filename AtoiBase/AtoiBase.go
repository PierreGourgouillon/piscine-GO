package main

func main() {
	AtoiBase("125", "0123456789")
}

func AtoiBase(s string, base string) int {
	if verifBase(base) == false {
		return 0
	}
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

/*********************************************A FINIR*********************************************/
