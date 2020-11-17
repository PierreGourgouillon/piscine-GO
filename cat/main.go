package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printLn(chaine string) {

	for _, r := range chaine {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func fatal(chaine string) {
	printLn(chaine)
	os.Exit(1)
}

func main() {
	compteur := 0

	for range os.Args {
		compteur++
	}

	if compteur == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fatal(err.Error())
		}
		printLn(string(b))
	} else {
		for _, file := range os.Args[1:] {
			a, err := ioutil.ReadFile(file)
			if err != nil {
				fatal("ERROR: " + err.Error())
			}
			printLn(string((a)))
		}
	}
}

//FINISH
