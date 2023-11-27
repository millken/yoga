package yoga

import (
	"fmt"
	"strings"
)

var (
	gPrintChanges = true
	gPrintSkips   = true
)

func vlog(config *Config,
	node *Node,
	level YGLogLevel,
	format string,
	args ...interface{}) {
	if config != nil {
		config.log(node, level, format, args...)
	} else {
		DefaultLogger(config, node, level, format, args...)
	}
}

func indent(base *strings.Builder, level uint32) {
	for i := uint32(0); i < level; i++ {
		base.WriteString("  ")
	}
}

func appendFloatOptionalIfDefined(base *strings.Builder, str string, value FloatOptional) {
	if value.isDefined() {
		base.WriteString(fmt.Sprintf("%s: %g; ", str, value.unwrap()))
	}
}

func appendNumberIfNotAuto(base *strings.Builder, str string, value YGValue) {
	if value.unit != YGUnitAuto {
		appendNumberIfNotUndefined(base, str, value)
	}
}

func appendNumberIfNotZero(base *strings.Builder, str string, number YGValue) {
	if number.unit == YGUnitAuto {
		base.WriteString(str)
		base.WriteString(": auto; ")
	} else if !number.isUndefined() && number.value != 0 {
		appendNumberIfNotUndefined(base, str, number)
	}
}

func areFourValuesEqual(four [EdgeCount]CompactValue) bool {
	return inexactEquals(four[0], four[1]) &&
		inexactEquals(four[0], four[2]) &&
		inexactEquals(four[0], four[3])
}

func appendEdges(base *strings.Builder, key string, edges [EdgeCount]CompactValue) {
	if areFourValuesEqual(edges) {
		edgeValue := (&nodeDefaults).computeEdgeValueForColumn(edges, YGEdgeLeft)
		appendNumberIfNotUndefined(base, key, edgeValue.YGValue())
	} else {
		for edge := YGEdgeLeft; edge < EdgeCount; edge++ {
			appendNumberIfNotZero(base, fmt.Sprintf("%s-%s", key, edge.String()), edges[edge].YGValue())
		}
	}
}

func appendNumberIfNotUndefined(base *strings.Builder, str string, number YGValue) {
	if number.unit != YGUnitUndefined {
		if number.unit == YGUnitAuto {
			base.WriteString(str)
			base.WriteString(": auto; ")
		} else {
			unit := If(number.unit == YGUnitPoint, "px", "%")
			base.WriteString(fmt.Sprintf("%s: %g%s; ", str, number.value, unit))
		}
	}
}

func nodeToString(str *strings.Builder, node *Node, options YGPrintOptions, level uint32) {
	if node == nil {
		return
	}
	indent(str, level)
	str.WriteString("<div ")
	if options&YGPrintOptionsLayout == YGPrintOptionsLayout {
		str.WriteString("layout=\"")
		str.WriteString(fmt.Sprintf("width: %g; ", node.getLayout().dimension(YGDimensionWidth)))
		str.WriteString(fmt.Sprintf("height: %g; ", node.getLayout().dimension(YGDimensionHeight)))
		str.WriteString(fmt.Sprintf("top: %g;", node.getLayout().position[YGEdgeTop]))
		str.WriteString(fmt.Sprintf("left: %g; ", node.getLayout().position[YGEdgeLeft]))
		str.WriteString("\" ")
	}

	if options&YGPrintOptionsStyle == YGPrintOptionsStyle {
		str.WriteString("style=\"")
		style := node.getStyle()
		oriStyle := (&nodeDefaults).getStyle()
		if style.flexDirection() != oriStyle.flexDirection() {
			str.WriteString(fmt.Sprintf("flex-direction: %s; ", style.flexDirection().String()))
		}
		if style.justifyContent() != oriStyle.justifyContent() {
			str.WriteString(fmt.Sprintf("justify-content: %s; ", style.justifyContent().String()))
		}
		if style.alignItems() != oriStyle.alignItems() {
			str.WriteString(fmt.Sprintf("align-items: %s; ", style.alignItems().String()))
		}
		if style.alignContent() != oriStyle.alignContent() {
			str.WriteString(fmt.Sprintf("align-content: %s; ", style.alignContent().String()))
		}
		if style.alignSelf() != oriStyle.alignSelf() {
			str.WriteString(fmt.Sprintf("align-self: %s; ", style.alignSelf().String()))
		}
		appendFloatOptionalIfDefined(str, "flex-grow", style.flexGrow())
		appendFloatOptionalIfDefined(str, "flex-shrink", style.flexShrink())
		appendNumberIfNotAuto(str, "flex-basis", style.flexBasis().YGValue())
		appendFloatOptionalIfDefined(str, "flex", style.flex())

		if style.flexWrap() != oriStyle.flexWrap() {
			str.WriteString(fmt.Sprintf("flex-wrap: %s; ", style.flexWrap().String()))
		}

		if style.overflow() != oriStyle.overflow() {
			str.WriteString(fmt.Sprintf("overflow: %s; ", style.overflow().String()))
		}

		if style.display() != oriStyle.display() {
			str.WriteString(fmt.Sprintf("display: %s; ", style.display().String()))
		}

		appendEdges(str, "margin", style.margin_)
		appendEdges(str, "padding", style.padding_)
		appendEdges(str, "border", style.border_)

		if style.gap(YGGutterAll).isDefined() {
			appendNumberIfNotUndefined(str, "gap", style.gap(YGGutterAll).YGValue())
		} else {
			appendNumberIfNotUndefined(str, "column-gap", style.gap(YGGutterColumn).YGValue())
			appendNumberIfNotUndefined(str, "row-gap", style.gap(YGGutterRow).YGValue())
		}

		appendNumberIfNotAuto(str, "width", style.dimension(YGDimensionWidth).YGValue())
		appendNumberIfNotAuto(str, "height", style.dimension(YGDimensionHeight).YGValue())
		appendNumberIfNotAuto(str, "max-width", style.maxDimension(YGDimensionWidth).YGValue())
		appendNumberIfNotAuto(str, "max-height", style.maxDimension(YGDimensionHeight).YGValue())
		appendNumberIfNotAuto(str, "min-width", style.minDimension(YGDimensionWidth).YGValue())
		appendNumberIfNotAuto(str, "min-height", style.minDimension(YGDimensionHeight).YGValue())

		if style.positionType() != oriStyle.positionType() {
			str.WriteString(fmt.Sprintf("position: %s; ", style.positionType().String()))
		}

		appendNumberIfNotUndefined(str, "left", style.position(YGEdgeLeft).YGValue())
		appendNumberIfNotUndefined(str, "right", style.position(YGEdgeRight).YGValue())
		appendNumberIfNotUndefined(str, "top", style.position(YGEdgeTop).YGValue())
		appendNumberIfNotUndefined(str, "bottom", style.position(YGEdgeBottom).YGValue())

		if node.hasMeasureFunc() {
			str.WriteString(fmt.Sprintf("has-custom-measure-func: true; "))
		}
	}
	str.WriteString(fmt.Sprintf(">"))

	childCount := node.getChildCount()
	if options&YGPrintOptionsChildren == YGPrintOptionsChildren && childCount > 0 {
		str.WriteString("\n")
		for i := uint32(0); i < childCount; i++ {
			nodeToString(str, node.getChild(i), options, level+1)
		}
		indent(str, level)
		str.WriteString("</div>")
	}
}

func vprint(node *Node, printOptions YGPrintOptions) {
	var str strings.Builder
	str.Reset()
	nodeToString(&str, node, printOptions, 0)
	vlog(node.getConfig(), node, YGLogLevelDebug, str.String())
}
