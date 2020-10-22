package student

func Atoi(s string) int {
	var num int
	var pow int = 1

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '+' {
			if i != 0 {
				return 0
			}
			continue
		}

		if s[i] == '-' {
			if i != 0 {
				return 0
			}
			num *= -1
			return num
		}

		if s[i] >= '0' && s[i] <= '9' {
			num += int(s[i]-'0') * pow
			pow *= 10
		} else {
			return 0
		}
	}
	return num
}
