package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	String1 := args[0]
	for i, s := range String1 {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			PrintStr("0")
			return
		}
	}

	String2 := args[2]
	for i, s := range String2 {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			PrintStr("0")
			return
		}
	}

	signe := args[1]
	if len(signe) != 1 {
		PrintStr("0")
		return
	}

	var nombre1 int64
	var checkbool bool
	var checkbool2 bool
	var nombre2 int64
	var resultat int64
	nombre1, checkbool = Atoi(String1)
	nombre2, checkbool2 = Atoi(String2)

	if checkbool == false || checkbool2 == false {
		PrintStr("0")
		return
	}

	if signe == "+" || signe == "-" || signe == "*" || signe == "/" || signe == "%" {
		switch signe {
		case "+":
			if nombre1 > 0 && nombre2 > 0 && nombre1+nombre2 < 0 { //OVERFLOW ADDITION
				PrintStr("0")
				return
			} else if nombre1 < 0 && nombre2 < 0 && nombre1+nombre2 > 0 { //OVERFLOW ADDITION
				PrintStr("0")
				return
			}
			resultat = nombre1 + nombre2
		case "-":
			if nombre1 > 0 && nombre2 > 0 && (nombre1-nombre2) > nombre1 {
				PrintStr("0")
				return
			}
			resultat = nombre1 - nombre2
		case "*":
			if nombre1 != (nombre1*nombre2)/nombre2 && nombre2 != 0 {
				PrintStr("0")
				return
			} else if nombre2 != (nombre1*nombre2)/nombre1 {
				PrintStr("0")
				return
			}
			resultat = nombre1 * nombre2
		case "/":
			if nombre2 == 0 {
				PrintStr("No division by 0")
				return
			}
			resultat = nombre1 / nombre2
		case "%":
			if nombre2 == 0 {
				PrintStr("No modulo by 0")
				return
			}
			resultat = nombre1 % nombre2
		}
	} else {
		PrintStr("0")
		return
	}

	PrintStr(Itoa(resultat))
}

func PrintStr(chaine string) {
	for i := range chaine {
		z01.PrintRune(rune(chaine[i]))
	}
	z01.PrintRune('\n')
}

func Atoi(s string) (int64, bool) {
	var res int64 = 0

	for i, val := range s {
		if s[0] >= '0' && s[0] <= '9' || s[0] == '+' {
			var dig int64 = 0

			for j := '1'; j <= val; j++ {
				dig++
			}
			res = res*10 + dig
		} else if s[0] == '-' {
			var dig int64 = 0
			for j := '1'; j <= val; j++ {
				dig--
			}
			res = res*10 + dig
		}

		if i > 0 {
			if !(val >= '0' && val <= '9') {
				res = 0
				return res, false
			}
		}
	}

	if res > 0 && s[0] == '-' {
		return 0, false
	} else if res < 0 && s[0] >= '0' && s[0] <= '9' || s[0] == '+' {
		return 0, false
	}

	return res, true
}

func Itoa(n int64) string {

	neg := false
	s := ""

	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		s = string(rune(n%10+48)) + s
		n /= 10
	}

	if neg {
		s = "-" + s
	}

	return s
}

//FINI
