package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[1:]

	if len(argument) < 1 {
		z01.PrintRune('\n')
	}

	tableauArgument := tableauArgument(argument)
	tableauVoyelle := changeVoyelle(argument)

	tableauFinish := moveLetter(tableauVoyelle, tableauArgument)

	printTableau(tableauFinish)

}

func tableauArgument(argument []string) []rune {
	var tableau []rune

	for i := range argument {
		runeArgs := []rune(argument[i])
		for j := range runeArgs {
			tableau = append(tableau, runeArgs[j])
		}
		tableau = append(tableau, ' ')
	}
	return tableau
}

func printTableau(tableau []rune) {
	for i := range tableau {
		z01.PrintRune(tableau[i])
	}
	z01.PrintRune('\n')
}

func PrintStr(argument []string) {

	for i := range argument {
		runeArgs := []rune(argument[i])

		for j := range runeArgs {
			z01.PrintRune(runeArgs[j])
		}
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}

func changeVoyelle(argument []string) []rune {
	var tableau []rune

	for i := range argument {
		runeArgs := []rune(argument[i])

		for j := range runeArgs {
			if runeArgs[j] == 'a' || runeArgs[j] == 'e' || runeArgs[j] == 'i' || runeArgs[j] == 'o' || runeArgs[j] == 'u' || runeArgs[j] == 'A' || runeArgs[j] == 'E' || runeArgs[j] == 'I' || runeArgs[j] == 'O' || runeArgs[j] == 'U' {
				tableau = append(tableau, runeArgs[j])
			}
		}
	}
	return tableau
}

func moveLetter(tableauVoyelle []rune, tableauArgument []rune) []rune {

	for i := len(tableauVoyelle) - 1; i >= 0; i-- {
		for j := range tableauArgument {
			if tableauArgument[j] == 'a' || tableauArgument[j] == 'e' || tableauArgument[j] == 'i' || tableauArgument[j] == 'o' || tableauArgument[j] == 'u' || tableauArgument[j] == 'A' || tableauArgument[j] == 'E' || tableauArgument[j] == 'I' || tableauArgument[j] == 'O' || tableauArgument[j] == 'U' {
				tableauArgument[j] = tableauVoyelle[i]
				i--
			}
		}
	}
	return tableauArgument
}

//FINISH
