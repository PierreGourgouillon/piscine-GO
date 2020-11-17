package student

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	garage := ""

	for range a {
		for i := 0; i < len(a)-1; i++ {

			if f(a[i], a[i+1]) == 1 { //a[i]>a[i+1] tu les changes
				garage = a[i]
				a[i] = a[i+1]
				a[i+1] = garage
			}
		}
	}
}

func Compare(a, b string) int {
	if a == b {
		return 0
	}

	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

//FINISH
