package yoga

func calculateBaseline(node *Node) float32 {
	if node.hasBaselineFunc() {
		//Event.Publish(Event.NodeBaselineStart, node)

		baseline := node.baseline(
			node.getLayout().measuredDimension(DimensionWidth),
			node.getLayout().measuredDimension(DimensionHeight),
		)

		//Event.Publish(Event.NodeBaselineEnd, node)
		if IsNaN(baseline) {
			panic("Expect custom baseline function to not return NaN")
		}
		return baseline
	}

	var baselineChild *Node
	childCount := node.getChildCount()
	for i := uint32(0); i < childCount; i++ {
		child := node.getChild(i)
		if child.getLineIndex() > 0 {
			break
		}
		if child.getStyle().positionType() == YGPositionTypeAbsolute {
			continue
		}
		if resolveChildAlignment(node, child) == AlignBaseline ||
			child.isReferenceBaseline() {
			baselineChild = child
			break
		}

		if baselineChild == nil {
			baselineChild = child
		}
	}

	if baselineChild == nil {
		return node.getLayout().measuredDimension(DimensionHeight)
	}

	baseline := calculateBaseline(baselineChild)
	return baseline + baselineChild.getLayout().position[EdgeTop]
}

func isBaselineLayout(node *Node) bool {
	if isColumn(node.getStyle().flexDirection()) {
		return false
	}
	if node.getStyle().alignItems() == AlignBaseline {
		return true
	}
	childCount := node.getChildCount()
	for i := uint32(0); i < childCount; i++ {
		child := node.getChild(i)
		if child.getStyle().positionType() != YGPositionTypeAbsolute &&
			child.getStyle().alignSelf() == AlignBaseline {
			return true
		}
	}

	return false
}
