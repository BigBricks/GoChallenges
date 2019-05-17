package kata

func MxDifLg(a1 []string, a2 []string) int {
	// your code
	maxA := 0
	minA := 10000
	maxB := 0
	minB := 10000
	if len(a1) <= 0 || len(a2) <= 0 {
		return -1
	}
	for i := 0; i < len(a1); i++ {
		if maxA < len(a1[i]) {
			maxA = len(a1[i])
		}
		if minA > len(a1[i]) {
			minA = len(a1[i])
		}
	}
	for i := 0; i < len(a2); i++ {
		if maxB < len(a2[i]) {
			maxB = len(a2[i])
		}
		if minB > len(a2[i]) {
			minB = len(a2[i])
		}
	}
	amax := maxA - minB
	bmax := maxB - minA
	max := 0
	if bmax > amax {
		max = bmax
	} else {
		max = amax
	}

	return max
}
