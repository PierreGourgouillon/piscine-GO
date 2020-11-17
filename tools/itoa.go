package tools

func Itoa(n int) string {

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
