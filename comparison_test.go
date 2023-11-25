package yoga

import "testing"

func TestInexactEquals(t *testing.T) {
	tests := []struct {
		a, b   any
		expect bool
	}{
		{float32(1), float32(1), true},
		{float32(1), float32(2), false},
		{float32(1), float64(1.00011), false},
		{float64(1), float64(1), true},
		{float64(1), float64(2), false},
		{float64(1), float64(1.0001), true},
		{float64(1), float64(1.00011), false},
		{YGValue{1, YGUnitPoint}, YGValue{1, YGUnitPoint}, true},
		{YGValue{1, YGUnitPoint}, YGValue{2, YGUnitPoint}, false},
		{YGValue{1, YGUnitPoint}, YGValue{1.00011, YGUnitPoint}, false},
		{CompactValueOfZero(), CompactValueOfZero(), true},
		{CompactValueOfZero(), CompactValueOfAuto(), false},
		{FloatOptional{1}, FloatOptional{1}, true},
		{FloatOptional{1}, FloatOptional{2}, false},
		{FloatOptional{1}, FloatOptional{1.00011}, false},
		{[]CompactValue{CompactValueOfZero()}, []CompactValue{CompactValueOfZero()}, true},
		{[]CompactValue{CompactValueOfZero()}, []CompactValue{CompactValueOfAuto()}, false},
		{[2]CompactValue{CompactValueOfZero(), CompactValueOfZero()}, [2]CompactValue{CompactValueOfZero(), CompactValueOfZero()}, false},
	}
	for _, test := range tests {
		if inexactEquals(test.a, test.b) != test.expect {
			t.Errorf("inexactEquals(%v, %v) != %v", test.a, test.b, test.expect)
		}
	}
}
