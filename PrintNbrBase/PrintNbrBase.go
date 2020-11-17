package main

import (
	"fmt"

	tools "../tools"
	"github.com/01-edu/z01"
)

func main() {
	PrintNbrBase(125, "0123456789")
}

func PrintNbrBase(nbr int, base string) {
	if verifBase(base) == false {
		tools.PrintStr("NV")
		z01.PrintRune('\n')
	}
	baseString := len(base)
	moduloBase(nbr, baseString)
}

func moduloBase(nbr, base int) int {
	var modulo int
	pow := 0

	fmt.Println(nbr)
	fmt.Println(base)
	fmt.Println()
	for nbr > 0 {

		modulo += nbr % base * tools.RecursivePower(10, pow)
		pow++
		nbr = nbr / base
	}
	fmt.Println(nbr)
	fmt.Println(modulo)
	return modulo
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
