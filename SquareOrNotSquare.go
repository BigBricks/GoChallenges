package kata

import "math"

func SquareOrSquareRoot(arr []int) []int {
	a := make([]int, len(arr))
	for x := 0; x < len(arr); x++ {
		j := math.Sqrt(float64(arr[x]))
		if j == math.Floor(j) {
			a[x] = int(j)
		} else {
			a[x] = int(math.Pow(float64(arr[x]), 2))
		}
	}
	return a
}
