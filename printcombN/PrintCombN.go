package main

import (
	tools "../tools"
	"github.com/01-edu/z01"
)

func main() {
	n := 4
	PrintCombN(n)
}

func PrintCombN(n int) {
	s := make([]rune, n)
	combN(&s, n)
	z01.PrintRune('\n')
}

func combN(s *[]rune, n int) bool {

	if n == 1 && (*s)[len(*s)-1] == '9' {
		tools.PrintStr(string(*s))
	} else if n == 2 && (*s)[len(*s)-1] == '9' && (*s)[len(*s)-2] == '8' {
		tools.PrintStr(string(*s))
	} else if n == 3 && (*s)[len(*s)-1] == '9' && (*s)[len(*s)-3] == '7' {
		tools.PrintStr(string(*s))
	} else if n == 4 && (*s)[len(*s)-1] == '9' && (*s)[len(*s)-4] == '6' {
		tools.PrintStr(string(*s))
	} else if n == 5 && (*s)[len(*s)-1] == '9' && (*s)[len(*s)-5] == '5' {
		tools.PrintStr(string(*s))
	} else if n == 6 && (*s)[len(*s)-1] == '9' && (*s)[len(*s)-6] == '4' {
		tools.PrintStr(string(*s))
	} else if n == 7 && (*s)[len(*s)-1] == '9' && (*s)[len(*s)-7] == '3' {
		tools.PrintStr(string(*s))
	} else if n == 8 && (*s)[len(*s)-1] == '9' && (*s)[len(*s)-8] == '2' {
		tools.PrintStr(string(*s))
	} else if n == 9 && (*s)[len(*s)-1] == '9' && (*s)[len(*s)-9] == '1' {
		tools.PrintStr(string(*s))
	} else if (*s)[len(*s)-1] != 0 {
		tools.PrintStr(string(*s))
		tools.PrintStr(", ")
		return true
	}

	for i, v := range *s {

		if v != 0 {
			continue
		}

		for k := '0'; k <= '9'; k++ {
			if canPlace(*s, k) && (i == 0 || (*s)[i-1] < k) {
				(*s)[i] = k
				if combN(s, n) {
					(*s)[i] = 0
				}
			}
		}
		(*s)[i] = 0
		return false
	}
	return true
}

func canPlace(s []rune, k rune) bool {
	for _, i := range s {
		if k == i {
			return false
		}
	}
	return true
}

/*********************************************A FINIR*********************************************/
