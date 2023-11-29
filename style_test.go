package yoga

import "testing"

func TestA(t *testing.T) {
	t.Log("Direction", minimumBitCount(Direction(0)))
	t.Log("FlexDirection", minimumBitCount(FlexDirection(0)))
	t.Log("Justify", minimumBitCount(Justify(0)))
	t.Log("Align", minimumBitCount(Align(0)))
	t.Log("PositionType", minimumBitCount(PositionType(0)))
	t.Log("Wrap", minimumBitCount(Wrap(0)))
	t.Log("Overflow", minimumBitCount(Overflow(0)))

}
