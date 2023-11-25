package yoga

func calculateBaseline(node *YGNode) float32 {
	if node.hasBaselineFunc() {
		//Event.Publish(Event.NodeBaselineStart, node)

		baseline := node.baseline(
			node.getLayout().measuredDimension(YGDimensionWidth),
			node.getLayout().measuredDimension(YGDimensionHeight),
		)

		//Event.Publish(Event.NodeBaselineEnd, node)
		if IsNaN(baseline) {
			panic("Expect custom baseline function to not return NaN")
		}
		return baseline
	}

	var baselineChild *YGNode
	childCount := node.getChildCount()
	for i := uint32(0); i < childCount; i++ {
		child := node.getChild(i)
		if child.getLineIndex() > 0 {
			break
		}
		if child.getStyle().positionType() == YGPositionTypeAbsolute {
			continue
		}
		if resolveChildAlignment(node, child) == YGAlignBaseline ||
			child.isReferenceBaseline() {
			baselineChild = child
			break
		}

		if baselineChild == nil {
			baselineChild = child
		}
	}

	if baselineChild == nil {
		return node.getLayout().measuredDimension(YGDimensionHeight)
	}

	baseline := calculateBaseline(baselineChild)
	return baseline + baselineChild.getLayout().position[YGEdgeTop]
}

func isBaselineLayout(node *YGNode) bool {
	if isColumn(node.getStyle().flexDirection()) {
		return false
	}
	if node.getStyle().alignItems() == YGAlignBaseline {
		return true
	}
	childCount := node.getChildCount()
	for i := uint32(0); i < childCount; i++ {
		child := node.getChild(i)
		if child.getStyle().positionType() != YGPositionTypeAbsolute &&
			child.getStyle().alignSelf() == YGAlignBaseline {
			return true
		}
	}

	return false
}
