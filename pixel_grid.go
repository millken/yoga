package yoga

import "math"

func roundValueToPixelGrid(value float64, pointScaleFactor float64, forceCeil bool, forceFloor bool) float32 {
	scaledValue := value * pointScaleFactor
	fractial := math.Mod(scaledValue, 1.0)
	if fractial < 0.00001 {
		fractial++
	}
	if inexactEqual(fractial, 0) {
		scaledValue = scaledValue - fractial
	} else if inexactEqual(fractial, 1) {
		scaledValue = scaledValue - fractial + 1
	} else if forceCeil {
		scaledValue = scaledValue - fractial + 1
	} else if forceFloor {
		scaledValue = scaledValue - fractial
	} else {
		scaledValue = scaledValue - fractial +
			If(!math.IsNaN(fractial) &&
				(fractial > 0.5 || inexactEqual(fractial, 0.5)), 1.0, 0.0)
	}
	return If(math.IsNaN(scaledValue) || math.IsNaN(pointScaleFactor), YGUndefined, float32(scaledValue/pointScaleFactor))
}

func roundLayoutResultsToPixelGrid(node *YGNode, absoluteLeft float64, absoluteTop float64) {
	pointScaleFactor := float64(node.getConfig().getPointScaleFactor())
	nodeLeft := float64(node.getLayout().position[YGEdgeLeft])
	nodeTop := float64(node.getLayout().position[YGEdgeTop])

	nodeWidth := float64(node.getLayout().dimension(YGDimensionWidth))
	nodeHeight := float64(node.getLayout().dimension(YGDimensionHeight))

	absoluteNodeLeft := absoluteLeft + nodeLeft
	absoluteNodeTop := absoluteTop + nodeTop

	absoluteNodeRight := absoluteNodeLeft + nodeWidth
	absoluteNodeBottom := absoluteNodeTop + nodeHeight

	if pointScaleFactor != 0.0 {
		// If a node has a custom measure function we never want to round down its
		// size as this could lead to unwanted text truncation.
		textRounding := node.getNodeType() == YGNodeTypeText

		node.setLayoutPosition(roundValueToPixelGrid(nodeLeft, pointScaleFactor, false, textRounding), YGEdgeLeft)

		node.setLayoutPosition(roundValueToPixelGrid(nodeTop, pointScaleFactor, false, textRounding), YGEdgeTop)

		// We multiply dimension by scale factor and if the result is close to the
		// whole number, we don't have any fraction To verify if the result is close
		// to whole number we want to check both floor and ceil numbers
		hasFractionalWidth := !inexactEqual(math.Mod(nodeWidth*pointScaleFactor, 1.0), 0) &&
			!inexactEqual(math.Mod(nodeWidth*pointScaleFactor, 1.0), 1.0)
		hasFractionalHeight := !inexactEqual(math.Mod(nodeHeight*pointScaleFactor, 1.0), 0) &&
			!inexactEqual(math.Mod(nodeHeight*pointScaleFactor, 1.0), 1.0)

		node.setLayoutDimension(
			roundValueToPixelGrid(
				absoluteNodeRight,
				pointScaleFactor,
				(textRounding && hasFractionalWidth),
				(textRounding && !hasFractionalWidth))-
				roundValueToPixelGrid(
					absoluteNodeLeft, pointScaleFactor, false, textRounding), YGDimensionWidth)

		node.setLayoutDimension(
			roundValueToPixelGrid(
				absoluteNodeBottom,
				pointScaleFactor,
				(textRounding && hasFractionalHeight),
				(textRounding && !hasFractionalHeight))-
				roundValueToPixelGrid(
					absoluteNodeTop, pointScaleFactor, false, textRounding),
			YGDimensionHeight)
	}

	for _, child := range node.getChildren() {
		roundLayoutResultsToPixelGrid(child, absoluteNodeLeft, absoluteNodeTop)
	}
}