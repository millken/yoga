package yoga

func resolveChildAlignment(node *YGNode, child *YGNode) YGAlign {
	align := child.getStyle().alignSelf()
	if align == YGAlignAuto {
		align = node.getStyle().alignItems()
	} else {
		align = node.getStyle().alignSelf()
	}
	if align == YGAlignBaseline && isColumn(node.getStyle().flexDirection()) {
		return YGAlignFlexStart
	}
	return align
}
