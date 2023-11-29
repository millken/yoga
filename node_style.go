package yoga

/*
YG_EXPORT void YGNodeCopyStyle(YGNodeRef dstNode, YGNodeConstRef srcNode);

YG_EXPORT void YGNodeStyleSetDirection(YGNodeRef node, Direction direction);
YG_EXPORT Direction YGNodeStyleGetDirection(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetFlexDirection(
    YGNodeRef node,
    FlexDirection flexDirection);
YG_EXPORT FlexDirection YGNodeStyleGetFlexDirection(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetJustifyContent(
    YGNodeRef node,
    Justify justifyContent);
YG_EXPORT Justify YGNodeStyleGetJustifyContent(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetAlignContent(YGNodeRef node, Align alignContent);
YG_EXPORT Align YGNodeStyleGetAlignContent(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetAlignItems(YGNodeRef node, Align alignItems);
YG_EXPORT Align YGNodeStyleGetAlignItems(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetAlignSelf(YGNodeRef node, Align alignSelf);
YG_EXPORT Align YGNodeStyleGetAlignSelf(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetPositionType(
    YGNodeRef node,
    PositionType positionType);
YG_EXPORT PositionType YGNodeStyleGetPositionType(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetFlexWrap(YGNodeRef node, Wrap flexWrap);
YG_EXPORT Wrap YGNodeStyleGetFlexWrap(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetOverflow(YGNodeRef node, Overflow overflow);
YG_EXPORT Overflow YGNodeStyleGetOverflow(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetDisplay(YGNodeRef node, Display display);
YG_EXPORT Display YGNodeStyleGetDisplay(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetFlex(YGNodeRef node, float flex);
YG_EXPORT float YGNodeStyleGetFlex(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetFlexGrow(YGNodeRef node, float flexGrow);
YG_EXPORT float YGNodeStyleGetFlexGrow(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetFlexShrink(YGNodeRef node, float flexShrink);
YG_EXPORT float YGNodeStyleGetFlexShrink(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetFlexBasis(YGNodeRef node, float flexBasis);
YG_EXPORT void YGNodeStyleSetFlexBasisPercent(YGNodeRef node, float flexBasis);
YG_EXPORT void YGNodeStyleSetFlexBasisAuto(YGNodeRef node);
YG_EXPORT YGValue YGNodeStyleGetFlexBasis(YGNodeConstRef node);

YG_EXPORT void
YGNodeStyleSetPosition(YGNodeRef node, Edge edge, float position);
YG_EXPORT void
YGNodeStyleSetPositionPercent(YGNodeRef node, Edge edge, float position);
YG_EXPORT YGValue YGNodeStyleGetPosition(YGNodeConstRef node, Edge edge);

YG_EXPORT void YGNodeStyleSetMargin(YGNodeRef node, Edge edge, float margin);
YG_EXPORT void
YGNodeStyleSetMarginPercent(YGNodeRef node, Edge edge, float margin);
YG_EXPORT void YGNodeStyleSetMarginAuto(YGNodeRef node, Edge edge);
YG_EXPORT YGValue YGNodeStyleGetMargin(YGNodeConstRef node, Edge edge);

YG_EXPORT void
YGNodeStyleSetPadding(YGNodeRef node, Edge edge, float padding);
YG_EXPORT void
YGNodeStyleSetPaddingPercent(YGNodeRef node, Edge edge, float padding);
YG_EXPORT YGValue YGNodeStyleGetPadding(YGNodeConstRef node, Edge edge);

YG_EXPORT void YGNodeStyleSetBorder(YGNodeRef node, Edge edge, float border);
YG_EXPORT float YGNodeStyleGetBorder(YGNodeConstRef node, Edge edge);

YG_EXPORT void
YGNodeStyleSetGap(YGNodeRef node, Gutter gutter, float gapLength);
YG_EXPORT float YGNodeStyleGetGap(YGNodeConstRef node, Gutter gutter);

YG_EXPORT void YGNodeStyleSetWidth(YGNodeRef node, float width);
YG_EXPORT void YGNodeStyleSetWidthPercent(YGNodeRef node, float width);
YG_EXPORT void YGNodeStyleSetWidthAuto(YGNodeRef node);
YG_EXPORT YGValue YGNodeStyleGetWidth(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetHeight(YGNodeRef node, float height);
YG_EXPORT void YGNodeStyleSetHeightPercent(YGNodeRef node, float height);
YG_EXPORT void YGNodeStyleSetHeightAuto(YGNodeRef node);
YG_EXPORT YGValue YGNodeStyleGetHeight(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetMinWidth(YGNodeRef node, float minWidth);
YG_EXPORT void YGNodeStyleSetMinWidthPercent(YGNodeRef node, float minWidth);
YG_EXPORT YGValue YGNodeStyleGetMinWidth(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetMinHeight(YGNodeRef node, float minHeight);
YG_EXPORT void YGNodeStyleSetMinHeightPercent(YGNodeRef node, float minHeight);
YG_EXPORT YGValue YGNodeStyleGetMinHeight(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetMaxWidth(YGNodeRef node, float maxWidth);
YG_EXPORT void YGNodeStyleSetMaxWidthPercent(YGNodeRef node, float maxWidth);
YG_EXPORT YGValue YGNodeStyleGetMaxWidth(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetMaxHeight(YGNodeRef node, float maxHeight);
YG_EXPORT void YGNodeStyleSetMaxHeightPercent(YGNodeRef node, float maxHeight);
YG_EXPORT YGValue YGNodeStyleGetMaxHeight(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetAspectRatio(YGNodeRef node, float aspectRatio);
YG_EXPORT float YGNodeStyleGetAspectRatio(YGNodeConstRef node);
*/

func updateStyle[T float32 | Direction](node *Node, prop string, value T) {

}

// StyleSetDirection
func (n *Node) StyleSetDirection(direction Direction) {
	if n.getStyle().direction() != direction {
		n.getStyle().setDirection(direction)
		n.markDirtyAndPropagate()
	}
}

// StyleGetDirection
func (n *Node) StyleGetDirection() Direction {
	return n.getStyle().direction()
}

// StyleSetFlexDirection
func (n *Node) StyleSetFlexDirection(flexDirection FlexDirection) {
	if n.getStyle().flexDirection() != flexDirection {
		n.getStyle().setFlexDirection(flexDirection)
		n.markDirtyAndPropagate()
	}
}

// StyleGetFlexDirection
func (n *Node) StyleGetFlexDirection() FlexDirection {
	return n.getStyle().flexDirection()
}

// StyleSetJustifyContent
func (n *Node) StyleSetJustifyContent(justifyContent Justify) {
	if n.getStyle().justifyContent() != justifyContent {
		n.getStyle().setJustifyContent(justifyContent)
		n.markDirtyAndPropagate()
	}
}

// StyleGetJustifyContent
func (n *Node) StyleGetJustifyContent() Justify {
	return n.getStyle().justifyContent()
}

// StyleSetAlignContent
func (n *Node) StyleSetAlignContent(alignContent Align) {
	if n.getStyle().alignContent() != alignContent {
		n.getStyle().setAlignContent(alignContent)
		n.markDirtyAndPropagate()
	}
}

// StyleGetAlignContent
func (n *Node) StyleGetAlignContent() Align {
	return n.getStyle().alignContent()
}

// StyleSetAlignItems
func (n *Node) StyleSetAlignItems(alignItems Align) {
	if n.getStyle().alignItems() != alignItems {
		n.getStyle().setAlignItems(alignItems)
		n.markDirtyAndPropagate()
	}
}

// StyleGetAlignItems
func (n *Node) StyleGetAlignItems() Align {
	return n.getStyle().alignItems()
}

// StyleSetAlignSelf
func (n *Node) StyleSetAlignSelf(alignSelf Align) {
	if n.getStyle().alignSelf() != alignSelf {
		n.getStyle().setAlignSelf(alignSelf)
		n.markDirtyAndPropagate()
	}
}

// StyleGetAlignSelf
func (n *Node) StyleGetAlignSelf() Align {
	return n.getStyle().alignSelf()
}

// StyleSetFlexWrap
func (n *Node) StyleSetFlexWrap(flexWrap Wrap) {
	if n.getStyle().flexWrap() != flexWrap {
		n.getStyle().setFlexWrap(flexWrap)
		n.markDirtyAndPropagate()
	}
}

// StyleGetFlexWrap
func (n *Node) StyleGetFlexWrap() Wrap {
	return n.getStyle().flexWrap()
}

// StyleSetOverflow
func (n *Node) StyleSetOverflow(overflow Overflow) {
	if n.getStyle().overflow() != overflow {
		n.getStyle().setOverflow(overflow)
		n.markDirtyAndPropagate()
	}
}

// StyleSetDisplay
func (n *Node) StyleSetDisplay(display Display) {
	if n.getStyle().display() != display {
		n.getStyle().setDisplay(display)
		n.markDirtyAndPropagate()
	}
}

// StyleGetDisplay
func (n *Node) StyleGetDisplay() Display {
	return n.getStyle().display()
}

// StyleSetFlex
func (n *Node) StyleSetFlex(flex float32) {
	value := NewFloatOptional(flex)
	if !n.getStyle().flex().equal(value) {
		n.getStyle().flex_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetFlex
func (n *Node) StyleGetFlex() float32 {
	return If(n.getStyle().flex().isUndefined(), Undefined, n.getStyle().flex().unwrap())
}

// StyleSetFlexGrow
func (n *Node) StyleSetFlexGrow(flexGrow float32) {
	value := NewFloatOptional(flexGrow)
	if !n.getStyle().flexGrow().equal(value) {
		n.getStyle().flexGrow_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetFlexGrow
func (n *Node) StyleGetFlexGrow() float32 {
	return If(n.getStyle().flexGrow().isUndefined(), Undefined, n.getStyle().flexGrow().unwrap())
}

// StyleSetFlexShrink
func (n *Node) StyleSetFlexShrink(flexShrink float32) {
	value := NewFloatOptional(flexShrink)
	if !n.getStyle().flexShrink().equal(value) {
		n.getStyle().flexShrink_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetFlexShrink
func (n *Node) StyleGetFlexShrink() float32 {
	return If(n.getStyle().flexShrink().isUndefined(), Undefined, n.getStyle().flexShrink().unwrap())
}

// StyleSetFlexBasis
func (n *Node) StyleSetFlexBasis(flexBasis float32) {
	value := CompactValueOfPoint(flexBasis)
	if !n.getStyle().flexBasis().equal(value) {
		n.getStyle().flexBasis_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleSetFlexBasisPercent
func (n *Node) StyleSetFlexBasisPercent(flexBasis float32) {
	value := CompactValuePercent(flexBasis)
	if !n.getStyle().flexBasis().equal(value) {
		n.getStyle().flexBasis_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleSetFlexBasisAuto
func (n *Node) StyleSetFlexBasisAuto() {
	value := CompactValueOfAuto()
	if !n.getStyle().flexBasis().equal(value) {
		n.getStyle().flexBasis_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetFlexBasis
func (n *Node) StyleGetFlexBasis() YGValue {
	flexBasis := n.getStyle().flexBasis().YGValue()
	if flexBasis.unit == UnitUndefined || flexBasis.unit == UnitAuto {
		flexBasis.value = Undefined
	}
	return flexBasis
}

// StyleSetWidth sets width
func (n *Node) StyleSetWidth(points float32) {
	value := CompactValueOfPoint(points)
	if !n.getStyle().dimension(DimensionWidth).equal(value) {
		n.getStyle().setDimension(DimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetWidthPercent sets width percent
func (n *Node) StyleSetWidthPercent(percent float32) {
	value := CompactValuePercent(percent)
	if !n.getStyle().dimension(DimensionWidth).equal(value) {
		n.getStyle().setDimension(DimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetWidthAuto sets width auto
func (n *Node) StyleSetWidthAuto() {
	value := CompactValueOfAuto()
	if !n.getStyle().dimension(DimensionWidth).equal(value) {
		n.getStyle().setDimension(DimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetWidth returns width
func (n *Node) StyleGetWidth() float32 {
	return n.getLayout().dimension(DimensionWidth)
}

// StyleSetHeight sets height
func (n *Node) StyleSetHeight(height float32) {
	value := CompactValueOfPoint(height)
	if !n.getStyle().dimension(DimensionHeight).equal(value) {
		n.getStyle().setDimension(DimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetHeightPercent sets height percent
func (n *Node) StyleSetHeightPercent(height float32) {
	value := CompactValuePercent(height)
	if !n.getStyle().dimension(DimensionHeight).equal(value) {
		n.getStyle().setDimension(DimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetHeightAuto sets height auto
func (n *Node) StyleSetHeightAuto() {
	value := CompactValueOfAuto()
	if !n.getStyle().dimension(DimensionHeight).equal(value) {
		n.getStyle().setDimension(DimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetHeight returns height
func (n *Node) StyleGetHeight() float32 {
	return n.getLayout().dimension(DimensionHeight)
}

// StyleSetPositionType
func (n *Node) StyleSetPositionType(positionType PositionType) {
	if n.getStyle().positionType() != positionType {
		n.getStyle().setPositionType(positionType)
		n.markDirtyAndPropagate()
	}
}

// StyleGetPositionType
func (n *Node) StyleGetPositionType() PositionType {
	return n.getStyle().positionType()
}

// StyleSetPosition
func (n *Node) StyleSetPosition(edge Edge, position float32) {
	value := CompactValueOfPoint(position)
	if !n.getStyle().position(edge).equal(value) {
		n.getStyle().setPosition(edge, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetPositionPercent
func (n *Node) StyleSetPositionPercent(edge Edge, position float32) {
	value := CompactValuePercent(position)
	if !n.getStyle().position(edge).equal(value) {
		n.getStyle().setPosition(edge, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetPosition
func (n *Node) StyleGetPosition(edge Edge) YGValue {
	position := n.getStyle().position(edge).YGValue()
	if position.unit == UnitUndefined || position.unit == UnitAuto {
		position.value = Undefined
	}
	return position
}

// StyleSetMargin
func (n *Node) StyleSetMargin(edge Edge, margin float32) {
	value := CompactValueOfPoint(margin)
	if !n.getStyle().margin(edge).equal(value) {
		n.getStyle().setMargin(edge, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMarginPercent
func (n *Node) StyleSetMarginPercent(edge Edge, margin float32) {
	value := CompactValuePercent(margin)
	if !n.getStyle().margin(edge).equal(value) {
		n.getStyle().setMargin(edge, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMarginAuto
func (n *Node) StyleSetMarginAuto(edge Edge) {
	value := CompactValueOfAuto()
	if !n.getStyle().margin(edge).equal(value) {
		n.getStyle().setMargin(edge, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetMargin
func (n *Node) StyleGetMargin(edge Edge) YGValue {
	margin := n.getStyle().margin(edge).YGValue()
	if margin.unit == UnitUndefined || margin.unit == UnitAuto {
		margin.value = Undefined
	}
	return margin
}

// StyleSetPadding
func (n *Node) StyleSetPadding(edge Edge, padding float32) {
	value := CompactValueOfPoint(padding)
	if !n.getStyle().padding(edge).equal(value) {
		n.getStyle().setPadding(edge, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetPaddingPercent
func (n *Node) StyleSetPaddingPercent(edge Edge, padding float32) {
	value := CompactValuePercent(padding)
	if !n.getStyle().padding(edge).equal(value) {
		n.getStyle().setPadding(edge, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetPadding
func (n *Node) StyleGetPadding(edge Edge) YGValue {
	padding := n.getStyle().padding(edge).YGValue()
	if padding.unit == UnitUndefined || padding.unit == UnitAuto {
		padding.value = Undefined
	}
	return padding
}

// StyleSetBorder
func (n *Node) StyleSetBorder(edge Edge, border float32) {
	value := CompactValueOfPoint(border)
	if !n.getStyle().border(edge).equal(value) {
		n.getStyle().setBorder(edge, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetBorder
func (n *Node) StyleGetBorder(edge Edge) float32 {
	border := n.getStyle().border(edge)
	if border.isUndefined() || border.isAuto() {
		return Undefined
	}
	return border.YGValue().value
}

// StyleSetGap
func (n *Node) StyleSetGap(gutter Gutter, gapLength float32) {
	value := CompactValueOfPoint(gapLength)
	if !n.getStyle().gap_[gutter].equal(value) {
		n.getStyle().gap_[gutter] = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetGap
func (n *Node) StyleGetGap(gutter Gutter) float32 {
	gap := n.getStyle().gap_[gutter]
	if gap.isUndefined() || gap.isAuto() {
		return Undefined
	}
	return gap.YGValue().value
}

// StyleSetMinWidth
func (n *Node) StyleSetMinWidth(minWidth float32) {
	value := CompactValueOfPoint(minWidth)
	if !n.getStyle().minDimension(DimensionWidth).equal(value) {
		n.getStyle().setMinDimension(DimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMinWidthPercent
func (n *Node) StyleSetMinWidthPercent(minWidth float32) {
	value := CompactValuePercent(minWidth)
	if !n.getStyle().minDimension(DimensionWidth).equal(value) {
		n.getStyle().setMinDimension(DimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetMinWidth
func (n *Node) StyleGetMinWidth() YGValue {
	return n.getStyle().minDimension(DimensionWidth).YGValue()
}

// StyleSetMinHeight
func (n *Node) StyleSetMinHeight(minHeight float32) {
	value := CompactValueOfPoint(minHeight)
	if !n.getStyle().minDimension(DimensionHeight).equal(value) {
		n.getStyle().setMinDimension(DimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMinHeightPercent
func (n *Node) StyleSetMinHeightPercent(minHeight float32) {
	value := CompactValuePercent(minHeight)
	if !n.getStyle().minDimension(DimensionHeight).equal(value) {
		n.getStyle().setMinDimension(DimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetMinHeight
func (n *Node) StyleGetMinHeight() YGValue {
	return n.getStyle().minDimension(DimensionHeight).YGValue()
}

// StyleSetMaxWidth
func (n *Node) StyleSetMaxWidth(maxWidth float32) {
	value := CompactValueOfPoint(maxWidth)
	if !n.getStyle().maxDimension(DimensionWidth).equal(value) {
		n.getStyle().setMaxDimension(DimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMaxWidthPercent
func (n *Node) StyleSetMaxWidthPercent(maxWidth float32) {
	value := CompactValuePercent(maxWidth)
	if !n.getStyle().maxDimension(DimensionWidth).equal(value) {
		n.getStyle().setMaxDimension(DimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetMaxWidth
func (n *Node) StyleGetMaxWidth() YGValue {
	return n.getStyle().maxDimension(DimensionWidth).YGValue()
}

// StyleSetMaxHeight
func (n *Node) StyleSetMaxHeight(maxHeight float32) {
	value := CompactValueOfPoint(maxHeight)
	if !n.getStyle().maxDimension(DimensionHeight).equal(value) {
		n.getStyle().setMaxDimension(DimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMaxHeightPercent
func (n *Node) StyleSetMaxHeightPercent(maxHeight float32) {
	value := CompactValuePercent(maxHeight)
	if !n.getStyle().maxDimension(DimensionHeight).equal(value) {
		n.getStyle().setMaxDimension(DimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetMaxHeight
func (n *Node) StyleGetMaxHeight() YGValue {
	return n.getStyle().maxDimension(DimensionHeight).YGValue()
}

// StyleSetAspectRatio
func (n *Node) StyleSetAspectRatio(aspectRatio float32) {
	value := NewFloatOptional(aspectRatio)
	if !n.getStyle().aspectRatio().equal(value) {
		n.getStyle().aspectRatio_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetAspectRatio
func (n *Node) StyleGetAspectRatio() float32 {
	return If(n.getStyle().aspectRatio().isUndefined(), Undefined, n.getStyle().aspectRatio().unwrap())
}
