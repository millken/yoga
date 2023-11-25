package yoga

type CachedMeasurement struct {
	availableWidth    float32
	availableHeight   float32
	widthMeasureMode  YGMeasureMode
	heightMeasureMode YGMeasureMode
	computedWidth     float32
	computedHeight    float32
}
