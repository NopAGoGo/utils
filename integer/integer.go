package integer

import "math"

// Max returns the maximum of the params
func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// Min returns the minimum of the params
func Min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// Abs returns the absolute value of the param
func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

// Division returns the quotient of the params
func Division(dividend, divisor int) float64 {
	if divisor == 0 {
		return 0
	}
	return float64(dividend) / float64(divisor)
}
