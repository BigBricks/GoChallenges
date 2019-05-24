package kata

import "strconv"

func PrinterError(s string) string {
	count := 0
	for x := range s {
		if s[x] > 109 {
			count += 1
		}
	}

	return "" + strconv.Itoa(count) + "/" + strconv.Itoa(len(s))
}
