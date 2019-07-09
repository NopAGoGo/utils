package naryconv

func Decimal2XnaryReverse(seed, x int) []int {
	r := make([]int, 0)
	mod := 0
	for {
		seed, mod = seed/x, seed%x
		r = append(r, mod)
		if seed == 0 {
			break
		}
	}
	return r
}
