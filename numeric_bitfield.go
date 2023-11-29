package yoga

import "math/bits"

type EnumBitset struct {
	bits uint
}

func NewEnumBitset() EnumBitset {
	return EnumBitset{bits: 0}
}

func (bitset *EnumBitset) Set(index uint) {
	bitset.bits |= 1 << index
}

func (bitset *EnumBitset) Reset(index uint) {
	bitset.bits &= ^(1 << index)
}

func (bitset *EnumBitset) Test(index uint) bool {
	return (bitset.bits & (1 << index)) != 0
}

func (bitset *EnumBitset) Count() int {
	return bits.OnesCount(bitset.bits)
}

func log2ceilFn(n uint8) uint8 {
	if n < 1 {
		return 0
	}
	return 1 + log2ceilFn(n/2)
}

func mask(bitWidth, index uint8) uint32 {
	return ((1 << bitWidth) - 1) << index
}

func minimumBitCount(enumT any) uint8 {
	count := ordinalCount(enumT)
	if count <= 0 {
		panic("Enums must have at least one entry")
	}
	return log2ceilFn(count - 1)
}

func getEnumData(flags uint32, index uint8, enumT any) uint8 {
	bitWidth := minimumBitCount(enumT)
	return uint8((flags & mask(bitWidth, index)) >> index)
}

func setEnumData(flags *uint32, index uint8, enumT any, newValue uint8) {
	mask := mask(minimumBitCount(enumT), index)
	*flags = (*flags & ^uint32(mask)) | (uint32(newValue) << index & mask)
}

func setBooleanData(flags *uint8, index uint8, value bool) {
	if value {
		*flags |= 1 << index
	} else {
		*flags &= ^(1 << index)
	}
}

func getBooleanData(flags uint8, index uint8) bool {
	return (flags>>index)&1 != 0
}

func ordinalCount(enum any) uint8 {
	switch enum.(type) {
	case Align:
		return AlignCount
	case Dimension:
		return DimensionCount
	case Direction:
		return DirectionCount
	case Display:
		return DisplayCount
	case Edge:
		return EdgeCount
	case ExperimentalFeature:
		return ExperimentalFeatureCount
	case FlexDirection:
		return FlexDirectionCount
	case Justify:
		return JustifyCount
	case LogLevel:
		return LogLevelCount
	case MeasureMode:
		return MeasureModeCount
	case NodeType:
		return NodeTypeCount
	case Overflow:
		return OverflowCount
	case PositionType:
		return PositionTypeCount
	case PrintOptions:
		return PrintOptionsCount
	case Unit:
		return UnitCount
	case Wrap:
		return WrapCount
	case Gutter:
		return GutterCount
	}
	return 0
}
