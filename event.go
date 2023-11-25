package yoga

type LayoutType int

const (
	LayoutTypeLayout LayoutType = iota
	LayoutTypeMeasure
	LayoutTypeCachedLayout
	LayoutTypeCachedMeasure
)

type LayoutPassReason uint8

const (
	LayoutPassReasonInitial LayoutPassReason = iota
	LayoutPassReasonAbsLayout
	LayoutPassReasonStretch
	LayoutPassReasonMultilineStretch
	LayoutPassReasonFlexLayout
	LayoutPassReasonMeasureChild
	LayoutPassReasonAbsMeasureChild
	LayoutPassReasonFlexMeasure
	LayoutPassReasonCount
)

func (r LayoutPassReason) String() string {
	switch r {
	case LayoutPassReasonInitial:
		return "Initial"
	case LayoutPassReasonAbsLayout:
		return "AbsLayout"
	case LayoutPassReasonStretch:
		return "Stretch"
	case LayoutPassReasonMultilineStretch:
		return "MultilineStretch"
	case LayoutPassReasonFlexLayout:
		return "FlexLayout"
	case LayoutPassReasonMeasureChild:
		return "MeasureChild"
	case LayoutPassReasonAbsMeasureChild:
		return "AbsMeasureChild"
	case LayoutPassReasonFlexMeasure:
		return "FlexMeasure"
	default:
		return "Unknown"
	}
}

type LayoutData struct {
	layouts                     int32
	measures                    int32
	maxMeasureCache             uint32
	cachedLayouts               int32
	cachedMeasures              int32
	measureCallbacks            int32
	measureCallbackReasonsCount [LayoutPassReasonCount]uint8
}
