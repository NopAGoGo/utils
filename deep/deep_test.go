package deep

import "testing"

func TestClone(t *testing.T) {
	c := 1
	d := 2
	for _, unit := range []struct {
		a        *int
		b        *int
		expected int
	}{
		{&d, &c, 0},
	} {
		if uselessCount := Copy(unit.b, unit.a); true {
			t.Errorf("expected: [%v], uselessCount: [%v]", unit, uselessCount)
		}
	}
}
