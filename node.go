package yoga

/*
#include "cgo_wrapper.h"
*/
import "C"
import (
	"runtime"
)

// Layout represents the computed layout information
type Layout struct {
	Left   float32
	Right  float32
	Top    float32
	Bottom float32
	Width  float32
	Height float32
}

// Size represents the measurement result
type Size struct {
	Width  float32
	Height float32
}

// MeasureFunc defines the type for the measurement callback function
type MeasureFunc func(width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size

// DirtiedFunc defines the callback for node state changes
type DirtiedFunc func()

type Node struct {
	node    C.YGNodeRef
	context interface{}
}

// NewNode creates a default node
func NewNode() *Node {
	n := &Node{node: C.YGNodeNew()}
	runtime.SetFinalizer(n, (*Node).Destroy)
	return n
}

// NewNodeWithConfig creates a node with the specified configuration
func NewNodeWithConfig(config *Config) *Node {
	node := C.YGNodeNewWithConfig(config.ref())
	if node == nil {
		return nil
	}
	n := &Node{node: node}
	runtime.SetFinalizer(n, (*Node).Destroy)
	return n
}

// Clone creates a copy of the node with the same context and children, but no owner set
func (n *Node) Clone() *Node {
	if n.node == nil {
		return nil
	}
	clonedNode := C.YGNodeClone(n.node)
	if clonedNode == nil {
		return nil
	}
	newNode := &Node{
		node:    clonedNode,
		context: n.context, // 直接复制 context
	}
	runtime.SetFinalizer(newNode, (*Node).Destroy)
	return newNode
}

// SetConfig sets the configuration for the node
func (n *Node) SetConfig(config *Config) {
	if n.node != nil && config != nil {
		C.YGNodeSetConfig(n.node, config.ref())
	}
}

// GetConfig gets the configuration of the node
func (n *Node) GetConfig() *Config {
	if n.node != nil {
		configPtr := C.YGNodeGetConfig(n.node)
		if configPtr == nil {
			return nil
		}
		return wrapConfigRef(configPtr)
	}
	return nil
}

// SetContext sets the context for the node
func (n *Node) SetContext(context interface{}) {
	if n.node == nil {
		return
	}
	n.context = context
}

// GetContext gets the context of the node
func (n *Node) GetContext() interface{} {
	if n.node == nil {
		return nil
	}
	return n.context
}

// SetNodeType Sets whether a leaf node's layout results may be truncated during layout rounding.
func (n *Node) SetNodeType(nodeType NodeType) {
	if n.node != nil {
		C.YGNodeSetNodeType(n.node, C.YGNodeType(nodeType))
	}
}

// GetNodeType gets the type of the node
func (n *Node) GetNodeType() NodeType {
	if n.node != nil {
		return NodeType(C.YGNodeGetNodeType(n.node))
	}
	return NodeTypeDefault
}

// Destroy releases the resources of the node
func (n *Node) Destroy() {
	if n.node != nil {
		// 清理节点上下文（包含所有回调句柄）
		deleteNodeContext(n.node)
		C.YGNodeFree(n.node)
		n.node = nil
	}
}

// Finalize frees the node without disconnecting it from its owner or children
func (n *Node) Finalize() {
	if n.node != nil {
		// 清理节点上下文（包含所有回调句柄）
		deleteNodeContext(n.node)
		C.YGNodeFinalize(n.node)
		n.node = nil
	}
}

func (n *Node) ref() C.YGNodeRef {
	return n.node
}

// Reset resets the node settings
func (n *Node) Reset() {
	if n.node != nil {
		C.YGNodeReset(n.node)
	}
}

// CopyStyle copies the style from another node
func (n *Node) CopyStyle(other *Node) {
	if n.node != nil && other.node != nil {
		C.YGNodeCopyStyle(n.node, other.node)
	}
}

// SetPositionType sets the position type
func (n *Node) SetPositionType(positionType PositionType) {
	if n.node != nil {
		C.YGNodeStyleSetPositionType(n.node, C.YGPositionType(positionType))
	}
}

// SetPosition sets the position
func (n *Node) SetPosition(edge Edge, position float32) {
	if n.node != nil {
		C.YGNodeStyleSetPosition(n.node, C.YGEdge(edge), C.float(position))
	}
}

// SetPositionPercent sets the position in percentage
func (n *Node) SetPositionPercent(edge Edge, position float32) {
	if n.node != nil {
		C.YGNodeStyleSetPositionPercent(n.node, C.YGEdge(edge), C.float(position))
	}
}

// SetAlignContent sets the alignment of content
func (n *Node) SetAlignContent(alignContent Align) {
	if n.node != nil {
		C.YGNodeStyleSetAlignContent(n.node, C.YGAlign(alignContent))
	}
}

// SetAlignItems sets the alignment of child items
func (n *Node) SetAlignItems(alignItems Align) {
	if n.node != nil {
		C.YGNodeStyleSetAlignItems(n.node, C.YGAlign(alignItems))
	}
}

// SetFlexDirection sets the flex direction
func (n *Node) SetFlexDirection(flexDirection FlexDirection) {
	if n.node != nil {
		C.YGNodeStyleSetFlexDirection(n.node, C.YGFlexDirection(flexDirection))
	}
}

// SetJustifyContent sets the alignment along the main axis
func (n *Node) SetJustifyContent(justifyContent Justify) {
	if n.node != nil {
		C.YGNodeStyleSetJustifyContent(n.node, C.YGJustify(justifyContent))
	}
}

// SetMargin sets the margin
func (n *Node) SetMargin(edge Edge, margin float32) {
	if n.node != nil {
		C.YGNodeStyleSetMargin(n.node, C.YGEdge(edge), C.float(margin))
	}
}

// SetWidth sets the width
func (n *Node) SetWidth(width float32) {
	if n.node != nil {
		C.YGNodeStyleSetWidth(n.node, C.float(width))
	}
}

// SetHeight sets the height
func (n *Node) SetHeight(height float32) {
	if n.node != nil {
		C.YGNodeStyleSetHeight(n.node, C.float(height))
	}
}

// SetPadding sets the padding
func (n *Node) SetPadding(edge Edge, padding float32) {
	if n.node != nil {
		C.YGNodeStyleSetPadding(n.node, C.YGEdge(edge), C.float(padding))
	}
}

// SetPaddingPercent sets the padding in percentage
func (n *Node) SetPaddingPercent(edge Edge, padding float32) {
	if n.node != nil {
		C.YGNodeStyleSetPaddingPercent(n.node, C.YGEdge(edge), C.float(padding))
	}
}

// InsertChild inserts a child node at the specified position
func (n *Node) InsertChild(child *Node, index uint32) {
	if n.node != nil && child.node != nil {
		C.YGNodeInsertChild(n.node, child.node, C.size_t(index))
	}
}

// RemoveChild removes a child node
func (n *Node) RemoveChild(child *Node) {
	if n.node != nil && child.node != nil {
		C.YGNodeRemoveChild(n.node, child.node)
	}
}

// GetChildCount gets the number of child nodes
func (n *Node) GetChildCount() uint32 {
	if n.node != nil {
		return uint32(C.YGNodeGetChildCount(n.node))
	}
	return 0
}

// GetParent gets the parent node
func (n *Node) GetParent() *Node {
	if n.node != nil {
		parentPtr := C.YGNodeGetParent(n.node)
		if parentPtr == nil {
			return nil
		}
		return wrapNodeRef(parentPtr)
	}
	return nil
}

// GetChild gets the child node at the specified index
func (n *Node) GetChild(index uint32) *Node {
	if n.node != nil {
		childPtr := C.YGNodeGetChild(n.node, C.size_t(index))
		if childPtr == nil {
			return nil
		}
		return wrapNodeRef(childPtr)
	}
	return nil
}

// CalculateLayout calculates the layout
func (n *Node) CalculateLayout(width, height float32, direction Direction) {
	if n.node != nil {
		C.YGNodeCalculateLayout(n.node, C.float(width), C.float(height), C.YGDirection(direction))
	}
}

// GetComputedLeft gets the computed left position
func (n *Node) GetComputedLeft() float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetLeft(n.node))
	}
	return 0
}

// GetComputedTop gets the computed top position
func (n *Node) GetComputedTop() float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetTop(n.node))
	}
	return 0
}

// GetComputedWidth gets the computed width
func (n *Node) GetComputedWidth() float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetWidth(n.node))
	}
	return 0
}

// GetComputedHeight gets the computed height
func (n *Node) GetComputedHeight() float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetHeight(n.node))
	}
	return 0
}

// GetComputedLayout gets the complete computed layout
func (n *Node) GetComputedLayout() Layout {
	if n.node != nil {
		return Layout{
			Left:   float32(C.YGNodeLayoutGetLeft(n.node)),
			Right:  float32(C.YGNodeLayoutGetRight(n.node)),
			Top:    float32(C.YGNodeLayoutGetTop(n.node)),
			Bottom: float32(C.YGNodeLayoutGetBottom(n.node)),
			Width:  float32(C.YGNodeLayoutGetWidth(n.node)),
			Height: float32(C.YGNodeLayoutGetHeight(n.node)),
		}
	}
	return Layout{}
}

// GetHadOverflow checks if the node had overflow
func (n *Node) HadOverflow() bool {
	if n.node != nil {
		return bool(C.YGNodeLayoutGetHadOverflow(n.node))
	}
	return false
}

// GetLayoutDirection gets the layout direction after layout calculation
func (n *Node) GetLayoutDirection() Direction {
	if n.node != nil {
		return Direction(C.YGNodeLayoutGetDirection(n.node))
	}
	return DirectionLTR
}

// GetRawWidth gets the measured width before layout rounding
func (n *Node) GetRawWidth() float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetRawWidth(n.node))
	}
	return 0
}

// GetRawHeight gets the measured height before layout rounding
func (n *Node) GetRawHeight() float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetRawHeight(n.node))
	}
	return 0
}

// BaselineFunc defines the type for the baseline callback function
type BaselineFunc func(width float32, height float32) float32

// SetMeasureFunc sets the measurement function
func (n *Node) SetMeasureFunc(measureFunc MeasureFunc) {
	if n.node == nil {
		return
	}
	if measureFunc == nil {
		// 取消回调并清理句柄
		C.YGNodeSetMeasureFunc(n.node, nil)
		deleteMeasureHandle(n.node)
		return
	}
	// 存储/替换该节点的 MeasureFunc 句柄，并设置 C 层回调
	setMeasureHandle(n.node, measureFunc)
	C.YGNodeSetMeasureFunc(n.node, (C.YGMeasureFunc)(C.goMeasureInvoke))

	// 标记节点为 dirty，强制重新计算布局
	C.YGNodeMarkDirty(n.node)
}

// GetMeasureFunc gets the current measurement callback function
func (n *Node) GetMeasureFunc() MeasureFunc {
	if n.node == nil {
		return nil
	}
	// 从 NodeContext 中获取 measure handle
	handle := getMeasureHandleByNode(n.node)
	if handle == 0 {
		return nil
	}
	if measureFunc, ok := handle.Value().(MeasureFunc); ok {
		return measureFunc
	}
	return nil
}

// SetBaselineFunc sets the baseline function
func (n *Node) SetBaselineFunc(baselineFunc BaselineFunc) {
	// This method requires additional cgo callback implementation
	// Similar to SetMeasureFunc
}

// HasBaselineFunc checks if a baseline function is set
func (n *Node) HasBaselineFunc() bool {
	if n.node != nil {
		return bool(C.YGNodeHasBaselineFunc(n.node))
	}
	return false
}

// MarkDirty marks the node as dirty (needs recalculation)
func (n *Node) MarkDirty() {
	if n.node != nil {
		C.YGNodeMarkDirty(n.node)
	}
}

// IsDirty checks if the node is dirty
func (n *Node) IsDirty() bool {
	if n.node != nil {
		return bool(C.YGNodeIsDirty(n.node))
	}
	return false
}

// HasNewLayout checks if the node has a new layout
func (n *Node) HasNewLayout() bool {
	if n.node != nil {
		return bool(C.YGNodeGetHasNewLayout(n.node))
	}
	return false
}

// SetHasNewLayout marks the layout as seen
func (n *Node) SetHasNewLayout(hasNewLayout bool) {
	if n.node != nil {
		C.YGNodeSetHasNewLayout(n.node, C.bool(hasNewLayout))
	}
}

// GetComputedMargin gets the computed margin
func (n *Node) GetComputedMargin(edge Edge) float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetMargin(n.node, C.YGEdge(edge)))
	}
	return 0
}

// GetComputedBorder gets the computed border width
func (n *Node) GetComputedBorder(edge Edge) float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetBorder(n.node, C.YGEdge(edge)))
	}
	return 0
}

// GetComputedPadding gets the computed padding
func (n *Node) GetComputedPadding(edge Edge) float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetPadding(n.node, C.YGEdge(edge)))
	}
	return 0
}

// GetComputedRight gets the computed right position
func (n *Node) GetComputedRight() float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetRight(n.node))
	}
	return 0
}

// GetComputedBottom gets the computed bottom position
func (n *Node) GetComputedBottom() float32 {
	if n.node != nil {
		return float32(C.YGNodeLayoutGetBottom(n.node))
	}
	return 0
}

// SetBoxSizing sets the box model type
func (n *Node) SetBoxSizing(boxSizing BoxSizing) {
	if n.node != nil {
		C.YGNodeStyleSetBoxSizing(n.node, C.YGBoxSizing(boxSizing))
	}
}

// GetBoxSizing gets the box model type
func (n *Node) GetBoxSizing() BoxSizing {
	if n.node != nil {
		return BoxSizing(C.YGNodeStyleGetBoxSizing(n.node))
	}
	return BoxSizingContentBox
}

// GetPositionType gets the position type
func (n *Node) GetPositionType() PositionType {
	if n.node != nil {
		return PositionType(C.YGNodeStyleGetPositionType(n.node))
	}
	return PositionTypeStatic
}

// SetPositionAuto sets the position to auto
func (n *Node) SetPositionAuto(edge Edge) {
	if n.node != nil {
		C.YGNodeStyleSetPositionAuto(n.node, C.YGEdge(edge))
	}
}

// GetPosition gets the position value
func (n *Node) GetPosition(edge Edge) Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetPosition(n.node, C.YGEdge(edge)))
	}
	return Value{}
}

// SetAlignSelf sets the alignment for the node itself
func (n *Node) SetAlignSelf(alignSelf Align) {
	if n.node != nil {
		C.YGNodeStyleSetAlignSelf(n.node, C.YGAlign(alignSelf))
	}
}

// GetAlignSelf gets the alignment for the node itself
func (n *Node) GetAlignSelf() Align {
	if n.node != nil {
		return Align(C.YGNodeStyleGetAlignSelf(n.node))
	}
	return AlignAuto
}

// GetAlignContent gets the alignment of content
func (n *Node) GetAlignContent() Align {
	if n.node != nil {
		return Align(C.YGNodeStyleGetAlignContent(n.node))
	}
	return AlignFlexStart
}

// GetAlignItems gets the alignment of child items
func (n *Node) GetAlignItems() Align {
	if n.node != nil {
		return Align(C.YGNodeStyleGetAlignItems(n.node))
	}
	return AlignStretch
}

// SetDirection sets the document flow direction
func (n *Node) SetDirection(direction Direction) {
	if n.node != nil {
		C.YGNodeStyleSetDirection(n.node, C.YGDirection(direction))
	}
}

// GetDirection gets the document flow direction
func (n *Node) GetDirection() Direction {
	if n.node != nil {
		return Direction(C.YGNodeStyleGetDirection(n.node))
	}
	return DirectionInherit
}

// SetFlexWrap sets the wrapping behavior
func (n *Node) SetFlexWrap(flexWrap Wrap) {
	if n.node != nil {
		C.YGNodeStyleSetFlexWrap(n.node, C.YGWrap(flexWrap))
	}
}

// GetFlexWrap gets the wrapping behavior
func (n *Node) GetFlexWrap() Wrap {
	if n.node != nil {
		return Wrap(C.YGNodeStyleGetFlexWrap(n.node))
	}
	return WrapNoWrap
}

// GetFlexDirection gets the flex direction
func (n *Node) GetFlexDirection() FlexDirection {
	if n.node != nil {
		return FlexDirection(C.YGNodeStyleGetFlexDirection(n.node))
	}
	return FlexDirectionColumn
}

// GetJustifyContent gets the alignment along the main axis
func (n *Node) GetJustifyContent() Justify {
	if n.node != nil {
		return Justify(C.YGNodeStyleGetJustifyContent(n.node))
	}
	return JustifyFlexStart
}

// SetMarginAuto sets the margin to auto
func (n *Node) SetMarginAuto(edge Edge) {
	if n.node != nil {
		C.YGNodeStyleSetMarginAuto(n.node, C.YGEdge(edge))
	}
}

// SetMarginPercent sets the margin in percentage
func (n *Node) SetMarginPercent(edge Edge, margin float32) {
	if n.node != nil {
		C.YGNodeStyleSetMarginPercent(n.node, C.YGEdge(edge), C.float(margin))
	}
}

// GetMargin gets the margin
func (n *Node) GetMargin(edge Edge) Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetMargin(n.node, C.YGEdge(edge)))
	}
	return Value{}
}

// SetOverflow sets the overflow behavior
func (n *Node) SetOverflow(overflow Overflow) {
	if n.node != nil {
		C.YGNodeStyleSetOverflow(n.node, C.YGOverflow(overflow))
	}
}

// GetOverflow gets the overflow behavior
func (n *Node) GetOverflow() Overflow {
	if n.node != nil {
		return Overflow(C.YGNodeStyleGetOverflow(n.node))
	}
	return OverflowVisible
}

// SetDisplay sets the display behavior
func (n *Node) SetDisplay(display Display) {
	if n.node != nil {
		C.YGNodeStyleSetDisplay(n.node, C.YGDisplay(display))
	}
}

// GetDisplay gets the display behavior
func (n *Node) GetDisplay() Display {
	if n.node != nil {
		return Display(C.YGNodeStyleGetDisplay(n.node))
	}
	return DisplayFlex
}

// SetFlex sets the flex layout
func (n *Node) SetFlex(flex float32) {
	if n.node != nil {
		C.YGNodeStyleSetFlex(n.node, C.float(flex))
	}
}

// GetFlex gets the flex layout value
func (n *Node) GetFlex() float32 {
	if n.node != nil {
		return float32(C.YGNodeStyleGetFlex(n.node))
	}
	return 0
}

// SetFlexGrow sets the flex grow factor
func (n *Node) SetFlexGrow(flexGrow float32) {
	if n.node != nil {
		C.YGNodeStyleSetFlexGrow(n.node, C.float(flexGrow))
	}
}

// GetFlexGrow gets the flex grow factor
func (n *Node) GetFlexGrow() float32 {
	if n.node != nil {
		return float32(C.YGNodeStyleGetFlexGrow(n.node))
	}
	return 0
}

// SetFlexShrink sets the flex shrink factor
func (n *Node) SetFlexShrink(flexShrink float32) {
	if n.node != nil {
		C.YGNodeStyleSetFlexShrink(n.node, C.float(flexShrink))
	}
}

// GetFlexShrink gets the flex shrink factor
func (n *Node) GetFlexShrink() float32 {
	if n.node != nil {
		return float32(C.YGNodeStyleGetFlexShrink(n.node))
	}
	return 0
}

// SetFlexBasis sets the flex basis size
func (n *Node) SetFlexBasis(flexBasis float32) {
	if n.node != nil {
		C.YGNodeStyleSetFlexBasis(n.node, C.float(flexBasis))
	}
}

// SetFlexBasisPercent sets the flex basis size in percentage
func (n *Node) SetFlexBasisPercent(flexBasis float32) {
	if n.node != nil {
		C.YGNodeStyleSetFlexBasisPercent(n.node, C.float(flexBasis))
	}
}

// SetFlexBasisAuto sets the flex basis size to auto
func (n *Node) SetFlexBasisAuto() {
	if n.node != nil {
		C.YGNodeStyleSetFlexBasisAuto(n.node)
	}
}

// GetFlexBasis gets the flex basis size
func (n *Node) GetFlexBasis() Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetFlexBasis(n.node))
	}
	return Value{}
}

// SetWidthPercent sets the width in percentage
func (n *Node) SetWidthPercent(width float32) {
	if n.node != nil {
		C.YGNodeStyleSetWidthPercent(n.node, C.float(width))
	}
}

// SetWidthAuto sets the width to auto
func (n *Node) SetWidthAuto() {
	if n.node != nil {
		C.YGNodeStyleSetWidthAuto(n.node)
	}
}

// GetWidth gets the width
func (n *Node) GetWidth() Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetWidth(n.node))
	}
	return Value{}
}

// SetHeightPercent sets the height in percentage
func (n *Node) SetHeightPercent(height float32) {
	if n.node != nil {
		C.YGNodeStyleSetHeightPercent(n.node, C.float(height))
	}
}

// SetHeightAuto sets the height to auto
func (n *Node) SetHeightAuto() {
	if n.node != nil {
		C.YGNodeStyleSetHeightAuto(n.node)
	}
}

// GetHeight gets the height
func (n *Node) GetHeight() Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetHeight(n.node))
	}
	return Value{}
}

// SetMinWidth sets the minimum width
func (n *Node) SetMinWidth(minWidth float32) {
	if n.node != nil {
		C.YGNodeStyleSetMinWidth(n.node, C.float(minWidth))
	}
}

// SetMinWidthPercent sets the minimum width in percentage
func (n *Node) SetMinWidthPercent(minWidth float32) {
	if n.node != nil {
		C.YGNodeStyleSetMinWidthPercent(n.node, C.float(minWidth))
	}
}

// GetMinWidth gets the minimum width
func (n *Node) GetMinWidth() Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetMinWidth(n.node))
	}
	return Value{}
}

// SetMinHeight sets the minimum height
func (n *Node) SetMinHeight(minHeight float32) {
	if n.node != nil {
		C.YGNodeStyleSetMinHeight(n.node, C.float(minHeight))
	}
}

// SetMinHeightPercent sets the minimum height in percentage
func (n *Node) SetMinHeightPercent(minHeight float32) {
	if n.node != nil {
		C.YGNodeStyleSetMinHeightPercent(n.node, C.float(minHeight))
	}
}

// GetMinHeight gets the minimum height
func (n *Node) GetMinHeight() Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetMinHeight(n.node))
	}
	return Value{}
}

// SetMaxWidth sets the maximum width
func (n *Node) SetMaxWidth(maxWidth float32) {
	if n.node != nil {
		C.YGNodeStyleSetMaxWidth(n.node, C.float(maxWidth))
	}
}

// SetMaxWidthPercent sets the maximum width in percentage
func (n *Node) SetMaxWidthPercent(maxWidth float32) {
	if n.node != nil {
		C.YGNodeStyleSetMaxWidthPercent(n.node, C.float(maxWidth))
	}
}

// GetMaxWidth gets the maximum width
func (n *Node) GetMaxWidth() Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetMaxWidth(n.node))
	}
	return Value{}
}

// SetMaxHeight sets the maximum height
func (n *Node) SetMaxHeight(maxHeight float32) {
	if n.node != nil {
		C.YGNodeStyleSetMaxHeight(n.node, C.float(maxHeight))
	}
}

// SetMaxHeightPercent sets the maximum height in percentage
func (n *Node) SetMaxHeightPercent(maxHeight float32) {
	if n.node != nil {
		C.YGNodeStyleSetMaxHeightPercent(n.node, C.float(maxHeight))
	}
}

// GetMaxHeight gets the maximum height
func (n *Node) GetMaxHeight() Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetMaxHeight(n.node))
	}
	return Value{}
}

// SetAspectRatio sets the aspect ratio
func (n *Node) SetAspectRatio(aspectRatio float32) {
	if n.node != nil {
		C.YGNodeStyleSetAspectRatio(n.node, C.float(aspectRatio))
	}
}

// GetAspectRatio gets the aspect ratio
func (n *Node) GetAspectRatio() float32 {
	if n.node != nil {
		return float32(C.YGNodeStyleGetAspectRatio(n.node))
	}
	return 0
}

// SetBorder sets the border width
func (n *Node) SetBorder(edge Edge, border float32) {
	if n.node != nil {
		C.YGNodeStyleSetBorder(n.node, C.YGEdge(edge), C.float(border))
	}
}

// GetBorder gets the border width
func (n *Node) GetBorder(edge Edge) float32 {
	if n.node != nil {
		return float32(C.YGNodeStyleGetBorder(n.node, C.YGEdge(edge)))
	}
	return 0
}

// GetPadding gets the padding
func (n *Node) GetPadding(edge Edge) Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetPadding(n.node, C.YGEdge(edge)))
	}
	return Value{}
}

// SetGap sets the gap
func (n *Node) SetGap(gutter Gutter, gapLength float32) {
	if n.node != nil {
		C.YGNodeStyleSetGap(n.node, C.YGGutter(gutter), C.float(gapLength))
	}
}

// SetGapPercent sets the gap in percentage
func (n *Node) SetGapPercent(gutter Gutter, gapLength float32) {
	if n.node != nil {
		C.YGNodeStyleSetGapPercent(n.node, C.YGGutter(gutter), C.float(gapLength))
	}
}

// GetGap gets the gap
func (n *Node) GetGap(gutter Gutter) Value {
	if n.node != nil {
		return valueFromYGValue(C.YGNodeStyleGetGap(n.node, C.YGGutter(gutter)))
	}
	return Value{}
}

// SetIsReferenceBaseline sets whether the node is a reference baseline
func (n *Node) SetIsReferenceBaseline(isReferenceBaseline bool) {
	if n.node != nil {
		C.YGNodeSetIsReferenceBaseline(n.node, C.bool(isReferenceBaseline))
	}
}

// IsReferenceBaseline checks if the node is a reference baseline
func (n *Node) IsReferenceBaseline() bool {
	if n.node != nil {
		return bool(C.YGNodeIsReferenceBaseline(n.node))
	}
	return false
}

// UnsetMeasureFunc unsets the measurement function
func (n *Node) UnsetMeasureFunc() {
	if n.node != nil {
		C.YGNodeSetMeasureFunc(n.node, nil)
		// 清理 measure handle，但保留 NodeContext 结构
		deleteMeasureHandle(n.node)
	}
}

// SetDirtiedFunc sets the dirtied callback function
func (n *Node) SetDirtiedFunc(dirtiedFunc DirtiedFunc) {
	// This method requires cgo callback implementation, see cgo_callbacks.go for details
}

// GetDirtiedFunc gets the current dirtied callback function
func (n *Node) GetDirtiedFunc() DirtiedFunc {
	// Note: This is a placeholder implementation
	// In practice, retrieving a cgo function pointer is complex and may not be necessary
	// This would require storing the callback reference for retrieval
	return nil
}

// UnsetDirtiedFunc unsets the dirtied callback function
func (n *Node) UnsetDirtiedFunc() {
	if n.node != nil {
		C.YGNodeSetDirtiedFunc(n.node, nil)
	}
}

// SetAlwaysFormsContainingBlock sets whether the node always forms a containing block
func (n *Node) SetAlwaysFormsContainingBlock(alwaysFormsContainingBlock bool) {
	if n.node != nil {
		C.YGNodeSetAlwaysFormsContainingBlock(n.node, C.bool(alwaysFormsContainingBlock))
	}
}

// GetAlwaysFormsContainingBlock checks if the node always forms a containing block
func (n *Node) GetAlwaysFormsContainingBlock() bool {
	if n.node != nil {
		return bool(C.YGNodeGetAlwaysFormsContainingBlock(n.node))
	}
	return false
}

// SetFlexBasisMaxContent sets the flex basis to max-content
func (n *Node) SetFlexBasisMaxContent() {
	if n.node != nil {
		C.YGNodeStyleSetFlexBasisMaxContent(n.node)
	}
}

// SetFlexBasisFitContent sets the flex basis to fit-content
func (n *Node) SetFlexBasisFitContent() {
	if n.node != nil {
		C.YGNodeStyleSetFlexBasisFitContent(n.node)
	}
}

// SetFlexBasisStretch sets the flex basis to stretch
func (n *Node) SetFlexBasisStretch() {
	if n.node != nil {
		C.YGNodeStyleSetFlexBasisStretch(n.node)
	}
}

// SetWidthMaxContent sets the width to max-content
func (n *Node) SetWidthMaxContent() {
	if n.node != nil {
		C.YGNodeStyleSetWidthMaxContent(n.node)
	}
}

// SetWidthFitContent sets the width to fit-content
func (n *Node) SetWidthFitContent() {
	if n.node != nil {
		C.YGNodeStyleSetWidthFitContent(n.node)
	}
}

// SetWidthStretch sets the width to stretch
func (n *Node) SetWidthStretch() {
	if n.node != nil {
		C.YGNodeStyleSetWidthStretch(n.node)
	}
}

// SetHeightMaxContent sets the height to max-content
func (n *Node) SetHeightMaxContent() {
	if n.node != nil {
		C.YGNodeStyleSetHeightMaxContent(n.node)
	}
}

// SetHeightFitContent sets the height to fit-content
func (n *Node) SetHeightFitContent() {
	if n.node != nil {
		C.YGNodeStyleSetHeightFitContent(n.node)
	}
}

// SetHeightStretch sets the height to stretch
func (n *Node) SetHeightStretch() {
	if n.node != nil {
		C.YGNodeStyleSetHeightStretch(n.node)
	}
}

// SetMinWidthMaxContent sets the minimum width to max-content
func (n *Node) SetMinWidthMaxContent() {
	if n.node != nil {
		C.YGNodeStyleSetMinWidthMaxContent(n.node)
	}
}

// SetMinWidthFitContent sets the minimum width to fit-content
func (n *Node) SetMinWidthFitContent() {
	if n.node != nil {
		C.YGNodeStyleSetMinWidthFitContent(n.node)
	}
}

// SetMinWidthStretch sets the minimum width to stretch
func (n *Node) SetMinWidthStretch() {
	if n.node != nil {
		C.YGNodeStyleSetMinWidthStretch(n.node)
	}
}

// SetMinHeightMaxContent sets the minimum height to max-content
func (n *Node) SetMinHeightMaxContent() {
	if n.node != nil {
		C.YGNodeStyleSetMinHeightMaxContent(n.node)
	}
}

// SetMinHeightFitContent sets the minimum height to fit-content
func (n *Node) SetMinHeightFitContent() {
	if n.node != nil {
		C.YGNodeStyleSetMinHeightFitContent(n.node)
	}
}

// SetMinHeightStretch sets the minimum height to stretch
func (n *Node) SetMinHeightStretch() {
	if n.node != nil {
		C.YGNodeStyleSetMinHeightStretch(n.node)
	}
}

// SetMaxWidthMaxContent sets the maximum width to max-content
func (n *Node) SetMaxWidthMaxContent() {
	if n.node != nil {
		C.YGNodeStyleSetMaxWidthMaxContent(n.node)
	}
}

// SetMaxWidthFitContent sets the maximum width to fit-content
func (n *Node) SetMaxWidthFitContent() {
	if n.node != nil {
		C.YGNodeStyleSetMaxWidthFitContent(n.node)
	}
}

// SetMaxWidthStretch sets the maximum width to stretch
func (n *Node) SetMaxWidthStretch() {
	if n.node != nil {
		C.YGNodeStyleSetMaxWidthStretch(n.node)
	}
}

// SetMaxHeightMaxContent sets the maximum height to max-content
func (n *Node) SetMaxHeightMaxContent() {
	if n.node != nil {
		C.YGNodeStyleSetMaxHeightMaxContent(n.node)
	}
}

// SetMaxHeightFitContent sets the maximum height to fit-content
func (n *Node) SetMaxHeightFitContent() {
	if n.node != nil {
		C.YGNodeStyleSetMaxHeightFitContent(n.node)
	}
}

// SetMaxHeightStretch sets the maximum height to stretch
func (n *Node) SetMaxHeightStretch() {
	if n.node != nil {
		C.YGNodeStyleSetMaxHeightStretch(n.node)
	}
}

// SwapChild replaces the child node at the specified index with a new one
func (n *Node) SwapChild(child *Node, index uint32) {
	if n.node != nil && child.node != nil {
		C.YGNodeSwapChild(n.node, child.node, C.size_t(index))
	}
}

// FreeRecursive frees the node and all its children recursively
func (n *Node) FreeRecursive() {
	if n.node != nil {
		C.YGNodeFreeRecursive(n.node)
	}
}

// RemoveAllChildren removes all child nodes from the node
func (n *Node) RemoveAllChildren() {
	if n.node != nil {
		C.YGNodeRemoveAllChildren(n.node)
	}
}

// SetChildren sets children according to the given list of nodes.
func (n *Node) SetChildren(children *Node, count uint32) {
	if n.node != nil && children != nil {
		C.YGNodeSetChildren(n.node, &children.node, C.size_t(count))
	}
}

// GetOwner returns the owner of the node.
func (n *Node) GetOwner() *Node {
	if n.node != nil {
		ownerPtr := C.YGNodeGetOwner(n.node)
		if ownerPtr == nil {
			return nil
		}
		return wrapNodeRef(ownerPtr)
	}
	return nil
}

// Free Frees the Yoga node without disconnecting it from its owner or children. Allows garbage collecting Yoga nodes in parallel when the entire tree is unreachable.
func (n *Node) Free() {
	if n.node != nil {
		// 清理节点上下文（包含所有回调句柄）
		deleteNodeContext(n.node)
		C.YGNodeFree(n.node)
		n.node = nil
	}
}
