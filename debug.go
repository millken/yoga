package yoga

import (
	"fmt"
	"strings"
)

type PrintOptions uint8

const (
	PrintOptionsLayout PrintOptions = 1 << iota
	PrintOptionsStyle
	PrintOptionsChildren
)

func (p PrintOptions) String() string {
	switch p {
	case PrintOptionsLayout:
		return "layout"
	case PrintOptionsStyle:
		return "style"
	case PrintOptionsChildren:
		return "children"
	}
	return "unknown"
}
func indent(base *strings.Builder, level uint32) {
	for i := uint32(0); i < level; i++ {
		base.WriteString("  ")
	}
}

func appendFloatOptionalIfDefined(base *strings.Builder, str string, value float32) {
	if !IsNaN(value) {
		base.WriteString(fmt.Sprintf("%s: %g; ", str, value))
	}
}

func appendNumberIfNotAuto(base *strings.Builder, str string, value Value) {
	if value.Unit != UnitAuto {
		appendNumberIfNotUndefined(base, str, value)
	}
}

func appendNumberIfNotZero(base *strings.Builder, str string, number Value) {
	if number.Unit == UnitAuto {
		base.WriteString(str)
		base.WriteString(": auto; ")
	} else if !number.IsUndefined() && number.Value != 0 {
		appendNumberIfNotUndefined(base, str, number)
	}
}

func appendNumberIfNotUndefined(base *strings.Builder, str string, number Value) {
	if number.Unit != UnitUndefined {
		if number.Unit == UnitAuto {
			base.WriteString(str)
			base.WriteString(": auto; ")
		} else {
			unit := If(number.Unit == UnitPoint, "px", "%")
			fmt.Fprintf(base, "%s: %g%s; ", str, number.Value, unit)
		}
	}
}
func nodeToString(str *strings.Builder, node *Node, options PrintOptions, level uint32) {
	if node == nil {
		return
	}
	indent(str, level)
	str.WriteString("<div ")
	if options&PrintOptionsLayout == PrintOptionsLayout {
		str.WriteString("layout=\"")
		fmt.Fprintf(str, "width: %g; ", node.GetComputedWidth())
		fmt.Fprintf(str, "height: %g; ", node.GetComputedHeight())
		fmt.Fprintf(str, "top: %g;", node.GetComputedTop())
		fmt.Fprintf(str, "left: %g; ", node.GetComputedLeft())
		str.WriteString("\" ")
	}

	if options&PrintOptionsStyle == PrintOptionsStyle {
		str.WriteString("style=\"")
		style := node
		oriStyle := NewNode()
		if style.GetFlexDirection() != oriStyle.GetFlexDirection() {
			fmt.Fprintf(str, "flex-direction: %s; ", style.GetFlexDirection().String())
		}
		if style.GetJustifyContent() != oriStyle.GetJustifyContent() {
			fmt.Fprintf(str, "justify-content: %s; ", style.GetJustifyContent().String())
		}
		if style.GetAlignItems() != oriStyle.GetAlignItems() {
			fmt.Fprintf(str, "align-items: %s; ", style.GetAlignItems().String())
		}
		if style.GetAlignContent() != oriStyle.GetAlignContent() {
			fmt.Fprintf(str, "align-content: %s; ", style.GetAlignContent().String())
		}
		if style.GetAlignSelf() != oriStyle.GetAlignSelf() {
			fmt.Fprintf(str, "align-self: %s; ", style.GetAlignSelf().String())
		}
		appendFloatOptionalIfDefined(str, "flex-grow", style.GetFlexGrow())
		appendFloatOptionalIfDefined(str, "flex-shrink", style.GetFlexShrink())
		appendNumberIfNotAuto(str, "flex-basis", style.GetFlexBasis())
		if style.GetFlexWrap() != oriStyle.GetFlexWrap() {
			fmt.Fprintf(str, "flex-wrap: %s; ", style.GetFlexWrap().String())
		}

		if style.GetOverflow() != oriStyle.GetOverflow() {
			fmt.Fprintf(str, "overflow: %s; ", style.GetOverflow().String())
		}

		if style.GetDisplay() != oriStyle.GetDisplay() {
			fmt.Fprintf(str, "display: %s; ", style.GetDisplay().String())
		}

		if style.GetPositionType() != oriStyle.GetPositionType() {
			fmt.Fprintf(str, "position: %s; ", style.GetPositionType().String())
		}
		if style.GetMargin(EdgeLeft).Value != 0 {
			fmt.Fprintf(str, "margin-left: %gpx; ", style.GetMargin(EdgeLeft).Value)
		}

	}
	str.WriteString("\">")

	childCount := node.GetChildCount()
	if options&PrintOptionsChildren == PrintOptionsChildren && childCount > 0 {
		str.WriteString("\n")
		for i := uint32(0); i < childCount; i++ {
			nodeToString(str, node.GetChild(i), options, level+1)
		}
		indent(str, level)
		str.WriteString("</div>")
	}
}
