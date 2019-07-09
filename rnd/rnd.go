package rnd

import (
	"math/rand"
	"time"
)

func IntSlice(l, m int) []int {
	rand.Seed(time.Now().UnixNano())
	r := make([]int, l)
	for i := range r {
		r[i] = rand.Int() % m
	}
	return r
}

func IntSliceNoDuplicate(l int) []int {
	rand.Seed(time.Now().UnixNano())
	p := make([]int, l)
	r := make([]int, l)
	for i := range p {
		sign := rand.Int() - rand.Int()
		if sign < 0 {
			p[i] = -i - 1
		} else {
			p[i] = i + 1
		}

	}
	for ir := range r {
		ip := rand.Int() % len(p)
		r[ir] = p[ip]
		p = append(p[:ip], p[ip+1:]...)
	}
	return r
}
