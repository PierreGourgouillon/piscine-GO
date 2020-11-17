package tools

func Atoi(s string) int {
	res := 0

	for i, val := range s {
		if s[0] >= '0' && s[0] <= '9' || s[0] == '+' {
			dig := 0

			for j := '1'; j <= val; j++ {
				dig++
			}
			res = res*10 + dig
		} else if s[0] == '-' {
			dig := 0
			for j := '1'; j <= val; j++ {
				dig--
			}
			res = res*10 + dig
		}

		if i > 0 {
			if !(val >= '0' && val <= '9') {
				res = 0
				return res
			}
		}
	}
	return res
}
