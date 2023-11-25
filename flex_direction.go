package yoga

func isRow(flexDirection YGFlexDirection) bool {
	return flexDirection == YGFlexDirectionRow || flexDirection == YGFlexDirectionRowReverse
}

func isColumn(flexDirection YGFlexDirection) bool {
	return flexDirection == YGFlexDirectionColumn || flexDirection == YGFlexDirectionColumnReverse
}

func resolveDirection(flexDirection YGFlexDirection, direction YGDirection) YGFlexDirection {
	if direction == YGDirectionRTL {
		if flexDirection == YGFlexDirectionRow {
			return YGFlexDirectionRowReverse
		} else if flexDirection == YGFlexDirectionRowReverse {
			return YGFlexDirectionRow
		}
	}
	return flexDirection
}

func resolveCrossDirection(flexDirection YGFlexDirection, direction YGDirection) YGFlexDirection {
	return If(isColumn(flexDirection), resolveDirection(YGFlexDirectionRow, direction), YGFlexDirectionColumn)
}

func flexStartEdge(flexDirection YGFlexDirection) YGEdge {
	switch flexDirection {
	case YGFlexDirectionColumn:
		return YGEdgeTop
	case YGFlexDirectionColumnReverse:
		return YGEdgeBottom
	case YGFlexDirectionRow:
		return YGEdgeLeft
	case YGFlexDirectionRowReverse:
		return YGEdgeRight
	}
	panic("Invalid FlexDirection")
}

func flexEndEdge(flexDirection YGFlexDirection) YGEdge {
	switch flexDirection {
	case YGFlexDirectionColumn:
		return YGEdgeBottom
	case YGFlexDirectionColumnReverse:
		return YGEdgeTop
	case YGFlexDirectionRow:
		return YGEdgeRight
	case YGFlexDirectionRowReverse:
		return YGEdgeLeft
	}
	panic("Invalid FlexDirection")
}

func inlineStartEdge(flexDirection YGFlexDirection, direction YGDirection) YGEdge {
	if isRow(flexDirection) {
		return If(direction == YGDirectionRTL, YGEdgeRight, YGEdgeLeft)
	}
	return YGEdgeTop
}

func inlineEndEdge(flexDirection YGFlexDirection, direction YGDirection) YGEdge {
	if isRow(flexDirection) {
		return If(direction == YGDirectionRTL, YGEdgeLeft, YGEdgeRight)
	}
	return YGEdgeBottom
}

func flexStartRelativeEdge(flexDirection YGFlexDirection, direction YGDirection) YGEdge {
	leadLayoutEdge := inlineStartEdge(flexDirection, direction)
	leadFlexItemEdge := flexStartEdge(flexDirection)
	return If(leadLayoutEdge == leadFlexItemEdge, YGEdgeStart, YGEdgeEnd)
}

func flexEndRelativeEdge(flexDirection YGFlexDirection, direction YGDirection) YGEdge {
	trailLayoutEdge := inlineEndEdge(flexDirection, direction)
	trailFlexItemEdge := flexEndEdge(flexDirection)
	return If(trailLayoutEdge == trailFlexItemEdge, YGEdgeEnd, YGEdgeStart)
}

func dimension(flexDirection YGFlexDirection) YGDimension {
	switch flexDirection {
	case YGFlexDirectionColumn, YGFlexDirectionColumnReverse:
		return YGDimensionHeight
	case YGFlexDirectionRow, YGFlexDirectionRowReverse:
		return YGDimensionWidth
	}
	panic("Invalid FlexDirection")
}
