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
	flexDirectionOffset  uint8 = directionOffset + minimumBitCount(Direction(0))
	justifyContentOffset uint8 = flexDirectionOffset + minimumBitCount(FlexDirection(0))
	alignContentOffset   uint8 = justifyContentOffset + minimumBitCount(Justify(0))
	alignItemsOffset     uint8 = alignContentOffset + minimumBitCount(Align(0))
	alignSelfOffset      uint8 = alignItemsOffset + minimumBitCount(Align(0))
	positionTypeOffset   uint8 = alignSelfOffset + minimumBitCount(Align(0))
	flexWrapOffset       uint8 = positionTypeOffset + minimumBitCount(PositionType(0))
	overflowOffset       uint8 = flexWrapOffset + minimumBitCount(Wrap(0))
	displayOffset        uint8 = overflowOffset + minimumBitCount(Overflow(0))

	edgeUndefinedCompactValue      = [EdgeCount]CompactValue{CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined()}
	gutterUndefinedCompactValue    = [GutterCount]CompactValue{CompactValueOfUndefined(), CompactValueOfUndefined(), CompactValueOfUndefined()}
	dimensionUndefinedCompactValue = [DimensionCount]CompactValue{CompactValueOfUndefined(), CompactValueOfUndefined()}

	defaultStyle = YGStyle{
		flags:          0,
		flex_:          undefinedFloatOptional,
		flexGrow_:      undefinedFloatOptional,
		flexShrink_:    undefinedFloatOptional,
		flexBasis_:     CompactValueOfUndefined(),
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

func (s *YGStyle) direction() Direction {
	return Direction(getEnumData(s.flags, directionOffset, Direction(0)))
}

func (s *YGStyle) flexDirection() FlexDirection {
	return FlexDirection(getEnumData(s.flags, flexDirectionOffset, FlexDirection(0)))
}

func (s *YGStyle) justifyContent() Justify {
	return Justify(getEnumData(s.flags, justifyContentOffset, Justify(0)))
}

func (s *YGStyle) alignContent() Align {
	return Align(getEnumData(s.flags, alignContentOffset, Align(0)))
}

func (s *YGStyle) alignItems() Align {
	return Align(getEnumData(s.flags, alignItemsOffset, Align(0)))
}

func (s *YGStyle) alignSelf() Align {
	return Align(getEnumData(s.flags, alignSelfOffset, Align(0)))
}

func (s *YGStyle) positionType() PositionType {
	return PositionType(getEnumData(s.flags, positionTypeOffset, PositionType(0)))
}

func (s *YGStyle) flexWrap() Wrap {
	return Wrap(getEnumData(s.flags, flexWrapOffset, Wrap(0)))
}

func (s *YGStyle) overflow() Overflow {
	return Overflow(getEnumData(s.flags, overflowOffset, Overflow(0)))
}

func (s *YGStyle) display() Display {
	return Display(getEnumData(s.flags, displayOffset, Display(0)))
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

func (s *YGStyle) margin(edge Edge) CompactValue {
	return s.margin_[edge]
}

// setMargin
func (s *YGStyle) setMargin(edge Edge, value CompactValue) {
	s.margin_[edge] = value
}

func (s *YGStyle) position(edge Edge) CompactValue {
	return s.position_[edge]
}

// setPosition
func (s *YGStyle) setPosition(edge Edge, value CompactValue) {
	s.position_[edge] = value
}

func (s *YGStyle) padding(edge Edge) CompactValue {
	return s.padding_[edge]
}

// setPadding
func (s *YGStyle) setPadding(edge Edge, value CompactValue) {
	s.padding_[edge] = value
}

func (s *YGStyle) border(edge Edge) CompactValue {
	return s.border_[edge]
}

// setBorder
func (s *YGStyle) setBorder(edge Edge, value CompactValue) {
	s.border_[edge] = value
}

func (s *YGStyle) gap(gutter Gutter) CompactValue {
	return s.gap_[gutter]
}

// setGap
func (s *YGStyle) setGap(gutter Gutter, value CompactValue) {
	s.gap_[gutter] = value
}

// dimension
func (s *YGStyle) dimension(dimension Dimension) CompactValue {
	return s.dimensions_[dimension]
}

// setDimension
func (s *YGStyle) setDimension(dimension Dimension, value CompactValue) {
	s.dimensions_[dimension] = value
}

// minDimension
func (s *YGStyle) minDimension(dimension Dimension) CompactValue {
	return s.minDimensions_[dimension]
}

// setMinDimension
func (s *YGStyle) setMinDimension(dimension Dimension, value CompactValue) {
	s.minDimensions_[dimension] = value
}

// maxDimension
func (s *YGStyle) maxDimension(dimension Dimension) CompactValue {
	return s.maxDimensions_[dimension]
}

// setMaxDimension
func (s *YGStyle) setMaxDimension(dimension Dimension, value CompactValue) {
	s.maxDimensions_[dimension] = value
}

// aspectRatio
func (s *YGStyle) aspectRatio() FloatOptional {
	return s.aspectRatio_
}

// resolveColumnGap
func (s *YGStyle) resolveColumnGap() CompactValue {
	if s.gap_[GutterColumn].isDefined() {
		return s.gap_[GutterColumn]
	} else {
		return s.gap_[GutterAll]
	}
}

// resolveRowGap
func (s *YGStyle) resolveRowGap() CompactValue {
	if s.gap_[GutterRow].isDefined() {
		return s.gap_[GutterRow]
	} else {
		return s.gap_[GutterAll]
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
