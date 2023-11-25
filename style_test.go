package yoga

import "testing"

func TestA(t *testing.T) {
	t.Log("Direction", minimumBitCount(YGDirection(0)))
	t.Log("FlexDirection", minimumBitCount(YGFlexDirection(0)))
	t.Log("Justify", minimumBitCount(YGJustify(0)))
	t.Log("Align", minimumBitCount(YGAlign(0)))
	t.Log("PositionType", minimumBitCount(YGPositionType(0)))
	t.Log("Wrap", minimumBitCount(YGWrap(0)))
	t.Log("Overflow", minimumBitCount(YGOverflow(0)))

}
