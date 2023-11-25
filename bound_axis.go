package yoga

func paddingAndBorderForAxis(node *Node, axis YGFlexDirection, widthSize float32) float32 {
	// The total padding/border for a given axis does not depend on the direction
	// so hardcoding LTR here to avoid piping direction to this function
	return node.getInlineStartPaddingAndBorder(axis, YGDirectionLTR, widthSize) +
		node.getInlineEndPaddingAndBorder(axis, YGDirectionLTR, widthSize)
}

func boundAxisWithinMinAndMax(node *Node, axis YGFlexDirection, value FloatOptional, axisSize float32) FloatOptional {
	var min, max FloatOptional

	if isColumn(axis) {
		min = resolveValue(
			node.getStyle().minDimension(YGDimensionHeight).YGValue(), axisSize)
		max = resolveValue(
			node.getStyle().maxDimension(YGDimensionHeight).YGValue(), axisSize)
	} else if isRow(axis) {
		min = resolveValue(
			node.getStyle().minDimension(YGDimensionWidth).YGValue(), axisSize)
		max = resolveValue(
			node.getStyle().maxDimension(YGDimensionWidth).YGValue(), axisSize)
	}

	if max.unwrap() >= 0 && value.unwrap() > max.unwrap() {
		return max
	}

	if min.unwrap() >= 0 && value.unwrap() < min.unwrap() {
		return min
	}

	return value
}

// Like boundAxisWithinMinAndMax but also ensures that the value doesn't
// go below the padding and border amount.
func boundAxis(
	node *Node,
	axis YGFlexDirection,
	value float32,
	axisSize float32,
	widthSize float32,
) float32 {
	return maxOrDefined(
		boundAxisWithinMinAndMax(node, axis, FloatOptional{value}, axisSize).unwrap(),
		paddingAndBorderForAxis(node, axis, widthSize),
	)
}
