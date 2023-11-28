package yoga

func resolveChildAlignment(node *Node, child *Node) YGAlign {
	align := If(child.getStyle().alignSelf() == YGAlignAuto, node.getStyle().alignItems(), child.getStyle().alignSelf())
	if align == YGAlignBaseline && isColumn(node.getStyle().flexDirection()) {
		return YGAlignFlexStart
	}
	return align
}
