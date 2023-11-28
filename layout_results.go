package yoga

const MaxCachedMeasurements int32 = 8

type LayoutResults struct {
	position            [4]float32
	margin              [4]float32
	border              [4]float32
	padding             [4]float32
	direction_          YGDirection
	hadOverflow_        bool
	dimension_          [2]float32
	measuredDimensions_ [2]float32

	computedFlexBasisGeneration uint32
	computedFlexBasis           FloatOptional
	generationCount             uint32
	lastOwnerDirection          YGDirection
	nextCachedMeasurementsIndex uint32
	cachedMeasurements          [MaxCachedMeasurements]CachedMeasurement
	cachedLayout                CachedMeasurement
}

func NewLayoutResults() *LayoutResults {
	return &LayoutResults{
		position:                    [4]float32{YGUndefined, YGUndefined, YGUndefined, YGUndefined},
		margin:                      [4]float32{YGUndefined, YGUndefined, YGUndefined, YGUndefined},
		border:                      [4]float32{YGUndefined, YGUndefined, YGUndefined, YGUndefined},
		padding:                     [4]float32{YGUndefined, YGUndefined, YGUndefined, YGUndefined},
		direction_:                  YGDirectionInherit,
		hadOverflow_:                false,
		dimension_:                  [2]float32{YGUndefined, YGUndefined},
		measuredDimensions_:         [2]float32{YGUndefined, YGUndefined},
		computedFlexBasisGeneration: 0,
		computedFlexBasis:           undefinedFloatOptional,
		lastOwnerDirection:          YGDirectionInherit,
		cachedMeasurements:          [MaxCachedMeasurements]CachedMeasurement{},
		cachedLayout:                CachedMeasurement{},
	}
}

func (l *LayoutResults) direction() YGDirection {
	return l.direction_
}

func (l *LayoutResults) setDirection(direction YGDirection) {
	l.direction_ = direction
}

func (l *LayoutResults) hadOverflow() bool {
	return l.hadOverflow_
}

func (l *LayoutResults) setHadOverflow(hadOverflow bool) {
	l.hadOverflow_ = hadOverflow
}

func (l *LayoutResults) dimension(axis YGDimension) float32 {
	return l.dimension_[axis]
}

func (l *LayoutResults) setDimension(axis YGDimension, value float32) {
	l.dimension_[axis] = value
}

func (l *LayoutResults) measuredDimension(axis YGDimension) float32 {
	return l.measuredDimensions_[axis]
}

func (l *LayoutResults) setMeasuredDimension(axis YGDimension, value float32) {
	l.measuredDimensions_[axis] = value
}

func (l *LayoutResults) isEqual(layout *LayoutResults) bool {
	isEqual := inexactEquals(l.position[:], layout.position[:]) &&
		inexactEquals(l.margin[:], layout.margin[:]) &&
		inexactEquals(l.border[:], layout.border[:]) &&
		inexactEquals(l.padding[:], layout.padding[:]) &&
		l.direction() == layout.direction() &&
		l.hadOverflow() == layout.hadOverflow() &&
		l.lastOwnerDirection == layout.lastOwnerDirection &&
		l.nextCachedMeasurementsIndex == layout.nextCachedMeasurementsIndex &&
		l.cachedLayout == layout.cachedLayout &&
		l.computedFlexBasis == layout.computedFlexBasis
	for i := uint32(0); i < uint32(MaxCachedMeasurements) && isEqual; i++ {
		isEqual = isEqual && l.cachedMeasurements[i] == layout.cachedMeasurements[i]
	}
	if !isUndefined(l.measuredDimensions_[0]) || !isUndefined(layout.measuredDimensions_[0]) {
		isEqual = isEqual && (l.measuredDimensions_[0] == layout.measuredDimensions_[0])
	}
	if !isUndefined(l.measuredDimensions_[1]) || !isUndefined(layout.measuredDimensions_[1]) {
		isEqual = isEqual && (l.measuredDimensions_[1] == layout.measuredDimensions_[1])
	}
	return isEqual
}
