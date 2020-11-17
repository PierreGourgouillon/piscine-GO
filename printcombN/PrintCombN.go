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
}

func combN(s *[]rune, n int) bool {

	if (*s)[len(*s)-1] != 0 {
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
