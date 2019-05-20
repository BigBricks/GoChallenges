package kata

import "strings"
import "sort"

func TwoToOne(s1 string, s2 string) string {
	str := s1 + s2
	j := strings.Split(str, "")
	key := make(map[string]bool)
	list := []string{}
	for _, x := range j {
		if _, l := key[x]; !l {
			key[x] = true
			list = append(list, x)
		}
	}
	sort.Strings(list)
	return strings.Join(list, "")
}
