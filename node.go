package yoga

type YGMeasureFunc func(
	node *Node,
	width float32,
	widthMode YGMeasureMode,
	height float32,
	heightMode YGMeasureMode,
) YGSize

type YGBaselineFunc func(node *Node, width, height float32) float32

type YGPrintFunc func(node *Node)

type YGDirtiedFunc func(node *Node)

type Node struct {
	hasNewLayout_        bool
	isReferenceBaseline_ bool
	isDirty_             bool
	nodeType_            YGNodeType
	context_             interface{}
	measureFunc_         YGMeasureFunc
	baselineFunc_        YGBaselineFunc
	printFunc_           YGPrintFunc
	dirtiedFunc_         YGDirtiedFunc
	style_               YGStyle
	layout_              LayoutResults
	lineIndex_           uint32
	owner_               *Node
	children_            []*Node
	config_              *Config
	resolvedDimensions_  [2]YGValue
}

var (
	nodeDefaults = Node{
		hasNewLayout_:        true,
		isReferenceBaseline_: false,
		isDirty_:             false,
		nodeType_:            YGNodeTypeDefault,
		context_:             nil,
		measureFunc_:         nil,
		baselineFunc_:        nil,
		printFunc_:           nil,
		dirtiedFunc_:         nil,
		style_:               defaultStyle,
		layout_:              LayoutResults{},
		lineIndex_:           0,
		owner_:               nil,
		children_:            make([]*Node, 0),
		config_:              nil,
		resolvedDimensions_:  [2]YGValue{},
	}
)

// NewNodeWithConfig
func NewNodeWithConfig(config *Config) *Node {
	node := nodeDefaults
	node.config_ = config
	return &node
}

func NewNode() *Node {

	return NewNodeWithConfig(&defaultConfig)
}

// getHasNewLayout
func (node *Node) getHasNewLayout() bool {
	return node.hasNewLayout_
}

// getNodeType
func (node *Node) getNodeType() YGNodeType {
	return node.nodeType_
}

// hasMeasureFunc
func (node *Node) hasMeasureFunc() bool {
	return node.measureFunc_ != nil
}

func (node *Node) measure(
	width float32,
	widthMode YGMeasureMode,
	height float32,
	heightMode YGMeasureMode,
) YGSize {
	return node.measureFunc_(node, width, widthMode, height, heightMode)
}

// hasBaselineFunc
func (node *Node) hasBaselineFunc() bool {
	return node.baselineFunc_ != nil
}

// baseline
func (node *Node) baseline(width, height float32) float32 {
	return node.baselineFunc_(node, width, height)
}

// hasErrata
func (node *Node) hasErrata(errata YGErrata) bool {
	return node.config_.hasErrata(errata)
}

// getDirtiedFunc
func (node *Node) getDirtiedFunc() YGDirtiedFunc {
	return node.dirtiedFunc_
}

// getStyle
func (node *Node) getStyle() *YGStyle {
	return &node.style_
}

// getLayout
func (node *Node) getLayout() *LayoutResults {
	return &node.layout_
}

// getLineIndex
func (node *Node) getLineIndex() uint32 {
	return node.lineIndex_
}

// isReferenceBaseline
func (node *Node) isReferenceBaseline() bool {
	return node.isReferenceBaseline_
}

// getOwner
func (node *Node) getOwner() *Node {
	return node.owner_
}

// getChildren
func (node *Node) getChildren() []*Node {
	return node.children_
}

// getChild
func (node *Node) getChild(index uint32) *Node {
	return node.children_[index]
}

// getChildCount
func (node *Node) getChildCount() uint32 {
	return uint32(len(node.children_))
}

// getConfig
func (node *Node) getConfig() *Config {
	return node.config_
}

// isDirty
func (node *Node) isDirty() bool {
	return node.isDirty_
}

// getResolvedDimensions
func (node *Node) getResolvedDimensions() [2]YGValue {
	return node.resolvedDimensions_
}

// getResolvedDimension
func (node *Node) getResolvedDimension(dimension YGDimension) YGValue {
	return node.resolvedDimensions_[dimension]
}

// computeEdgeValueForColumn
func (node *Node) computeEdgeValueForColumn(
	edges [EdgeCount]CompactValue,
	edge YGEdge,
) CompactValue {
	if edges[edge].isDefined() {
		return edges[edge]
	} else if edges[YGEdgeVertical].isDefined() {
		return edges[YGEdgeVertical]
	} else {
		return edges[YGEdgeAll]
	}
}

// computeEdgeValueForRow
func (node *Node) computeEdgeValueForRow(
	edges [EdgeCount]CompactValue,
	rowEdge YGEdge,
	edge YGEdge,
) CompactValue {
	if edges[rowEdge].isDefined() {
		return edges[rowEdge]
	} else if edges[edge].isDefined() {
		return edges[edge]
	} else if edges[YGEdgeHorizontal].isDefined() {
		return edges[YGEdgeHorizontal]
	} else {
		return edges[YGEdgeAll]
	}
}

// getInlineStartEdgeUsingErrata
func (node *Node) getInlineStartEdgeUsingErrata(flexDirection YGFlexDirection, direction YGDirection) YGEdge {
	return If(node.hasErrata(YGErrataStartingEndingEdgeFromFlexDirection), flexStartEdge(flexDirection), inlineStartEdge(flexDirection, direction))
}

// getInlineEndEdgeUsingErrata
func (node *Node) getInlineEndEdgeUsingErrata(flexDirection YGFlexDirection, direction YGDirection) YGEdge {
	return If(node.hasErrata(YGErrataStartingEndingEdgeFromFlexDirection), flexEndEdge(flexDirection), inlineEndEdge(flexDirection, direction))
}

// isFlexStartPositionDefined
func (node *Node) isFlexStartPositionDefined(axis YGFlexDirection) bool {
	startEdge := flexStartEdge(axis)
	leadingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position_, YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().position_, startEdge))
	return leadingPosition.isDefined()
}

// isInlineStartPositionDefined
func (node *Node) isInlineStartPositionDefined(axis YGFlexDirection, direction YGDirection) bool {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position_, YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().position_, startEdge))
	return leadingPosition.isDefined()
}

// isFlexEndPositionDefined
func (node *Node) isFlexEndPositionDefined(axis YGFlexDirection) bool {
	endEdge := flexEndEdge(axis)
	trailingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position_, YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().position_, endEdge))
	return trailingPosition.isDefined()
}

// isInlineEndPositionDefined
func (node *Node) isInlineEndPositionDefined(axis YGFlexDirection, direction YGDirection) bool {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position_, YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().position_, endEdge))
	return trailingPosition.isDefined()
}

// getFlexStartPosition
func (node *Node) getFlexStartPosition(axis YGFlexDirection, axisSize float32) float32 {
	startEdge := flexStartEdge(axis)
	leadingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position_, YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().position_, startEdge))
	return resolveCompactValue(leadingPosition, axisSize).unwrapOrDefault(0.0)
}

// getInlineStartPosition
func (node *Node) getInlineStartPosition(axis YGFlexDirection, direction YGDirection, axisSize float32) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position_, YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().position_, startEdge))
	return resolveCompactValue(leadingPosition, axisSize).unwrapOrDefault(0.0)
}

// getFlexEndPosition
func (node *Node) getFlexEndPosition(axis YGFlexDirection, axisSize float32) float32 {
	endEdge := flexEndEdge(axis)
	trailingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position_, YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().position_, endEdge))
	return resolveCompactValue(trailingPosition, axisSize).unwrapOrDefault(0.0)
}

// getInlineEndPosition
func (node *Node) getInlineEndPosition(axis YGFlexDirection, direction YGDirection, axisSize float32) float32 {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position_, YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().position_, endEdge))
	return resolveCompactValue(trailingPosition, axisSize).unwrapOrDefault(0.0)
}

// getFlexStartMargin
func (node *Node) getFlexStartMargin(axis YGFlexDirection, widthSize float32) float32 {
	startEdge := flexStartEdge(axis)
	leadingMargin := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().margin_, YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().margin_, startEdge))
	return resolveCompactValue(leadingMargin, widthSize).unwrapOrDefault(0.0)
}

// getInlineStartMargin
func (node *Node) getInlineStartMargin(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingMargin := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().margin_, YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().margin_, startEdge))
	return resolveCompactValue(leadingMargin, widthSize).unwrapOrDefault(0.0)
}

// getFlexEndMargin
func (node *Node) getFlexEndMargin(axis YGFlexDirection, widthSize float32) float32 {
	endEdge := flexEndEdge(axis)
	trailingMargin := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().margin_, YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().margin_, endEdge))
	return resolveCompactValue(trailingMargin, widthSize).unwrapOrDefault(0.0)
}

// getInlineEndMargin
func (node *Node) getInlineEndMargin(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingMargin := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().margin_, YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().margin_, endEdge))
	return resolveCompactValue(trailingMargin, widthSize).unwrapOrDefault(0.0)
}

// getFlexStartBorder
func (node *Node) getFlexStartBorder(axis YGFlexDirection, direction YGDirection) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingBorder := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().border_, YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().border_, startEdge))
	return maxOrDefined(leadingBorder.YGValue().value, 0)
}

// getInlineStartBorder
func (node *Node) getInlineStartBorder(axis YGFlexDirection, direction YGDirection) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingBorder := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().border_, YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().border_, startEdge))
	return maxOrDefined(leadingBorder.YGValue().value, 0)
}

// getFlexEndBorder
func (node *Node) getFlexEndBorder(axis YGFlexDirection, direction YGDirection) float32 {
	trailRelativeFlexItemEdge := flexEndRelativeEdge(axis, direction)
	trailingBorder := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().border_, trailRelativeFlexItemEdge, flexEndEdge(axis)), node.computeEdgeValueForColumn(node.getStyle().border_, flexEndEdge(axis)))
	return maxOrDefined(trailingBorder.YGValue().value, 0)
}

// getInlineEndBorder
func (node *Node) getInlineEndBorder(axis YGFlexDirection, direction YGDirection) float32 {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingBorder := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().border_, YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().border_, endEdge))
	return maxOrDefined(trailingBorder.YGValue().value, 0)
}

// getFlexStartPadding
func (node *Node) getFlexStartPadding(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	leadRelativeFlexItemEdge := flexStartRelativeEdge(axis, direction)
	leadingPadding := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().padding_, leadRelativeFlexItemEdge, flexStartEdge(axis)), node.computeEdgeValueForColumn(node.getStyle().padding_, flexStartEdge(axis)))
	return maxOrDefined(resolveCompactValue(leadingPadding, widthSize).unwrap(), 0)
}

// getInlineStartPadding
func (node *Node) getInlineStartPadding(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingPadding := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().padding_, YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().padding_, startEdge))
	return maxOrDefined(resolveCompactValue(leadingPadding, widthSize).unwrap(), 0)
}

// getFlexEndPadding
func (node *Node) getFlexEndPadding(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	trailRelativeFlexItemEdge := flexEndRelativeEdge(axis, direction)
	trailingPadding := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().padding_, trailRelativeFlexItemEdge, flexEndEdge(axis)), node.computeEdgeValueForColumn(node.getStyle().padding_, flexEndEdge(axis)))
	return maxOrDefined(resolveCompactValue(trailingPadding, widthSize).unwrap(), 0)
}

// getInlineEndPadding
func (node *Node) getInlineEndPadding(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingPadding := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().padding_, YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().padding_, endEdge))
	return maxOrDefined(resolveCompactValue(trailingPadding, widthSize).unwrap(), 0)
}

// getFlexStartPaddingAndBorder
func (node *Node) getFlexStartPaddingAndBorder(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	return node.getFlexStartPadding(axis, direction, widthSize) + node.getFlexStartBorder(axis, direction)
}

// getInlineStartPaddingAndBorder
func (node *Node) getInlineStartPaddingAndBorder(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	return node.getInlineStartPadding(axis, direction, widthSize) + node.getInlineStartBorder(axis, direction)
}

// getFlexEndPaddingAndBorder
func (node *Node) getFlexEndPaddingAndBorder(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	return node.getFlexEndPadding(axis, direction, widthSize) + node.getFlexEndBorder(axis, direction)
}

// getInlineEndPaddingAndBorder
func (node *Node) getInlineEndPaddingAndBorder(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	return node.getInlineEndPadding(axis, direction, widthSize) + node.getInlineEndBorder(axis, direction)
}

// getMarginForAxis
func (node *Node) getMarginForAxis(axis YGFlexDirection, widthSize float32) float32 {
	return node.getInlineStartMargin(axis, YGDirectionLTR, widthSize) + node.getInlineEndMargin(axis, YGDirectionLTR, widthSize)
}

// getGapForAxis
func (node *Node) getGapForAxis(axis YGFlexDirection) float32 {
	gap := If(isRow(axis), node.getStyle().resolveColumnGap(), node.getStyle().resolveRowGap())
	return maxOrDefined(resolveCompactValue(gap, 0).unwrap(), 0)
}

// setContext
func (node *Node) setContext(context interface{}) {
	node.context_ = context
}

// setPrintFunc
func (node *Node) setPrintFunc(printFunc YGPrintFunc) {
	node.printFunc_ = printFunc
}

// setHasNewLayout
func (node *Node) setHasNewLayout(hasNewLayout bool) {
	node.hasNewLayout_ = hasNewLayout
}

// setNodeType
func (node *Node) setNodeType(nodeType YGNodeType) {
	node.nodeType_ = nodeType
}

// setMeasureFunc
func (node *Node) setMeasureFunc(measureFunc YGMeasureFunc) {
	if measureFunc == nil {
		node.setNodeType(YGNodeTypeDefault)
	} else {
		if node.getChildCount() == 0 {
			panic("Cannot set measure function: Nodes with measure functions cannot have children.")
		}
		node.setNodeType(YGNodeTypeText)
	}
	node.measureFunc_ = measureFunc
}

// setBaselineFunc
func (node *Node) setBaselineFunc(baselineFunc YGBaselineFunc) {
	node.baselineFunc_ = baselineFunc
}

// setDirtiedFunc
func (node *Node) setDirtiedFunc(dirtiedFunc YGDirtiedFunc) {
	node.dirtiedFunc_ = dirtiedFunc
}

// setStyle
func (node *Node) setStyle(style YGStyle) {
	node.style_ = style
}

// setLayout
func (node *Node) setLayout(layout LayoutResults) {
	node.layout_ = layout
}

// setLineIndex
func (node *Node) setLineIndex(lineIndex uint32) {
	node.lineIndex_ = lineIndex
}

// setIsReferenceBaseline
func (node *Node) setIsReferenceBaseline(isReferenceBaseline bool) {
	node.isReferenceBaseline_ = isReferenceBaseline
}

// setOwner
func (node *Node) setOwner(owner *Node) {
	node.owner_ = owner
}

// setChildren
func (node *Node) setChildren(children []*Node) {
	node.children_ = children
}

// setConfig
func (node *Node) setConfig(config *Config) {
	if config == nil {
		panic("config cannot be nil")
	}
	if config.UseWebDefaults() != node.config_.UseWebDefaults() {
		panic("UseWebDefaults may not be changed after constructing a Node")
	}
	if configUpdateInvalidatesLayout(node.config_, config) {
		node.markDirtyAndPropagate()
	}
	node.config_ = config
}

// setDirty
func (node *Node) setDirty(isDirty bool) {
	if isDirty == node.isDirty_ {
		return
	}
	node.isDirty_ = isDirty
	if isDirty && node.dirtiedFunc_ != nil {
		node.dirtiedFunc_(node)
	}
}

// setLayoutLastOwnerDirection
func (node *Node) setLayoutLastOwnerDirection(direction YGDirection) {
	node.getLayout().lastOwnerDirection = direction
}

// setLayoutComputedFlexBasis
func (node *Node) setLayoutComputedFlexBasis(computedFlexBasis FloatOptional) {
	node.getLayout().computedFlexBasis = computedFlexBasis
}

// setLayoutComputedFlexBasisGeneration
func (node *Node) setLayoutComputedFlexBasisGeneration(computedFlexBasisGeneration uint32) {
	node.getLayout().computedFlexBasisGeneration = computedFlexBasisGeneration
}

// setLayoutMeasuredDimension
func (node *Node) setLayoutMeasuredDimension(
	measuredDimension float32,
	dimension YGDimension,
) {
	node.getLayout().setMeasuredDimension(dimension, measuredDimension)
}

// setLayoutHadOverflow
func (node *Node) setLayoutHadOverflow(hadOverflow bool) {
	node.getLayout().setHadOverflow(hadOverflow)
}

// setLayoutDimension
func (node *Node) setLayoutDimension(dimensionValue float32, dimension YGDimension) {
	node.getLayout().setDimension(dimension, dimensionValue)
}

// setLayoutDirection
func (node *Node) setLayoutDirection(direction YGDirection) {
	node.getLayout().setDirection(direction)
}

// setLayoutMargin
func (node *Node) setLayoutMargin(margin float32, edge YGEdge) {
	if int(edge) >= len(node.getLayout().margin) {
		panic("Edge must be top/left/bottom/right")
	}
	node.getLayout().margin[edge] = margin
}

// setLayoutBorder
func (node *Node) setLayoutBorder(border float32, edge YGEdge) {
	if int(edge) >= len(node.getLayout().border) {
		panic("Edge must be top/left/bottom/right")
	}
	node.getLayout().border[edge] = border
}

// setLayoutPadding
func (node *Node) setLayoutPadding(padding float32, edge YGEdge) {
	if int(edge) >= len(node.getLayout().padding) {
		panic("Edge must be top/left/bottom/right")
	}
	node.getLayout().padding[edge] = padding
}

// setLayoutPosition
func (node *Node) setLayoutPosition(position float32, edge YGEdge) {
	if int(edge) >= len(node.getLayout().position) {
		panic("Edge must be top/left/bottom/right")
	}
	node.getLayout().position[edge] = position
}

// relativePosition
func (node *Node) relativePosition(axis YGFlexDirection, direction YGDirection, axisSize float32) float32 {
	if node.isInlineStartPositionDefined(axis, direction) {
		return node.getInlineStartPosition(axis, direction, axisSize)
	}
	return -1 * node.getInlineEndPosition(axis, direction, axisSize)
}

// setPosition
func (node *Node) setPosition(
	direction YGDirection,
	mainSize float32,
	crossSize float32,
	ownerWidth float32) {
	/* Root nodes should be always layouted as LTR, so we don't return negative
	 * values. */
	directionRespectingRoot := If(node.owner_ != nil, direction, YGDirectionLTR)
	mainAxis := resolveDirection(node.getStyle().flexDirection(), directionRespectingRoot)
	crossAxis := resolveCrossDirection(mainAxis, directionRespectingRoot)
	// Here we should check for `PositionType::Static` and in this case zero inset
	// properties (left, right, top, bottom, begin, end).
	// https://www.w3.org/TR/css-position-3/#valdef-position-static
	relativePositionMain := node.relativePosition(mainAxis, directionRespectingRoot, mainSize)
	relativePositionCross := node.relativePosition(crossAxis, directionRespectingRoot, crossSize)

	mainAxisLeadingEdge := node.getInlineStartEdgeUsingErrata(mainAxis, direction)
	mainAxisTrailingEdge := node.getInlineEndEdgeUsingErrata(mainAxis, direction)
	crossAxisLeadingEdge := node.getInlineStartEdgeUsingErrata(crossAxis, direction)
	crossAxisTrailingEdge := node.getInlineEndEdgeUsingErrata(crossAxis, direction)

	node.setLayoutPosition((node.getInlineStartMargin(mainAxis, direction, ownerWidth) + relativePositionMain), mainAxisLeadingEdge)
	node.setLayoutPosition((node.getInlineEndMargin(mainAxis, direction, ownerWidth) + relativePositionMain), mainAxisTrailingEdge)
	node.setLayoutPosition((node.getInlineStartMargin(crossAxis, direction, ownerWidth) + relativePositionCross), crossAxisLeadingEdge)
	node.setLayoutPosition((node.getInlineEndMargin(crossAxis, direction, ownerWidth) + relativePositionCross), crossAxisTrailingEdge)
}

// getFlexStartMarginValue
func (node *Node) getFlexStartMarginValue(axis YGFlexDirection) YGValue {
	if isRow(axis) && node.getStyle().margin_[YGEdgeStart].isDefined() {
		return node.getStyle().margin_[YGEdgeStart].YGValue()
	} else {
		return node.getStyle().margin_[flexStartEdge(axis)].YGValue()
	}
}

// marginTrailingValue
func (node *Node) marginTrailingValue(axis YGFlexDirection) YGValue {
	if isRow(axis) && node.getStyle().margin_[YGEdgeEnd].isDefined() {
		return node.getStyle().margin_[YGEdgeEnd].YGValue()
	} else {
		return node.getStyle().margin_[flexEndEdge(axis)].YGValue()
	}
}

// resolveFlexBasisPtr
func (node *Node) resolveFlexBasisPtr() YGValue {
	flexBasis := node.getStyle().flexBasis().YGValue()
	if flexBasis.unit != YGUnitAuto && flexBasis.unit != YGUnitUndefined {
		return flexBasis
	}
	if !node.getStyle().flex().isDefined() && node.getStyle().flex().unwrap() > 0 {
		return If(node.getConfig().UseWebDefaults(), YGValueAuto, YGValueZero)
	}
	return YGValueAuto
}

// resolveDimension
func (node *Node) resolveDimension() {
	style := node.getStyle()
	for dim := YGDimensionWidth; dim < DimensionCount; dim++ {
		if style.maxDimension(dim).isDefined() &&
			inexactEquals(style.maxDimension(dim), style.minDimension(dim)) {
			node.resolvedDimensions_[dim] = style.maxDimension(dim).YGValue()
		} else {
			node.resolvedDimensions_[dim] = style.dimension(dim).YGValue()
		}
	}
}

// resolveDirection
func (node *Node) resolveDirection(ownerDirection YGDirection) YGDirection {
	if node.getStyle().direction() == YGDirectionInherit {
		return If(ownerDirection != YGDirectionInherit, ownerDirection, YGDirectionLTR)
	} else {
		return node.getStyle().direction()
	}
}

// clearChildren
func (node *Node) clearChildren() {
	node.children_ = make([]*Node, 0)
}

// replaceChild
func (node *Node) replaceChild(oldChild, newChild *Node) {
	for i, child := range node.children_ {
		if child == oldChild {
			node.children_[i] = newChild
			break
		}
	}
}

func (node *Node) replaceChildIdx(child *Node, index uint32) {
	node.children_[index] = child
}

// insertChild
func (node *Node) InsertChild(child *Node, index uint32) {
	node.children_ = append(node.children_, nil)
	copy(node.children_[index+1:], node.children_[index:])
	node.children_[index] = child
}

func (node *Node) removeChild(child *Node) {
	for i, c := range node.children_ {
		if c == child {
			node.removeChildIdx(uint32(i))
			break
		}
	}
}

func (node *Node) removeChildIdx(index uint32) {
	copy(node.children_[index:], node.children_[index+1:])
	node.children_[len(node.children_)-1] = nil
	node.children_ = node.children_[:len(node.children_)-1]
}

// cloneChildrenIfNeeded
func (node *Node) cloneChildrenIfNeeded() {
	for i, child := range node.children_ {
		if child.getOwner() != node {
			child := node.config_.cloneNode(child, node, uint32(i))
			child.setOwner(node)
		}
	}
}

// markDirtyAndPropagate
func (node *Node) markDirtyAndPropagate() {
	if !node.isDirty_ {
		node.setDirty(true)
		node.setLayoutComputedFlexBasis(FloatOptional{})
		if node.getOwner() != nil {
			node.getOwner().markDirtyAndPropagate()
		}
	}
}

// resolveFlexGrow
func (node *Node) resolveFlexGrow() float32 {
	if node.getOwner() == nil {
		return 0
	}
	if node.getStyle().flexGrow().isDefined() {
		return node.getStyle().flexGrow().unwrap()
	}
	if node.getStyle().flex().isDefined() && node.getStyle().flex().unwrap() > 0 {
		return node.getStyle().flex().unwrap()
	}
	return DefaultFlexGrow
}

// resolveFlexShrink
func (node *Node) resolveFlexShrink() float32 {
	if node.getOwner() == nil {
		return 0
	}
	if node.getStyle().flexShrink().isDefined() {
		return node.getStyle().flexShrink().unwrap()
	}
	if !node.getConfig().UseWebDefaults() && node.getStyle().flex().isDefined() && node.getStyle().flex().unwrap() < 0 {
		return -node.getStyle().flex().unwrap()
	}
	return If(node.getConfig().UseWebDefaults(), WebDefaultFlexShrink, DefaultFlexShrink)
}

// isNodeFlexible
func (node *Node) isNodeFlexible() bool {
	return (node.getStyle().positionType() != YGPositionTypeAbsolute) && (node.resolveFlexGrow() != 0 || node.resolveFlexShrink() != 0)
}

// reset
func (node *Node) reset() {
	if node.getChildCount() != 0 {
		panic("Cannot reset a node which still has children attached")
	}
	if node.getOwner() != nil {
		panic("Cannot reset a node still attached to a owner")
	}
	node = NewNodeWithConfig(node.getConfig())
}
