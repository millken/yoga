package yoga

type YGMeasureFunc func(
	node *YGNode,
	width float32,
	widthMode YGMeasureMode,
	height float32,
	heightMode YGMeasureMode,
) YGSize

type YGBaselineFunc func(node *YGNode, width, height float32) float32

type YGPrintFunc func(node *YGNode)

type YGDirtiedFunc func(node *YGNode)

type YGNode struct {
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
	owner_               *YGNode
	children_            []*YGNode
	config_              *YGConfig
	resolvedDimensions_  [2]YGValue
}

func NewNode(config *YGConfig) *YGNode {
	node := &YGNode{}
	return node
}

// getHasNewLayout
func (node *YGNode) getHasNewLayout() bool {
	return node.hasNewLayout_
}

// getNodeType
func (node *YGNode) getNodeType() YGNodeType {
	return node.nodeType_
}

// hasMeasureFunc
func (node *YGNode) hasMeasureFunc() bool {
	return node.measureFunc_ != nil
}

func (node *YGNode) measure(
	width float32,
	widthMode YGMeasureMode,
	height float32,
	heightMode YGMeasureMode,
) YGSize {
	return node.measureFunc_(node, width, widthMode, height, heightMode)
}

// hasBaselineFunc
func (node *YGNode) hasBaselineFunc() bool {
	return node.baselineFunc_ != nil
}

// baseline
func (node *YGNode) baseline(width, height float32) float32 {
	return node.baselineFunc_(node, width, height)
}

// hasErrata
func (node *YGNode) hasErrata(errata YGErrata) bool {
	return node.config_.hasErrata(errata)
}

// getDirtiedFunc
func (node *YGNode) getDirtiedFunc() YGDirtiedFunc {
	return node.dirtiedFunc_
}

// getStyle
func (node *YGNode) getStyle() *YGStyle {
	return &node.style_
}

// getLayout
func (node *YGNode) getLayout() *LayoutResults {
	return &node.layout_
}

// getLineIndex
func (node *YGNode) getLineIndex() uint32 {
	return node.lineIndex_
}

// isReferenceBaseline
func (node *YGNode) isReferenceBaseline() bool {
	return node.isReferenceBaseline_
}

// getOwner
func (node *YGNode) getOwner() *YGNode {
	return node.owner_
}

// getChildren
func (node *YGNode) getChildren() []*YGNode {
	return node.children_
}

// getChild
func (node *YGNode) getChild(index uint32) *YGNode {
	return node.children_[index]
}

// getChildCount
func (node *YGNode) getChildCount() uint32 {
	return uint32(len(node.children_))
}

// getConfig
func (node *YGNode) getConfig() *YGConfig {
	return node.config_
}

// isDirty
func (node *YGNode) isDirty() bool {
	return node.isDirty_
}

// getResolvedDimensions
func (node *YGNode) getResolvedDimensions() [2]YGValue {
	return node.resolvedDimensions_
}

// getResolvedDimension
func (node *YGNode) getResolvedDimension(dimension YGDimension) YGValue {
	return node.resolvedDimensions_[dimension]
}

// computeEdgeValueForColumn
func (node *YGNode) computeEdgeValueForColumn(
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
func (node *YGNode) computeEdgeValueForRow(
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
func (node *YGNode) getInlineStartEdgeUsingErrata(flexDirection YGFlexDirection, direction YGDirection) YGEdge {
	return If(node.hasErrata(YGErrataStartingEndingEdgeFromFlexDirection), flexStartEdge(flexDirection), inlineStartEdge(flexDirection, direction))
}

// getInlineEndEdgeUsingErrata
func (node *YGNode) getInlineEndEdgeUsingErrata(flexDirection YGFlexDirection, direction YGDirection) YGEdge {
	return If(node.hasErrata(YGErrataStartingEndingEdgeFromFlexDirection), flexEndEdge(flexDirection), inlineEndEdge(flexDirection, direction))
}

// isFlexStartPositionDefined
func (node *YGNode) isFlexStartPositionDefined(axis YGFlexDirection) bool {
	startEdge := flexStartEdge(axis)
	leadingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position(), YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().position(), startEdge))
	return leadingPosition.isDefined()
}

// isInlineStartPositionDefined
func (node *YGNode) isInlineStartPositionDefined(axis YGFlexDirection, direction YGDirection) bool {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position(), YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().position(), startEdge))
	return leadingPosition.isDefined()
}

// isFlexEndPositionDefined
func (node *YGNode) isFlexEndPositionDefined(axis YGFlexDirection) bool {
	endEdge := flexEndEdge(axis)
	trailingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position(), YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().position(), endEdge))
	return trailingPosition.isDefined()
}

// isInlineEndPositionDefined
func (node *YGNode) isInlineEndPositionDefined(axis YGFlexDirection, direction YGDirection) bool {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position(), YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().position(), endEdge))
	return trailingPosition.isDefined()
}

// getFlexStartPosition
func (node *YGNode) getFlexStartPosition(axis YGFlexDirection, axisSize float32) float32 {
	startEdge := flexStartEdge(axis)
	leadingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position(), YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().position(), startEdge))
	return resolveCompactValue(leadingPosition, axisSize).unwrapOrDefault(0.0)
}

// getInlineStartPosition
func (node *YGNode) getInlineStartPosition(axis YGFlexDirection, direction YGDirection, axisSize float32) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position(), YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().position(), startEdge))
	return resolveCompactValue(leadingPosition, axisSize).unwrapOrDefault(0.0)
}

// getFlexEndPosition
func (node *YGNode) getFlexEndPosition(axis YGFlexDirection, axisSize float32) float32 {
	endEdge := flexEndEdge(axis)
	trailingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position(), YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().position(), endEdge))
	return resolveCompactValue(trailingPosition, axisSize).unwrapOrDefault(0.0)
}

// getInlineEndPosition
func (node *YGNode) getInlineEndPosition(axis YGFlexDirection, direction YGDirection, axisSize float32) float32 {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingPosition := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().position(), YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().position(), endEdge))
	return resolveCompactValue(trailingPosition, axisSize).unwrapOrDefault(0.0)
}

// getFlexStartMargin
func (node *YGNode) getFlexStartMargin(axis YGFlexDirection, widthSize float32) float32 {
	startEdge := flexStartEdge(axis)
	leadingMargin := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().margin(), YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().margin(), startEdge))
	return resolveCompactValue(leadingMargin, widthSize).unwrapOrDefault(0.0)
}

// getInlineStartMargin
func (node *YGNode) getInlineStartMargin(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingMargin := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().margin(), YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().margin(), startEdge))
	return resolveCompactValue(leadingMargin, widthSize).unwrapOrDefault(0.0)
}

// getFlexEndMargin
func (node *YGNode) getFlexEndMargin(axis YGFlexDirection, widthSize float32) float32 {
	endEdge := flexEndEdge(axis)
	trailingMargin := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().margin(), YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().margin(), endEdge))
	return resolveCompactValue(trailingMargin, widthSize).unwrapOrDefault(0.0)
}

// getInlineEndMargin
func (node *YGNode) getInlineEndMargin(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingMargin := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().margin(), YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().margin(), endEdge))
	return resolveCompactValue(trailingMargin, widthSize).unwrapOrDefault(0.0)
}

// getFlexStartBorder
func (node *YGNode) getFlexStartBorder(axis YGFlexDirection, direction YGDirection) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingBorder := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().border(), YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().border(), startEdge))
	return maxOrDefined(leadingBorder.YGValue().value, 0)
}

// getInlineStartBorder
func (node *YGNode) getInlineStartBorder(axis YGFlexDirection, direction YGDirection) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingBorder := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().border(), YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().border(), startEdge))
	return maxOrDefined(leadingBorder.YGValue().value, 0)
}

// getFlexEndBorder
func (node *YGNode) getFlexEndBorder(axis YGFlexDirection, direction YGDirection) float32 {
	trailRelativeFlexItemEdge := flexEndRelativeEdge(axis, direction)
	trailingBorder := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().border(), trailRelativeFlexItemEdge, flexEndEdge(axis)), node.computeEdgeValueForColumn(node.getStyle().border(), flexEndEdge(axis)))
	return maxOrDefined(trailingBorder.YGValue().value, 0)
}

// getInlineEndBorder
func (node *YGNode) getInlineEndBorder(axis YGFlexDirection, direction YGDirection) float32 {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingBorder := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().border(), YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().border(), endEdge))
	return maxOrDefined(trailingBorder.YGValue().value, 0)
}

// getFlexStartPadding
func (node *YGNode) getFlexStartPadding(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	leadRelativeFlexItemEdge := flexStartRelativeEdge(axis, direction)
	leadingPadding := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().padding(), leadRelativeFlexItemEdge, flexStartEdge(axis)), node.computeEdgeValueForColumn(node.getStyle().padding(), flexStartEdge(axis)))
	return maxOrDefined(resolveCompactValue(leadingPadding, widthSize).unwrap(), 0)
}

// getInlineStartPadding
func (node *YGNode) getInlineStartPadding(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	startEdge := node.getInlineStartEdgeUsingErrata(axis, direction)
	leadingPadding := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().padding(), YGEdgeStart, startEdge), node.computeEdgeValueForColumn(node.getStyle().padding(), startEdge))
	return maxOrDefined(resolveCompactValue(leadingPadding, widthSize).unwrap(), 0)
}

// getFlexEndPadding
func (node *YGNode) getFlexEndPadding(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	trailRelativeFlexItemEdge := flexEndRelativeEdge(axis, direction)
	trailingPadding := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().padding(), trailRelativeFlexItemEdge, flexEndEdge(axis)), node.computeEdgeValueForColumn(node.getStyle().padding(), flexEndEdge(axis)))
	return maxOrDefined(resolveCompactValue(trailingPadding, widthSize).unwrap(), 0)
}

// getInlineEndPadding
func (node *YGNode) getInlineEndPadding(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	endEdge := node.getInlineEndEdgeUsingErrata(axis, direction)
	trailingPadding := If(isRow(axis), node.computeEdgeValueForRow(node.getStyle().padding(), YGEdgeEnd, endEdge), node.computeEdgeValueForColumn(node.getStyle().padding(), endEdge))
	return maxOrDefined(resolveCompactValue(trailingPadding, widthSize).unwrap(), 0)
}

// getFlexStartPaddingAndBorder
func (node *YGNode) getFlexStartPaddingAndBorder(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	return node.getFlexStartPadding(axis, direction, widthSize) + node.getFlexStartBorder(axis, direction)
}

// getInlineStartPaddingAndBorder
func (node *YGNode) getInlineStartPaddingAndBorder(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	return node.getInlineStartPadding(axis, direction, widthSize) + node.getInlineStartBorder(axis, direction)
}

// getFlexEndPaddingAndBorder
func (node *YGNode) getFlexEndPaddingAndBorder(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	return node.getFlexEndPadding(axis, direction, widthSize) + node.getFlexEndBorder(axis, direction)
}

// getInlineEndPaddingAndBorder
func (node *YGNode) getInlineEndPaddingAndBorder(axis YGFlexDirection, direction YGDirection, widthSize float32) float32 {
	return node.getInlineEndPadding(axis, direction, widthSize) + node.getInlineEndBorder(axis, direction)
}

// getMarginForAxis
func (node *YGNode) getMarginForAxis(axis YGFlexDirection, widthSize float32) float32 {
	return node.getInlineStartMargin(axis, YGDirectionLTR, widthSize) + node.getInlineEndMargin(axis, YGDirectionLTR, widthSize)
}

// getGapForAxis
func (node *YGNode) getGapForAxis(axis YGFlexDirection) float32 {
	gap := If(isRow(axis), node.getStyle().resolveColumnGap(), node.getStyle().resolveRowGap())
	return maxOrDefined(resolveCompactValue(gap, 0).unwrap(), 0)
}

// setContext
func (node *YGNode) setContext(context interface{}) {
	node.context_ = context
}

// setPrintFunc
func (node *YGNode) setPrintFunc(printFunc YGPrintFunc) {
	node.printFunc_ = printFunc
}

// setHasNewLayout
func (node *YGNode) setHasNewLayout(hasNewLayout bool) {
	node.hasNewLayout_ = hasNewLayout
}

// setNodeType
func (node *YGNode) setNodeType(nodeType YGNodeType) {
	node.nodeType_ = nodeType
}

// setMeasureFunc
func (node *YGNode) setMeasureFunc(measureFunc YGMeasureFunc) {
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
func (node *YGNode) setBaselineFunc(baselineFunc YGBaselineFunc) {
	node.baselineFunc_ = baselineFunc
}

// setDirtiedFunc
func (node *YGNode) setDirtiedFunc(dirtiedFunc YGDirtiedFunc) {
	node.dirtiedFunc_ = dirtiedFunc
}

// setStyle
func (node *YGNode) setStyle(style YGStyle) {
	node.style_ = style
}

// setLayout
func (node *YGNode) setLayout(layout LayoutResults) {
	node.layout_ = layout
}

// setLineIndex
func (node *YGNode) setLineIndex(lineIndex uint32) {
	node.lineIndex_ = lineIndex
}

// setIsReferenceBaseline
func (node *YGNode) setIsReferenceBaseline(isReferenceBaseline bool) {
	node.isReferenceBaseline_ = isReferenceBaseline
}

// setOwner
func (node *YGNode) setOwner(owner *YGNode) {
	node.owner_ = owner
}

// setChildren
func (node *YGNode) setChildren(children []*YGNode) {
	node.children_ = children
}

// setConfig
func (node *YGNode) setConfig(config *YGConfig) {
	if config == nil {
		panic("config cannot be nil")
	}
	if config.useWebDefaults() != node.config_.useWebDefaults() {
		panic("UseWebDefaults may not be changed after constructing a Node")
	}
	if configUpdateInvalidatesLayout(node.config_, config) {
		node.markDirtyAndPropagate()
	}
	node.config_ = config
}

// setDirty
func (node *YGNode) setDirty(isDirty bool) {
	if isDirty == node.isDirty_ {
		return
	}
	node.isDirty_ = isDirty
	if isDirty && node.dirtiedFunc_ != nil {
		node.dirtiedFunc_(node)
	}
}

// setLayoutLastOwnerDirection
func (node *YGNode) setLayoutLastOwnerDirection(direction YGDirection) {
	node.layout_.lastOwnerDirection = direction
}

// setLayoutComputedFlexBasis
func (node *YGNode) setLayoutComputedFlexBasis(computedFlexBasis FloatOptional) {
	node.layout_.computedFlexBasis = computedFlexBasis
}

// setLayoutComputedFlexBasisGeneration
func (node *YGNode) setLayoutComputedFlexBasisGeneration(computedFlexBasisGeneration uint32) {
	node.layout_.computedFlexBasisGeneration = computedFlexBasisGeneration
}

// setLayoutMeasuredDimension
func (node *YGNode) setLayoutMeasuredDimension(
	measuredDimension float32,
	dimension YGDimension,
) {
	node.layout_.setMeasuredDimension(dimension, measuredDimension)
}

// setLayoutHadOverflow
func (node *YGNode) setLayoutHadOverflow(hadOverflow bool) {
	node.layout_.setHadOverflow(hadOverflow)
}

// setLayoutDimension
func (node *YGNode) setLayoutDimension(dimensionValue float32, dimension YGDimension) {
	node.layout_.setDimension(dimension, dimensionValue)
}

// setLayoutDirection
func (node *YGNode) setLayoutDirection(direction YGDirection) {
	node.layout_.setDirection(direction)
}

// setLayoutMargin
func (node *YGNode) setLayoutMargin(margin float32, edge YGEdge) {
	if int(edge) >= len(node.layout_.margin) {
		panic("Edge must be top/left/bottom/right")
	}
	node.layout_.margin[edge] = margin
}

// setLayoutBorder
func (node *YGNode) setLayoutBorder(border float32, edge YGEdge) {
	if int(edge) >= len(node.layout_.border) {
		panic("Edge must be top/left/bottom/right")
	}
	node.layout_.border[edge] = border
}

// setLayoutPadding
func (node *YGNode) setLayoutPadding(padding float32, edge YGEdge) {
	if int(edge) >= len(node.layout_.padding) {
		panic("Edge must be top/left/bottom/right")
	}
	node.layout_.padding[edge] = padding
}

// setLayoutPosition
func (node *YGNode) setLayoutPosition(position float32, edge YGEdge) {
	if int(edge) >= len(node.layout_.position) {
		panic("Edge must be top/left/bottom/right")
	}
	node.layout_.position[edge] = position
}

// relativePosition
func (node *YGNode) relativePosition(axis YGFlexDirection, direction YGDirection, axisSize float32) float32 {
	if node.isInlineStartPositionDefined(axis, direction) {
		return node.getInlineStartPosition(axis, direction, axisSize)
	}
	return -1 * node.getInlineEndPosition(axis, direction, axisSize)
}

// setPosition
func (node *YGNode) setPosition(
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
func (node *YGNode) getFlexStartMarginValue(axis YGFlexDirection) YGValue {
	if isRow(axis) && node.getStyle().margin()[YGEdgeStart].isDefined() {
		return node.getStyle().margin()[YGEdgeStart].YGValue()
	} else {
		return node.getStyle().margin()[flexStartEdge(axis)].YGValue()
	}
}

// marginTrailingValue
func (node *YGNode) marginTrailingValue(axis YGFlexDirection) YGValue {
	if isRow(axis) && node.getStyle().margin()[YGEdgeEnd].isDefined() {
		return node.getStyle().margin()[YGEdgeEnd].YGValue()
	} else {
		return node.getStyle().margin()[flexEndEdge(axis)].YGValue()
	}
}

// resolveFlexBasisPtr
func (node *YGNode) resolveFlexBasisPtr() YGValue {
	flexBasis := node.getStyle().flexBasis().YGValue()
	if flexBasis.unit != YGUnitAuto && flexBasis.unit != YGUnitUndefined {
		return flexBasis
	}
	if !node.getStyle().flex().isDefined() && node.getStyle().flex().unwrap() > 0 {
		return If(node.getConfig().useWebDefaults(), YGValueAuto, YGValueZero)
	}
	return YGValueAuto
}

// resolveDimension
func (node *YGNode) resolveDimension() {
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
func (node *YGNode) resolveDirection(ownerDirection YGDirection) YGDirection {
	if node.getStyle().direction() == YGDirectionInherit {
		return If(ownerDirection != YGDirectionInherit, ownerDirection, YGDirectionLTR)
	} else {
		return node.getStyle().direction()
	}
}

// clearChildren
func (node *YGNode) clearChildren() {
	node.children_ = make([]*YGNode, 0)
}

// replaceChild
func (node *YGNode) replaceChild(oldChild, newChild *YGNode) {
	for i, child := range node.children_ {
		if child == oldChild {
			node.children_[i] = newChild
			break
		}
	}
}

func (node *YGNode) replaceChildIdx(child *YGNode, index uint32) {
	node.children_[index] = child
}

// insertChild
func (node *YGNode) insertChild(child *YGNode, index uint32) {
	node.children_ = append(node.children_, nil)
	copy(node.children_[index+1:], node.children_[index:])
	node.children_[index] = child
}

func (node *YGNode) removeChild(child *YGNode) {
	for i, c := range node.children_ {
		if c == child {
			node.removeChildIdx(uint32(i))
			break
		}
	}
}

func (node *YGNode) removeChildIdx(index uint32) {
	copy(node.children_[index:], node.children_[index+1:])
	node.children_[len(node.children_)-1] = nil
	node.children_ = node.children_[:len(node.children_)-1]
}

// cloneChildrenIfNeeded
func (node *YGNode) cloneChildrenIfNeeded() {
	for i, child := range node.children_ {
		if child.getOwner() != node {
			child := node.config_.cloneNode(child, node, uint32(i))
			child.setOwner(node)
		}
	}
}

// markDirtyAndPropagate
func (node *YGNode) markDirtyAndPropagate() {
	if !node.isDirty_ {
		node.setDirty(true)
		node.setLayoutComputedFlexBasis(FloatOptional{})
		if node.getOwner() != nil {
			node.getOwner().markDirtyAndPropagate()
		}
	}
}

// resolveFlexGrow
func (node *YGNode) resolveFlexGrow() float32 {
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
func (node *YGNode) resolveFlexShrink() float32 {
	if node.getOwner() == nil {
		return 0
	}
	if node.getStyle().flexShrink().isDefined() {
		return node.getStyle().flexShrink().unwrap()
	}
	if !node.getConfig().useWebDefaults() && node.getStyle().flex().isDefined() && node.getStyle().flex().unwrap() < 0 {
		return -node.getStyle().flex().unwrap()
	}
	return If(node.getConfig().useWebDefaults(), WebDefaultFlexShrink, DefaultFlexShrink)
}

// isNodeFlexible
func (node *YGNode) isNodeFlexible() bool {
	return (node.getStyle().positionType() != YGPositionTypeAbsolute) && (node.resolveFlexGrow() != 0 || node.resolveFlexShrink() != 0)
}

// reset
func (node *YGNode) reset() {
	if node.getChildCount() != 0 {
		panic("Cannot reset a node which still has children attached")
	}
	if node.getOwner() != nil {
		panic("Cannot reset a node still attached to a owner")
	}
	node = NewNode(node.getConfig())
}
