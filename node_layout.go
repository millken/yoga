package yoga

/*
YG_EXPORT float YGNodeLayoutGetLeft(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetTop(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetRight(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetBottom(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetWidth(YGNodeConstRef node);
YG_EXPORT float YGNodeLayoutGetHeight(YGNodeConstRef node);
YG_EXPORT Direction YGNodeLayoutGetDirection(YGNodeConstRef node);
YG_EXPORT bool YGNodeLayoutGetHadOverflow(YGNodeConstRef node);

// Get the computed values for these nodes after performing layout. If they were
// set using point values then the returned value will be the same as
// YGNodeStyleGetXXX. However if they were set using a percentage value then the
// returned value is the computed value used during layout.
YG_EXPORT float YGNodeLayoutGetMargin(YGNodeConstRef node, Edge edge);
YG_EXPORT float YGNodeLayoutGetBorder(YGNodeConstRef node, Edge edge);
YG_EXPORT float YGNodeLayoutGetPadding(YGNodeConstRef node, Edge edge);
*/

// LayoutLeft returns left
func (n *Node) LayoutLeft() float32 {
	return n.getLayout().position(EdgeLeft)
}

// LayoutTop
func (n *Node) LayoutTop() float32 {
	return n.getLayout().position(EdgeTop)
}

// LayoutRight
func (n *Node) LayoutRight() float32 {
	return n.getLayout().position(EdgeRight)
}

// LayoutBottom
func (n *Node) LayoutBottom() float32 {
	return n.getLayout().position(EdgeBottom)
}

// LayoutWidth returns width
func (n *Node) LayoutWidth() float32 {
	return n.getLayout().dimension(DimensionWidth)
}

// LayoutHeight returns height
func (n *Node) LayoutHeight() float32 {
	return n.getLayout().dimension(DimensionHeight)
}

// LayoutDirection
func (n *Node) LayoutDirection() Direction {
	return n.getLayout().direction()
}

// LayoutHadOverflow
func (n *Node) LayoutHadOverflow() bool {
	return n.getLayout().hadOverflow()
}

// LayoutMargin
func (n *Node) LayoutMargin(edge Edge) float32 {
	member := n.getLayout().margin_
	if edge > EdgeEnd {
		panic("Cannot get layout properties of multi-edge shorthands")
	}
	if edge == EdgeStart {
		if n.getLayout().direction() == DirectionRTL {
			return member[EdgeRight]
		} else {
			return member[EdgeLeft]
		}
	}

	if edge == EdgeEnd {
		if n.getLayout().direction() == DirectionRTL {
			return member[EdgeLeft]
		} else {
			return member[EdgeRight]
		}
	}

	return member[edge]
}

// LayoutBorder
func (n *Node) LayoutBorder(edge Edge) float32 {
	member := n.getLayout().border_
	if edge > EdgeEnd {
		panic("Cannot get layout properties of multi-edge shorthands")
	}
	if edge == EdgeStart {
		if n.getLayout().direction() == DirectionRTL {
			return member[EdgeRight]
		} else {
			return member[EdgeLeft]
		}
	}

	if edge == EdgeEnd {
		if n.getLayout().direction() == DirectionRTL {
			return member[EdgeLeft]
		} else {
			return member[EdgeRight]
		}
	}

	return member[edge]
}

// LayoutPadding
func (n *Node) LayoutPadding(edge Edge) float32 {
	member := n.getLayout().padding_
	if edge > EdgeEnd {
		panic("Cannot get layout properties of multi-edge shorthands")
	}
	if edge == EdgeStart {
		if n.getLayout().direction() == DirectionRTL {
			return member[EdgeRight]
		} else {
			return member[EdgeLeft]
		}
	}

	if edge == EdgeEnd {
		if n.getLayout().direction() == DirectionRTL {
			return member[EdgeLeft]
		} else {
			return member[EdgeRight]
		}
	}

	return member[edge]
}
