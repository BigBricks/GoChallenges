package kata

import "strconv"
import "strings"

// import "fmt"

func FizzBuzzCuckooClock(time string) string {
	// your code here
	f := strings.Split(time, ":")
	first := f[0]
	second := f[1]
	firstNum, fn := strconv.Atoi(first)
	secondNum, sn := strconv.Atoi(second)
	if fn != nil || sn != nil {
		return strings.Repeat("Cuckoo", 12)
	}
	if secondNum == 0 {

		if firstNum == 12 || firstNum == 0 {
			st := strings.Repeat("Cuckoo ", 12)
			return st[:len(st)-1]
		}
		total := firstNum % 12
		st := strings.Repeat("Cuckoo ", total)
		return st[:len(st)-1]
	}
	if secondNum == 30 {
		return "Cuckoo"
	}
	if secondNum%5 == 0 && secondNum%3 == 0 {
		return "Fizz Buzz"
	}
	if secondNum%3 == 0 {
		return "Fizz"
	}
	if secondNum%5 == 0 {
		return "Buzz"
	}
	return "tick"

}
