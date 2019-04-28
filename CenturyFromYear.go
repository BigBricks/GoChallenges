package kata

func century(year int) int {
	// Finish this :)
	if year%100 > 0 {
		return (year / 100) + 1
	}
	return (year / 100)
}
