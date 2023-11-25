package yoga

/*
YG_EXPORT void YGNodeCopyStyle(YGNodeRef dstNode, YGNodeConstRef srcNode);

YG_EXPORT void YGNodeStyleSetDirection(YGNodeRef node, YGDirection direction);
YG_EXPORT YGDirection YGNodeStyleGetDirection(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetFlexDirection(
    YGNodeRef node,
    YGFlexDirection flexDirection);
YG_EXPORT YGFlexDirection YGNodeStyleGetFlexDirection(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetJustifyContent(
    YGNodeRef node,
    YGJustify justifyContent);
YG_EXPORT YGJustify YGNodeStyleGetJustifyContent(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetAlignContent(YGNodeRef node, YGAlign alignContent);
YG_EXPORT YGAlign YGNodeStyleGetAlignContent(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetAlignItems(YGNodeRef node, YGAlign alignItems);
YG_EXPORT YGAlign YGNodeStyleGetAlignItems(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetAlignSelf(YGNodeRef node, YGAlign alignSelf);
YG_EXPORT YGAlign YGNodeStyleGetAlignSelf(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetPositionType(
    YGNodeRef node,
    YGPositionType positionType);
YG_EXPORT YGPositionType YGNodeStyleGetPositionType(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetFlexWrap(YGNodeRef node, YGWrap flexWrap);
YG_EXPORT YGWrap YGNodeStyleGetFlexWrap(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetOverflow(YGNodeRef node, YGOverflow overflow);
YG_EXPORT YGOverflow YGNodeStyleGetOverflow(YGNodeConstRef node);

YG_EXPORT void YGNodeStyleSetDisplay(YGNodeRef node, YGDisplay display);
YG_EXPORT YGDisplay YGNodeStyleGetDisplay(YGNodeConstRef node);

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
YGNodeStyleSetPosition(YGNodeRef node, YGEdge edge, float position);
YG_EXPORT void
YGNodeStyleSetPositionPercent(YGNodeRef node, YGEdge edge, float position);
YG_EXPORT YGValue YGNodeStyleGetPosition(YGNodeConstRef node, YGEdge edge);

YG_EXPORT void YGNodeStyleSetMargin(YGNodeRef node, YGEdge edge, float margin);
YG_EXPORT void
YGNodeStyleSetMarginPercent(YGNodeRef node, YGEdge edge, float margin);
YG_EXPORT void YGNodeStyleSetMarginAuto(YGNodeRef node, YGEdge edge);
YG_EXPORT YGValue YGNodeStyleGetMargin(YGNodeConstRef node, YGEdge edge);

YG_EXPORT void
YGNodeStyleSetPadding(YGNodeRef node, YGEdge edge, float padding);
YG_EXPORT void
YGNodeStyleSetPaddingPercent(YGNodeRef node, YGEdge edge, float padding);
YG_EXPORT YGValue YGNodeStyleGetPadding(YGNodeConstRef node, YGEdge edge);

YG_EXPORT void YGNodeStyleSetBorder(YGNodeRef node, YGEdge edge, float border);
YG_EXPORT float YGNodeStyleGetBorder(YGNodeConstRef node, YGEdge edge);

YG_EXPORT void
YGNodeStyleSetGap(YGNodeRef node, YGGutter gutter, float gapLength);
YG_EXPORT float YGNodeStyleGetGap(YGNodeConstRef node, YGGutter gutter);

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

func updateStyle[T float32 | YGDirection](node *Node, prop string, value T) {

}

// StyleSetDirection
func (n *Node) StyleSetDirection(direction YGDirection) {
	if n.getStyle().direction() != direction {
		setEnumData(&n.style_.flags, directionOffset, direction, uint8(direction))
		n.markDirtyAndPropagate()
	}
}

// StyleGetDirection
func (n *Node) StyleGetDirection() YGDirection {
	return n.getStyle().direction()
}

// StyleSetFlexDirection
func (n *Node) StyleSetFlexDirection(flexDirection YGFlexDirection) {
	if n.getStyle().flexDirection() != flexDirection {
		setEnumData(&n.style_.flags, flexDirectionOffset, flexDirection, uint8(flexDirection))
		n.markDirtyAndPropagate()
	}
}

// StyleGetFlexDirection
func (n *Node) StyleGetFlexDirection() YGFlexDirection {
	return n.getStyle().flexDirection()
}

// StyleSetJustifyContent
func (n *Node) StyleSetJustifyContent(justifyContent YGJustify) {
	if n.getStyle().justifyContent() != justifyContent {
		setEnumData(&n.style_.flags, justifyContentOffset, justifyContent, uint8(justifyContent))
		n.markDirtyAndPropagate()
	}
}

// StyleGetJustifyContent
func (n *Node) StyleGetJustifyContent() YGJustify {
	return n.getStyle().justifyContent()
}

// StyleSetAlignContent
func (n *Node) StyleSetAlignContent(alignContent YGAlign) {
	if n.getStyle().alignContent() != alignContent {
		setEnumData(&n.style_.flags, alignContentOffset, alignContent, uint8(alignContent))
		n.markDirtyAndPropagate()
	}
}

// StyleGetAlignContent
func (n *Node) StyleGetAlignContent() YGAlign {
	return n.getStyle().alignContent()
}

// StyleSetAlignItems
func (n *Node) StyleSetAlignItems(alignItems YGAlign) {
	if n.getStyle().alignItems() != alignItems {
		setEnumData(&n.style_.flags, alignItemsOffset, alignItems, uint8(alignItems))
		n.markDirtyAndPropagate()
	}
}

// StyleGetAlignItems
func (n *Node) StyleGetAlignItems() YGAlign {
	return n.getStyle().alignItems()
}

// StyleSetAlignSelf
func (n *Node) StyleSetAlignSelf(alignSelf YGAlign) {
	if n.getStyle().alignSelf() != alignSelf {
		setEnumData(&n.style_.flags, alignSelfOffset, alignSelf, uint8(alignSelf))
		n.markDirtyAndPropagate()
	}
}

// StyleGetAlignSelf
func (n *Node) StyleGetAlignSelf() YGAlign {
	return n.getStyle().alignSelf()
}

// StyleSetFlexWrap
func (n *Node) StyleSetFlexWrap(flexWrap YGWrap) {
	if n.getStyle().flexWrap() != flexWrap {
		setEnumData(&n.style_.flags, flexWrapOffset, flexWrap, uint8(flexWrap))
		n.markDirtyAndPropagate()
	}
}

// StyleGetFlexWrap
func (n *Node) StyleGetFlexWrap() YGWrap {
	return n.getStyle().flexWrap()
}

// StyleSetOverflow
func (n *Node) StyleSetOverflow(overflow YGOverflow) {
	if n.getStyle().overflow() != overflow {
		setEnumData(&n.style_.flags, overflowOffset, overflow, uint8(overflow))
		n.markDirtyAndPropagate()
	}
}

// StyleSetDisplay
func (n *Node) StyleSetDisplay(display YGDisplay) {
	if n.getStyle().display() != display {
		setEnumData(&n.style_.flags, displayOffset, display, uint8(display))
		n.markDirtyAndPropagate()
	}
}

// StyleGetDisplay
func (n *Node) StyleGetDisplay() YGDisplay {
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
	return If(n.getStyle().flex().isUndefined(), YGUndefined, n.getStyle().flex().unwrap())
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
	return If(n.getStyle().flexGrow().isUndefined(), YGUndefined, n.getStyle().flexGrow().unwrap())
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
	return If(n.getStyle().flexShrink().isUndefined(), YGUndefined, n.getStyle().flexShrink().unwrap())
}

// StyleSetFlexBasis
func (n *Node) StyleSetFlexBasis(flexBasis float32) {
	value := CompactValueOfMaybe(YGUnitPoint, flexBasis)
	if !n.getStyle().flexBasis().equal(value) {
		n.getStyle().flexBasis_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleSetFlexBasisPercent
func (n *Node) StyleSetFlexBasisPercent(flexBasis float32) {
	value := CompactValueOfMaybe(YGUnitPercent, flexBasis)
	if !n.getStyle().flexBasis().equal(value) {
		n.getStyle().flexBasis_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleSetFlexBasisAuto
func (n *Node) StyleSetFlexBasisAuto() {
	value := CompactValueOfMaybe(YGUnitAuto, 0.0)
	if !n.getStyle().flexBasis().equal(value) {
		n.getStyle().flexBasis_ = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetFlexBasis
func (n *Node) StyleGetFlexBasis() YGValue {
	flexBasis := n.getStyle().flexBasis().YGValue()
	if flexBasis.unit == YGUnitUndefined || flexBasis.unit == YGUnitAuto {
		flexBasis.value = YGUndefined
	}
	return flexBasis
}

// StyleSetWidth sets width
func (n *Node) StyleSetWidth(points float32) {
	value := CompactValueOfMaybe(YGUnitPoint, points)
	if !n.getStyle().dimension(YGDimensionWidth).equal(value) {
		n.getStyle().setDimension(YGDimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetWidthPercent sets width percent
func (n *Node) StyleSetWidthPercent(percent float32) {
	value := CompactValueOfMaybe(YGUnitPercent, percent)
	if !n.getStyle().dimension(YGDimensionWidth).equal(value) {
		n.getStyle().setDimension(YGDimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetWidthAuto sets width auto
func (n *Node) StyleSetWidthAuto() {
	value := CompactValueOfMaybe(YGUnitAuto, 0.0)
	if !n.getStyle().dimension(YGDimensionWidth).equal(value) {
		n.getStyle().setDimension(YGDimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetWidth returns width
func (n *Node) StyleGetWidth() float32 {
	return n.getLayout().dimension(YGDimensionWidth)
}

// StyleSetHeight sets height
func (n *Node) StyleSetHeight(height float32) {
	value := CompactValueOfMaybe(YGUnitPoint, height)
	if !n.getStyle().dimension(YGDimensionHeight).equal(value) {
		n.getStyle().setDimension(YGDimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetHeightPercent sets height percent
func (n *Node) StyleSetHeightPercent(height float32) {
	value := CompactValueOfMaybe(YGUnitPercent, height)
	if !n.getStyle().dimension(YGDimensionHeight).equal(value) {
		n.getStyle().setDimension(YGDimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetHeightAuto sets height auto
func (n *Node) StyleSetHeightAuto() {
	value := CompactValueOfMaybe(YGUnitAuto, 0.0)
	if !n.getStyle().dimension(YGDimensionHeight).equal(value) {
		n.getStyle().setDimension(YGDimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetHeight returns height
func (n *Node) StyleGetHeight() float32 {
	return n.getLayout().dimension(YGDimensionHeight)
}

// StyleSetPositionType
func (n *Node) StyleSetPositionType(positionType YGPositionType) {
	if n.getStyle().positionType() != positionType {
		setEnumData(&n.style_.flags, positionTypeOffset, positionType, uint8(positionType))
		n.markDirtyAndPropagate()
	}
}

// StyleGetPositionType
func (n *Node) StyleGetPositionType() YGPositionType {
	return n.getStyle().positionType()
}

// StyleSetPosition
func (n *Node) StyleSetPosition(edge YGEdge, position float32) {
	value := CompactValueOfMaybe(YGUnitPoint, position)
	if !n.getStyle().position()[edge].equal(value) {
		n.getStyle().position_[edge] = value
		n.markDirtyAndPropagate()
	}
}

// StyleSetPositionPercent
func (n *Node) StyleSetPositionPercent(edge YGEdge, position float32) {
	value := CompactValueOfMaybe(YGUnitPercent, position)
	if !n.getStyle().position()[edge].equal(value) {
		n.getStyle().position_[edge] = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetPosition
func (n *Node) StyleGetPosition(edge YGEdge) YGValue {
	position := n.getStyle().position()[edge].YGValue()
	if position.unit == YGUnitUndefined || position.unit == YGUnitAuto {
		position.value = YGUndefined
	}
	return position
}

// StyleSetMargin
func (n *Node) StyleSetMargin(edge YGEdge, margin float32) {
	value := CompactValueOfMaybe(YGUnitPoint, margin)
	if !n.getStyle().margin()[edge].equal(value) {
		n.getStyle().margin_[edge] = value
		n.markDirtyAndPropagate()
	}
}

// StyleSetMarginPercent
func (n *Node) StyleSetMarginPercent(edge YGEdge, margin float32) {
	value := CompactValueOfMaybe(YGUnitPercent, margin)
	if !n.getStyle().margin()[edge].equal(value) {
		n.getStyle().margin_[edge] = value
		n.markDirtyAndPropagate()
	}
}

// StyleSetMarginAuto
func (n *Node) StyleSetMarginAuto(edge YGEdge) {
	value := CompactValueOfMaybe(YGUnitAuto, 0.0)
	if !n.getStyle().margin()[edge].equal(value) {
		n.getStyle().margin_[edge] = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetMargin
func (n *Node) StyleGetMargin(edge YGEdge) YGValue {
	margin := n.getStyle().margin()[edge].YGValue()
	if margin.unit == YGUnitUndefined || margin.unit == YGUnitAuto {
		margin.value = YGUndefined
	}
	return margin
}

// StyleSetPadding
func (n *Node) StyleSetPadding(edge YGEdge, padding float32) {
	value := CompactValueOfMaybe(YGUnitPoint, padding)
	if !n.getStyle().padding()[edge].equal(value) {
		n.getStyle().padding_[edge] = value
		n.markDirtyAndPropagate()
	}
}

// StyleSetPaddingPercent
func (n *Node) StyleSetPaddingPercent(edge YGEdge, padding float32) {
	value := CompactValueOfMaybe(YGUnitPercent, padding)
	if !n.getStyle().padding()[edge].equal(value) {
		n.getStyle().padding_[edge] = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetPadding
func (n *Node) StyleGetPadding(edge YGEdge) YGValue {
	padding := n.getStyle().padding()[edge].YGValue()
	if padding.unit == YGUnitUndefined || padding.unit == YGUnitAuto {
		padding.value = YGUndefined
	}
	return padding
}

// StyleSetBorder
func (n *Node) StyleSetBorder(edge YGEdge, border float32) {
	value := CompactValueOfMaybe(YGUnitPoint, border)
	if !n.getStyle().border()[edge].equal(value) {
		n.getStyle().border_[edge] = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetBorder
func (n *Node) StyleGetBorder(edge YGEdge) float32 {
	border := n.getStyle().border()[edge]
	if border.isUndefined() || border.isAuto() {
		return YGUndefined
	}
	return border.YGValue().value
}

// StyleSetGap
func (n *Node) StyleSetGap(gutter YGGutter, gapLength float32) {
	value := CompactValueOfMaybe(YGUnitPoint, gapLength)
	if !n.getStyle().gap_[gutter].equal(value) {
		n.getStyle().gap_[gutter] = value
		n.markDirtyAndPropagate()
	}
}

// StyleGetGap
func (n *Node) StyleGetGap(gutter YGGutter) float32 {
	gap := n.getStyle().gap_[gutter]
	if gap.isUndefined() || gap.isAuto() {
		return YGUndefined
	}
	return gap.YGValue().value
}

// StyleSetMinWidth
func (n *Node) StyleSetMinWidth(minWidth float32) {
	value := CompactValueOfMaybe(YGUnitPoint, minWidth)
	if !n.getStyle().minDimension(YGDimensionWidth).equal(value) {
		n.getStyle().setMinDimension(YGDimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMinWidthPercent
func (n *Node) StyleSetMinWidthPercent(minWidth float32) {
	value := CompactValueOfMaybe(YGUnitPercent, minWidth)
	if !n.getStyle().minDimension(YGDimensionWidth).equal(value) {
		n.getStyle().setMinDimension(YGDimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetMinWidth
func (n *Node) StyleGetMinWidth() YGValue {
	return n.getStyle().minDimension(YGDimensionWidth).YGValue()
}

// StyleSetMinHeight
func (n *Node) StyleSetMinHeight(minHeight float32) {
	value := CompactValueOfMaybe(YGUnitPoint, minHeight)
	if !n.getStyle().minDimension(YGDimensionHeight).equal(value) {
		n.getStyle().setMinDimension(YGDimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMinHeightPercent
func (n *Node) StyleSetMinHeightPercent(minHeight float32) {
	value := CompactValueOfMaybe(YGUnitPercent, minHeight)
	if !n.getStyle().minDimension(YGDimensionHeight).equal(value) {
		n.getStyle().setMinDimension(YGDimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetMinHeight
func (n *Node) StyleGetMinHeight() YGValue {
	return n.getStyle().minDimension(YGDimensionHeight).YGValue()
}

// StyleSetMaxWidth
func (n *Node) StyleSetMaxWidth(maxWidth float32) {
	value := CompactValueOfMaybe(YGUnitPoint, maxWidth)
	if !n.getStyle().maxDimension(YGDimensionWidth).equal(value) {
		n.getStyle().setMaxDimension(YGDimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMaxWidthPercent
func (n *Node) StyleSetMaxWidthPercent(maxWidth float32) {
	value := CompactValueOfMaybe(YGUnitPercent, maxWidth)
	if !n.getStyle().maxDimension(YGDimensionWidth).equal(value) {
		n.getStyle().setMaxDimension(YGDimensionWidth, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetMaxWidth
func (n *Node) StyleGetMaxWidth() YGValue {
	return n.getStyle().maxDimension(YGDimensionWidth).YGValue()
}

// StyleSetMaxHeight
func (n *Node) StyleSetMaxHeight(maxHeight float32) {
	value := CompactValueOfMaybe(YGUnitPoint, maxHeight)
	if !n.getStyle().maxDimension(YGDimensionHeight).equal(value) {
		n.getStyle().setMaxDimension(YGDimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleSetMaxHeightPercent
func (n *Node) StyleSetMaxHeightPercent(maxHeight float32) {
	value := CompactValueOfMaybe(YGUnitPercent, maxHeight)
	if !n.getStyle().maxDimension(YGDimensionHeight).equal(value) {
		n.getStyle().setMaxDimension(YGDimensionHeight, value)
		n.markDirtyAndPropagate()
	}
}

// StyleGetMaxHeight
func (n *Node) StyleGetMaxHeight() YGValue {
	return n.getStyle().maxDimension(YGDimensionHeight).YGValue()
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
	return If(n.getStyle().aspectRatio().isUndefined(), YGUndefined, n.getStyle().aspectRatio().unwrap())
}
