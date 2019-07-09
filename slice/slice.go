package slice

func Reverse(s []int) []int {
	for h, t := 0, len(s)-1; h < t; h, t = h+1, t-1 {
		s[h], s[t] = s[t], s[h]
	}
	return s
}
