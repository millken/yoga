package yoga

func sizeIsExactAndMatchesOldMeasuredSize(sizeMode YGMeasureMode, size, lastComputedSize float32) bool {
	return sizeMode == YGMeasureModeExactly && inexactEqual(size, lastComputedSize)
}

func oldSizeIsUnspecifiedAndStillFits(sizeMode YGMeasureMode, size float32, lastSizeMode YGMeasureMode, lastComputedSize float32) bool {
	return sizeMode == YGMeasureModeAtMost && lastSizeMode == YGMeasureModeUndefined &&
		(size >= lastComputedSize || inexactEqual(size, lastComputedSize))
}

func newMeasureSizeIsStricterAndStillValid(sizeMode YGMeasureMode, size float32, lastSizeMode YGMeasureMode, lastSize, lastComputedSize float32) bool {
	return lastSizeMode == YGMeasureModeAtMost && sizeMode == YGMeasureModeAtMost && !IsNaN(lastSize) &&
		!IsNaN(size) && !IsNaN(lastComputedSize) && lastSize > size &&
		(lastComputedSize <= size || inexactEqual(size, lastComputedSize))
}

func canUseCachedMeasurement(widthMode YGMeasureMode, availableWidth float32, heightMode YGMeasureMode, availableHeight float32,
	lastWidthMode YGMeasureMode, lastAvailableWidth float32, lastHeightMode YGMeasureMode, lastAvailableHeight float32,
	lastComputedWidth, lastComputedHeight, marginRow, marginColumn float32, config *YGConfig) bool {
	if (!IsNaN(lastComputedHeight) && lastComputedHeight < 0) ||
		(!IsNaN(lastComputedWidth) && lastComputedWidth < 0) {
		return false
	}

	pointScaleFactor := config.getPointScaleFactor()

	useRoundedComparison := config != nil && pointScaleFactor != 0
	effectiveWidth := availableWidth
	effectiveHeight := availableHeight
	effectiveLastWidth := lastAvailableWidth
	effectiveLastHeight := lastAvailableHeight

	if useRoundedComparison {
		effectiveWidth = roundValueToPixelGrid(float64(availableWidth), float64(pointScaleFactor), false, false)
		effectiveHeight = roundValueToPixelGrid(float64(availableHeight), float64(pointScaleFactor), false, false)
		effectiveLastWidth = roundValueToPixelGrid(float64(lastAvailableWidth), float64(pointScaleFactor), false, false)
		effectiveLastHeight = roundValueToPixelGrid(float64(lastAvailableHeight), float64(pointScaleFactor), false, false)
	}

	hasSameWidthSpec := lastWidthMode == widthMode && inexactEqual(effectiveLastWidth, effectiveWidth)
	hasSameHeightSpec := lastHeightMode == heightMode && inexactEqual(effectiveLastHeight, effectiveHeight)

	widthIsCompatible :=
		hasSameWidthSpec ||
			sizeIsExactAndMatchesOldMeasuredSize(widthMode, availableWidth-marginRow, lastComputedWidth) ||
			oldSizeIsUnspecifiedAndStillFits(
				widthMode,
				availableWidth-marginRow,
				lastWidthMode,
				lastComputedWidth) ||
			newMeasureSizeIsStricterAndStillValid(
				widthMode,
				availableWidth-marginRow,
				lastWidthMode,
				lastAvailableWidth,
				lastComputedWidth)

	heightIsCompatible :=
		hasSameHeightSpec ||
			sizeIsExactAndMatchesOldMeasuredSize(
				heightMode, availableHeight-marginColumn, lastComputedHeight) ||
			oldSizeIsUnspecifiedAndStillFits(
				heightMode,
				availableHeight-marginColumn,
				lastHeightMode,
				lastComputedHeight) ||
			newMeasureSizeIsStricterAndStillValid(
				heightMode,
				availableHeight-marginColumn,
				lastHeightMode,
				lastAvailableHeight,
				lastComputedHeight)

	return widthIsCompatible && heightIsCompatible
}
