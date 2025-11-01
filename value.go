package yoga

/*
#include "cgo_wrapper.h"
*/
import "C"

// Value represents the Yoga style value
type Value struct {
	Value float32
	Unit  Unit
}

func (v Value) IsUndefined() bool {
	return v.Unit == UnitUndefined
}

func (v Value) IsAuto() bool {
	return v.Unit == UnitAuto
}

func (v Value) Equal(other Value) bool {
	if v.Unit != other.Unit {
		return false
	}
	switch v.Unit {
	case UnitUndefined, UnitAuto:
		return true
	case UnitPoint, UnitPercent:
		return v.Value == other.Value
	}
	return false
}

func (v *Value) NotEqual(other Value) bool {
	return !v.Equal(other)
}

// valueFromYGValue converts C's YGValue to Go's Value
func valueFromYGValue(v C.YGValue) Value {
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}
