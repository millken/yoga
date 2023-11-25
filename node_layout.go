package yoga

/*
YG_EXPORT float YGNodeLayoutGetLeft(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetTop(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetRight(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetBottom(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetWidth(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetHeight(YGNodeConstRef node);
YG_EXPORT YGDirection YGNodeLayoutGetDirection(YGNodeConstRef node);
YG_EXPORT bool YGNodeLayoutGetHadOverflow(YGNodeConstRef node);

// Get the computed values for these nodes after performing layout. If they were
// set using point values then the returned value will be the same as
// YGNodeStyleGetXXX. However if they were set using a percentage value then the
// returned value is the computed value used during layout.
YG_EXPORT float YGNodeLayoutGetMargin(YGNodeConstRef node, YGEdge edge);
YG_EXPORT float YGNodeLayoutGetBorder(YGNodeConstRef node, YGEdge edge);
YG_EXPORT float YGNodeLayoutGetPadding(YGNodeConstRef node, YGEdge edge);
*/

// LayoutLeft returns left
func (n *Node) LayoutLeft() float32 {
	return n.getLayout().position[YGEdgeLeft]
}

// LayoutTop
func (n *Node) LayoutTop() float32 {
	return n.getLayout().position[YGEdgeTop]
}

// LayoutRight
func (n *Node) LayoutRight() float32 {
	return n.getLayout().position[YGEdgeRight]
}

// LayoutBottom
func (n *Node) LayoutBottom() float32 {
	return n.getLayout().position[YGEdgeBottom]
}

// LayoutWidth returns width
func (n *Node) LayoutWidth() float32 {
	return n.getLayout().dimension(YGDimensionWidth)
}

// LayoutHeight returns height
func (n *Node) LayoutHeight() float32 {
	return n.getLayout().dimension(YGDimensionHeight)
}

// LayoutDirection
func (n *Node) LayoutDirection() YGDirection {
	return n.getLayout().direction()
}

// LayoutHadOverflow
func (n *Node) LayoutHadOverflow() bool {
	return n.getLayout().hadOverflow()
}

// LayoutMargin
func (n *Node) LayoutMargin(edge YGEdge) float32 {
	member := n.getLayout().margin
	if edge > YGEdgeEnd {
		panic("Cannot get layout properties of multi-edge shorthands")
	}
	if edge == YGEdgeStart {
		if n.getLayout().direction() == YGDirectionRTL {
			return member[YGEdgeRight]
		} else {
			return member[YGEdgeLeft]
		}
	}

	if edge == YGEdgeEnd {
		if n.getLayout().direction() == YGDirectionRTL {
			return member[YGEdgeLeft]
		} else {
			return member[YGEdgeRight]
		}
	}

	return member[edge]
}

// LayoutBorder
func (n *Node) LayoutBorder(edge YGEdge) float32 {
	member := n.getLayout().border
	if edge > YGEdgeEnd {
		panic("Cannot get layout properties of multi-edge shorthands")
	}
	if edge == YGEdgeStart {
		if n.getLayout().direction() == YGDirectionRTL {
			return member[YGEdgeRight]
		} else {
			return member[YGEdgeLeft]
		}
	}

	if edge == YGEdgeEnd {
		if n.getLayout().direction() == YGDirectionRTL {
			return member[YGEdgeLeft]
		} else {
			return member[YGEdgeRight]
		}
	}

	return member[edge]
}

// LayoutPadding
func (n *Node) LayoutPadding(edge YGEdge) float32 {
	member := n.getLayout().padding
	if edge > YGEdgeEnd {
		panic("Cannot get layout properties of multi-edge shorthands")
	}
	if edge == YGEdgeStart {
		if n.getLayout().direction() == YGDirectionRTL {
			return member[YGEdgeRight]
		} else {
			return member[YGEdgeLeft]
		}
	}

	if edge == YGEdgeEnd {
		if n.getLayout().direction() == YGDirectionRTL {
			return member[YGEdgeLeft]
		} else {
			return member[YGEdgeRight]
		}
	}

	return member[edge]
}
