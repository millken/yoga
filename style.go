package yoga

type YGStyle struct {
	flags uint32

	flex_          FloatOptional
	flexGrow_      FloatOptional
	flexShrink_    FloatOptional
	flexBasis_     CompactValue
	margin_        [EdgeCount]CompactValue
	position_      [EdgeCount]CompactValue
	padding_       [EdgeCount]CompactValue
	border_        [EdgeCount]CompactValue
	gap_           [GutterCount]CompactValue
	dimensions_    [DimensionCount]CompactValue
	minDimensions_ [DimensionCount]CompactValue
	maxDimensions_ [DimensionCount]CompactValue
	// Yoga specific properties, not compatible with flexbox specification
	aspectRatio_ FloatOptional
}

const (
	DefaultFlexGrow      float32 = 0.0
	DefaultFlexShrink    float32 = 0.0
	WebDefaultFlexShrink float32 = 1.0
)

var (
	directionOffset      uint8 = 0
	flexDirectionOffset  uint8 = directionOffset + minimumBitCount(YGDirection(0))
	justifyContentOffset uint8 = flexDirectionOffset + minimumBitCount(YGFlexDirection(0))
	alignContentOffset   uint8 = justifyContentOffset + minimumBitCount(YGJustify(0))
	alignItemsOffset     uint8 = alignContentOffset + minimumBitCount(YGAlign(0))
	alignSelfOffset      uint8 = alignItemsOffset + minimumBitCount(YGAlign(0))
	positionTypeOffset   uint8 = alignSelfOffset + minimumBitCount(YGAlign(0))
	flexWrapOffset       uint8 = positionTypeOffset + minimumBitCount(YGPositionType(0))
	overflowOffset       uint8 = flexWrapOffset + minimumBitCount(YGWrap(0))
	displayOffset        uint8 = overflowOffset + minimumBitCount(YGOverflow(0))

	edgeUndefinedCompactValue      = [EdgeCount]CompactValue{CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined()}
	gutterUndefinedCompactValue    = [GutterCount]CompactValue{CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined()}
	dimensionUndefinedCompactValue = [DimensionCount]CompactValue{CompactValueOfUndefined(), CompactValueOfUndefined()}

	defaultStyle = YGStyle{
		flags: 0,

		flex_:          undefinedFloatOptional,
		flexGrow_:      undefinedFloatOptional,
		flexShrink_:    undefinedFloatOptional,
		margin_:        edgeUndefinedCompactValue,
		position_:      edgeUndefinedCompactValue,
		padding_:       edgeUndefinedCompactValue,
		border_:        edgeUndefinedCompactValue,
		gap_:           gutterUndefinedCompactValue,
		dimensions_:    dimensionUndefinedCompactValue,
		minDimensions_: dimensionUndefinedCompactValue,
		maxDimensions_: dimensionUndefinedCompactValue,
		aspectRatio_:   undefinedFloatOptional,
	}
)

func NewStyle() *YGStyle {
	return &YGStyle{
		flags: 0,

		flex_:          undefinedFloatOptional,
		flexGrow_:      undefinedFloatOptional,
		flexShrink_:    undefinedFloatOptional,
		margin_:        edgeUndefinedCompactValue,
		position_:      edgeUndefinedCompactValue,
		padding_:       edgeUndefinedCompactValue,
		border_:        edgeUndefinedCompactValue,
		gap_:           gutterUndefinedCompactValue,
		dimensions_:    dimensionUndefinedCompactValue,
		minDimensions_: dimensionUndefinedCompactValue,
		maxDimensions_: dimensionUndefinedCompactValue,
		aspectRatio_:   undefinedFloatOptional,
	}
}

func (s *YGStyle) direction() YGDirection {
	return YGDirection(getEnumData(s.flags, directionOffset, YGDirection(0)))
}

func (s *YGStyle) flexDirection() YGFlexDirection {
	return YGFlexDirection(getEnumData(s.flags, flexDirectionOffset, YGFlexDirection(0)))
}

func (s *YGStyle) justifyContent() YGJustify {
	return YGJustify(getEnumData(s.flags, justifyContentOffset, YGJustify(0)))
}

func (s *YGStyle) alignContent() YGAlign {
	return YGAlign(getEnumData(s.flags, alignContentOffset, YGAlign(0)))
}

func (s *YGStyle) alignItems() YGAlign {
	return YGAlign(getEnumData(s.flags, alignItemsOffset, YGAlign(0)))
}

func (s *YGStyle) alignSelf() YGAlign {
	return YGAlign(getEnumData(s.flags, alignSelfOffset, YGAlign(0)))
}

func (s *YGStyle) positionType() YGPositionType {
	return YGPositionType(getEnumData(s.flags, positionTypeOffset, YGPositionType(0)))
}

func (s *YGStyle) flexWrap() YGWrap {
	return YGWrap(getEnumData(s.flags, flexWrapOffset, YGWrap(0)))
}

func (s *YGStyle) overflow() YGOverflow {
	return YGOverflow(getEnumData(s.flags, overflowOffset, YGOverflow(0)))
}

func (s *YGStyle) display() YGDisplay {
	return YGDisplay(getEnumData(s.flags, displayOffset, YGDisplay(0)))
}

func (s *YGStyle) flex() FloatOptional {
	return s.flex_
}

func (s *YGStyle) flexGrow() FloatOptional {
	return s.flexGrow_
}

func (s *YGStyle) flexShrink() FloatOptional {
	return s.flexShrink_
}

func (s *YGStyle) flexBasis() CompactValue {
	return s.flexBasis_
}

func (s *YGStyle) margin() [EdgeCount]CompactValue {
	return s.margin_
}

func (s *YGStyle) position() [EdgeCount]CompactValue {
	return s.position_
}

func (s *YGStyle) padding() [EdgeCount]CompactValue {
	return s.padding_
}

func (s *YGStyle) border() [EdgeCount]CompactValue {
	return s.border_
}

func (s *YGStyle) gap(gutter YGGutter) CompactValue {
	return s.gap_[gutter]
}

// setGap
func (s *YGStyle) setGap(gutter YGGutter, value CompactValue) {
	s.gap_[gutter] = value
}

// dimension
func (s *YGStyle) dimension(dimension YGDimension) CompactValue {
	return s.dimensions_[dimension]
}

// setDimension
func (s *YGStyle) setDimension(dimension YGDimension, value CompactValue) {
	s.dimensions_[dimension] = value
}

// minDimension
func (s *YGStyle) minDimension(dimension YGDimension) CompactValue {
	return s.minDimensions_[dimension]
}

// setMinDimension
func (s *YGStyle) setMinDimension(dimension YGDimension, value CompactValue) {
	s.minDimensions_[dimension] = value
}

// maxDimension
func (s *YGStyle) maxDimension(dimension YGDimension) CompactValue {
	return s.maxDimensions_[dimension]
}

// setMaxDimension
func (s *YGStyle) setMaxDimension(dimension YGDimension, value CompactValue) {
	s.maxDimensions_[dimension] = value
}

// aspectRatio
func (s *YGStyle) aspectRatio() FloatOptional {
	return s.aspectRatio_
}

// resolveColumnGap
func (s *YGStyle) resolveColumnGap() CompactValue {
	if s.gap_[YGGutterColumn].isDefined() {
		return s.gap_[YGGutterColumn]
	} else {
		return s.gap_[YGGutterAll]
	}
}

// resolveRowGap
func (s *YGStyle) resolveRowGap() CompactValue {
	if s.gap_[YGGutterRow].isDefined() {
		return s.gap_[YGGutterRow]
	} else {
		return s.gap_[YGGutterAll]
	}
}

// equal
func (s *YGStyle) equal(other *YGStyle) bool {
	return s.flags == other.flags && inexactEqual(s.flex_.unwrap(), other.flex_.unwrap()) &&
		inexactEquals(s.flexGrow_, other.flexGrow_) &&
		inexactEquals(s.flexShrink_, other.flexShrink_) &&
		inexactEquals(s.flexBasis_, other.flexBasis_) &&
		inexactEquals(s.margin_[:], other.margin_[:]) &&
		inexactEquals(s.position_[:], other.position_[:]) &&
		inexactEquals(s.padding_[:], other.padding_[:]) &&
		inexactEquals(s.border_[:], other.border_[:]) &&
		inexactEquals(s.gap_[:], other.gap_[:]) &&
		inexactEquals(s.dimensions_[:], other.dimensions_[:]) &&
		inexactEquals(s.minDimensions_[:], other.minDimensions_[:]) &&
		inexactEquals(s.maxDimensions_[:], other.maxDimensions_[:]) &&
		inexactEquals(s.aspectRatio_, other.aspectRatio_)
}
