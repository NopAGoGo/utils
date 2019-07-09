package rnd

import (
	"fmt"
	"testing"
)

func BenchmarkIntSliceNoDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSliceNoDuplicate(10000)
	}
}

func BenchmarkIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSlice(100, 2)
	}
}

func TestIntSlice(t *testing.T) {
	fmt.Println(IntSlice(100, 2))
}
