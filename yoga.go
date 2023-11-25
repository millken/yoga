package yoga

import (
	"fmt"
	"math"
	"os"
)

type YGSize struct {
	width  float32
	height float32
}

func DefaultLogger(config *Config,
	node *Node,
	level YGLogLevel,
	format string,
	args ...interface{}) int {
	switch level {
	case YGLogLevelError, YGLogLevelFatal:
		n, _ := fmt.Fprintf(os.Stderr, format, args...)
		return n
	case YGLogLevelWarn, YGLogLevelInfo, YGLogLevelDebug, YGLogLevelVerbose:
		fallthrough
	default:
		n, _ := fmt.Printf(format, args...)
		return n
	}
}

const (
	uvnan = 0x7FC00001
)

var (
	NaN                 = math.Float32frombits(uvnan)
	YGUndefined float32 = NaN
)

// IsNaN reports whether f is an IEEE 754 “not-a-number” value.
func IsNaN(f float32) (is bool) {
	return f != f
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

var (
	YGValueZero      = YGValue{0, YGUnitPoint}
	YGValueUndefined = YGValue{YGUndefined, YGUnitUndefined}
	YGValueAuto      = YGValue{YGUndefined, YGUnitAuto}
)

type YGValue struct {
	value float32
	unit  YGUnit
}

func (v YGValue) isUndefined() bool {
	return v.unit == YGUnitUndefined
}

func (v YGValue) equal(other YGValue) bool {
	if v.unit != other.unit {
		return false
	}
	switch v.unit {
	case YGUnitUndefined, YGUnitAuto:
		return true
	case YGUnitPoint, YGUnitPercent:
		return v.value == other.value
	}
	return false
}

func (v *YGValue) notEqual(other YGValue) bool {
	return !v.equal(other)
}

func resolveValue(value YGValue, ownerSize float32) FloatOptional {
	switch value.unit {
	case YGUnitPoint:
		return NewFloatOptional(value.value)
	case YGUnitPercent:
		return NewFloatOptional(value.value * ownerSize * 0.01)
	default:
		return undefinedFloatOptional
	}
}

func resolveCompactValue(value CompactValue, ownerSize float32) FloatOptional {
	return resolveValue(value.YGValue(), ownerSize)
}
