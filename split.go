package student

func Split(s, sep string) []string {
	var tableau []string
	longSep := len(sep)
	longS := len(s)
	index := 0

	for i := 0; i <= longS-longSep; i++ {
		if sep == s[i:i+longSep] {
			tableau = append(tableau, s[index:i])
			index = i + longSep
		}
		if i == longS-longSep {
			tableau = append(tableau, s[index:longS])
		}
	}
	return tableau
}

//FINI
