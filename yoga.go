package yoga

/*
#cgo CXXFLAGS: -std=c++20
#cgo CPPFLAGS: -I${SRCDIR}/include
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/_libs/darwin/arm64 -lyogacore -lstdc++
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/_libs/linux/amd64 -lyogacore -lstdc++ -lm
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/_libs/windows/amd64 -lyogacore -lstdc++
#include <yoga/Yoga.h>


*/
import "C"
import "math"

const (
	uvnan = 0x7FC00001
)

var (
	NaN               = math.Float32frombits(uvnan)
	Undefined float32 = NaN
	Zero      float32 = 0.0
)

// IsNaN reports whether f is an IEEE 754 "not-a-number" value.
func IsNaN(f float32) (is bool) {
	return f != f
}

// FloatIsUndefined reports whether a float value represents undefined.
func FloatIsUndefined(value float32) bool {
	return IsNaN(value)
}

func IsInf(f float32, sign int) bool {
	// Test for infinity by comparing against maximum float.
	// To avoid the floating-point hardware, could use:
	//	x := Float64bits(f);
	//	return sign >= 0 && x == uvinf || sign <= 0 && x == uvneginf;
	return sign >= 0 && f > math.MaxFloat32 || sign <= 0 && f < -math.MaxFloat32
}

func If[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}
