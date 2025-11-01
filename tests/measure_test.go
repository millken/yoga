/*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package yoga_test

import (
	"math"
	"strings"

	"github.com/millken/yoga"
)

// intrinsicSizeMeasureFunc is a text measurement function for intrinsic size tests
// This implements the same algorithm as the C++ version in Yoga's test utilities
func intrinsicSizeMeasureFunc(text string, flexDirection yoga.FlexDirection) yoga.MeasureFunc {
	return func(width float32, widthMode yoga.MeasureMode, height float32, heightMode yoga.MeasureMode) yoga.Size {
		const widthPerChar float32 = 10.0
		const heightPerChar float32 = 10.0
		var measuredWidth float32
		var measuredHeight float32

		// Calculate width
		if widthMode == yoga.MeasureModeExactly {
			measuredWidth = width
		} else if widthMode == yoga.MeasureModeAtMost {
			measuredWidth = float32(math.Min(float64(len(text))*float64(widthPerChar), float64(width)))
		} else {
			measuredWidth = float32(len(text)) * widthPerChar
		}

		// Calculate effective width for text wrapping
		var effectiveWidth float32
		if flexDirection == yoga.FlexDirectionColumn {
			effectiveWidth = measuredWidth
		} else {
			// For row flex direction, ensure width is at least as wide as the longest word
			longestWordWidth := findLongestWordWidth(text, widthPerChar)
			effectiveWidth = float32(math.Max(float64(longestWordWidth), float64(measuredWidth)))
		}

		// Calculate height with proper text wrapping
		if heightMode == yoga.MeasureModeExactly {
			measuredHeight = height
		} else if heightMode == yoga.MeasureModeAtMost {
			calculatedHeight := calculateTextHeight(text, effectiveWidth, widthPerChar, heightPerChar)
			measuredHeight = float32(math.Min(float64(calculatedHeight), float64(height)))
		} else {
			measuredHeight = calculateTextHeight(text, effectiveWidth, widthPerChar, heightPerChar)
		}

		return yoga.Size{Width: measuredWidth, Height: measuredHeight}
	}
}

// findLongestWordWidth finds the width of the longest word in the text
func findLongestWordWidth(text string, widthPerChar float32) float32 {
	maxLength := 0
	currentLength := 0

	for _, char := range text {
		if char == ' ' {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 0
		} else {
			currentLength++
		}
	}

	if currentLength > maxLength {
		maxLength = currentLength
	}

	return float32(maxLength) * widthPerChar
}

// calculateTextHeight implements proper word-based text wrapping
func calculateTextHeight(text string, measuredWidth, widthPerChar, heightPerChar float32) float32 {
	// If text fits on one line, return single line height
	if float32(len(text))*widthPerChar <= measuredWidth {
		return heightPerChar
	}

	// Split text into words
	words := strings.Split(text, " ")

	// Calculate line wrapping
	lines := 1
	currentLineLength := float32(0)

	for i, word := range words {
		wordWidth := float32(len(word)) * widthPerChar

		if wordWidth > measuredWidth {
			// Word exceeds line width - force new line
			if currentLineLength > 0 {
				lines++
			}
			lines++
			currentLineLength = 0
		} else {
			// Calculate if word fits on current line (including space if not first word)
			spaceWidth := float32(0)
			if currentLineLength > 0 {
				spaceWidth = widthPerChar
			}

			if currentLineLength+spaceWidth+wordWidth <= measuredWidth {
				// Word fits on current line
				currentLineLength += spaceWidth + wordWidth
			} else {
				// Word doesn't fit - start new line
				if i > 0 { // Don't increment lines for the very first word
					lines++
				}
				currentLineLength = wordWidth
			}
		}
	}

	// Handle empty line at the end (same as C++ version)
	if currentLineLength == 0 {
		lines--
	}
	return float32(lines) * heightPerChar
}
