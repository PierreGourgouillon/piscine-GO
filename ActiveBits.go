package student

func ActiveBits(n int) uint {
	var compteur uint

	for n != 0 {

		if n >= 256 {
			n = n - 256
			compteur++

		} else if n >= 128 {
			n = n - 128
			compteur++

		} else if n >= 64 {
			n = n - 64
			compteur++

		} else if n >= 32 {
			n = n - 32
			compteur++

		} else if n >= 16 {
			n = n - 16
			compteur++

		} else if n >= 8 {
			n = n - 8
			compteur++

		} else if n >= 4 {
			n = n - 4
			compteur++

		} else if n >= 2 {
			n = n - 2
			compteur++

		} else if n >= 1 {
			n = n - 1
			compteur++

		}
	}

	return compteur
}

//FINISH
