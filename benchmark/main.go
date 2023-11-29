package main

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/millken/yoga"
)

const NUM_REPETITIONS = 1000

func clock() int64 {
	return time.Now().UnixNano()
}

func compareDoubles(a, b *float64) int {
	arg1 := *a
	arg2 := *b

	if arg1 < arg2 {
		return -1
	}

	if arg1 > arg2 {
		return 1
	}

	return 0
}

func printBenchmarkResult(name string, start int64, endTimes []int64) {
	timesInMs := make([]float64, NUM_REPETITIONS)
	mean := float64(0)
	lastEnd := start
	for i := 0; i < NUM_REPETITIONS; i++ {
		timesInMs[i] = float64(endTimes[i] - lastEnd)
		lastEnd = endTimes[i]
		mean += timesInMs[i]
	}
	mean /= NUM_REPETITIONS

	sort.Float64s(timesInMs)
	median := timesInMs[NUM_REPETITIONS/2]

	variance := 0.0
	for i := 0; i < NUM_REPETITIONS; i++ {
		variance += math.Pow(timesInMs[i]-mean, 2)
	}
	variance /= float64(NUM_REPETITIONS)
	stddev := math.Sqrt(variance)

	fmt.Printf("%s: median: %f ms, stddev: %f ms\n", name, median/1e6, stddev/1e6)
}

func measure(node *yoga.Node, width float32, widthMode yoga.MeasureMode, height float32, heightMode yoga.MeasureMode) yoga.YGSize {
	return yoga.YGSize{
		Width:  yoga.If(widthMode == yoga.MeasureModeUndefined, 10, width),
		Height: yoga.If(heightMode == yoga.MeasureModeUndefined, 10, height),
	}

}

func benchmark(name string, benchFn func()) {
	start := clock()
	endTimes := make([]int64, NUM_REPETITIONS)
	for i := 0; i < NUM_REPETITIONS; i++ {
		benchFn()
		endTimes[i] = clock()
	}
	printBenchmarkResult(name, start, endTimes)
}

func main() {
	benchmark("Stack with flex", func() {
		root := yoga.NewNode()
		root.StyleSetWidth(100)
		root.StyleSetHeight(100)

		for i := 0; i < 10; i++ {
			child := yoga.NewNode()
			child.SetMeasureFunc(measure)
			child.StyleSetFlex(1)
			root.InsertChild(child, 0)
		}

		yoga.CalculateLayout(root, yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)
	})

	benchmark("Align stretch in undefined axis", func() {
		root := yoga.NewNode()

		for i := 0; i < 10; i++ {
			child := yoga.NewNode()
			child.StyleSetHeight(20)
			child.SetMeasureFunc(measure)
			root.InsertChild(child, 0)
		}

		yoga.CalculateLayout(root, yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)
	})

	benchmark("Nested flex", func() {
		root := yoga.NewNode()

		for i := 0; i < 10; i++ {
			child := yoga.NewNode()
			child.StyleSetFlex(1)
			root.InsertChild(child, 0)
			for ii := 0; ii < 10; ii++ {
				grandChild := yoga.NewNode()
				grandChild.SetMeasureFunc(measure)
				grandChild.StyleSetFlex(1)
				child.InsertChild(grandChild, 0)
			}
		}

		yoga.CalculateLayout(root, yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)
	})

	benchmark("Huge nested layout", func() {
		root := yoga.NewNode()

		for i := 0; i < 10; i++ {
			child := yoga.NewNode()
			child.StyleSetFlexGrow(1)
			child.StyleSetWidth(10)
			child.StyleSetHeight(10)
			root.InsertChild(child, 0)
			for ii := 0; ii < 10; ii++ {
				grandChild := yoga.NewNode()
				grandChild.StyleSetFlexDirection(yoga.FlexDirectionRow)
				grandChild.StyleSetFlexGrow(1)
				grandChild.StyleSetWidth(10)
				grandChild.StyleSetHeight(10)
				child.InsertChild(grandChild, 0)

				for iii := 0; iii < 10; iii++ {
					greatGrandChild := yoga.NewNode()
					greatGrandChild.StyleSetFlexGrow(1)
					greatGrandChild.StyleSetWidth(10)
					greatGrandChild.StyleSetHeight(10)
					grandChild.InsertChild(greatGrandChild, 0)

					for iiii := 0; iiii < 10; iiii++ {
						greatGreatGrandChild := yoga.NewNode()
						greatGreatGrandChild.StyleSetFlexDirection(yoga.FlexDirectionRow)
						greatGreatGrandChild.StyleSetFlexGrow(1)
						greatGreatGrandChild.StyleSetWidth(10)
						greatGreatGrandChild.StyleSetHeight(10)
						greatGrandChild.InsertChild(greatGreatGrandChild, 0)
					}
				}
			}
		}

		yoga.CalculateLayout(root, yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)
	})
}
