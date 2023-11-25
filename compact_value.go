package yoga

import "math"

type CompactValue struct {
	repr uint32
}

const (
	LOWER_BOUND         float32 = 1.08420217e-19
	UPPER_BOUND_POINT   float32 = 36893485948395847680.0
	UPPER_BOUND_PERCENT float32 = 18446742974197923840.0
	PERCENT_BIT         uint32  = 0x40000000
	AUTO_BITS           uint32  = 0x7faaaaaa
	ZERO_BITS_POINT     uint32  = 0x7f8f0f0f
	ZERO_BITS_PERCENT   uint32  = 0x7f80f0f0
	BIAS                uint32  = 0x20000000
)

func NewCompactValue() CompactValue {
	return CompactValue{repr: 0x7FC00000}
}

func (cv CompactValue) isUndefined() bool {
	return cv.repr != AUTO_BITS && cv.repr != ZERO_BITS_POINT && cv.repr != ZERO_BITS_PERCENT && IsNaN(math.Float32frombits(cv.repr))
}

func (cv CompactValue) isDefined() bool {
	return !cv.isUndefined()
}

func (cv CompactValue) isAuto() bool {
	return cv.repr == AUTO_BITS
}

func (cv CompactValue) YGValue() YGValue {
	switch cv.repr {
	case AUTO_BITS:
		return YGValueAuto
	case ZERO_BITS_POINT:
		return YGValue{value: 0.0, unit: YGUnitPoint}
	case ZERO_BITS_PERCENT:
		return YGValue{value: 0.0, unit: YGUnitPercent}
	}

	if IsNaN(math.Float32frombits(cv.repr)) {
		return YGValueUndefined
	}

	data := cv.repr
	data &= ^PERCENT_BIT
	data += BIAS

	return YGValue{value: math.Float32frombits(data), unit: If(cv.repr&0x40000000 != 0, YGUnitPercent, YGUnitPoint)}
}

func CompactValueOf(value float32, unit YGUnit) CompactValue {
	if value == 0.0 || (value < LOWER_BOUND && value > -LOWER_BOUND) {
		zero := ZERO_BITS_POINT
		if unit == YGUnitPercent {
			zero = ZERO_BITS_PERCENT
		}
		return CompactValue{repr: zero}
	}

	upperBound := UPPER_BOUND_POINT
	if unit == YGUnitPercent {
		upperBound = UPPER_BOUND_PERCENT
	}
	if value > upperBound || value < -upperBound {
		value = float32(math.Copysign(float64(upperBound), float64(value)))
	}

	unitBit := uint32(0)
	if unit == YGUnitPercent {
		unitBit = PERCENT_BIT
	}
	data := math.Float32bits(value)
	data -= BIAS
	data |= unitBit
	return CompactValue{repr: data}
}

func CompactValueOfZero() CompactValue {
	return CompactValue{repr: ZERO_BITS_POINT}
}

func CompactValueOfUndefined() CompactValue {
	return NewCompactValue()
}

func CompactValueOfAuto() CompactValue {
	return CompactValue{repr: AUTO_BITS}
}

func CompactValueOfMaybe(unit YGUnit, value float32) CompactValue {
	if math.IsNaN(float64(value)) || math.IsInf(float64(value), 0) {
		return CompactValueOfUndefined()
	}
	return CompactValueOf(value, unit)
}
