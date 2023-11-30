/**
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

// @generated by gentest/gentest.rb from gentest/fixtures/YGAlignSelfTest.html

package tests

import (
  "testing"
 "github.com/millken/yoga"
 "github.com/stretchr/testify/assert"
)
func TestAlignSelfCenter(t *testing.T) {
  config := yoga.ConfigNew()
  config.SetExperimentalFeatureEnabled(yoga.ExperimentalFeatureAbsolutePercentageAgainstPaddingEdge, true)

  root := yoga.NewNodeWithConfig(config)
  root.StyleSetPositionType(yoga.PositionTypeAbsolute)
  root.StyleSetWidth(100)
  root.StyleSetHeight(100)

  root_child0 := yoga.NewNodeWithConfig(config)
  root_child0.StyleSetAlignSelf(yoga.AlignCenter)
  root_child0.StyleSetWidth(10)
  root_child0.StyleSetHeight(10)
  root.InsertChild(root_child0, 0)
  root.StyleSetDirection(yoga.DirectionLTR)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 45, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 10, root_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child0.LayoutHeight())

  root.StyleSetDirection(yoga.DirectionRTL)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionRTL)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 45, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 10, root_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child0.LayoutHeight())
}

func TestAlignSelfFlexEnd(t *testing.T) {
  config := yoga.ConfigNew()
  config.SetExperimentalFeatureEnabled(yoga.ExperimentalFeatureAbsolutePercentageAgainstPaddingEdge, true)

  root := yoga.NewNodeWithConfig(config)
  root.StyleSetPositionType(yoga.PositionTypeAbsolute)
  root.StyleSetWidth(100)
  root.StyleSetHeight(100)

  root_child0 := yoga.NewNodeWithConfig(config)
  root_child0.StyleSetAlignSelf(yoga.AlignFlexEnd)
  root_child0.StyleSetWidth(10)
  root_child0.StyleSetHeight(10)
  root.InsertChild(root_child0, 0)
  root.StyleSetDirection(yoga.DirectionLTR)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 90, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 10, root_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child0.LayoutHeight())

  root.StyleSetDirection(yoga.DirectionRTL)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionRTL)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 0, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 10, root_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child0.LayoutHeight())
}

func TestAlignSelfFlexStart(t *testing.T) {
  config := yoga.ConfigNew()
  config.SetExperimentalFeatureEnabled(yoga.ExperimentalFeatureAbsolutePercentageAgainstPaddingEdge, true)

  root := yoga.NewNodeWithConfig(config)
  root.StyleSetPositionType(yoga.PositionTypeAbsolute)
  root.StyleSetWidth(100)
  root.StyleSetHeight(100)

  root_child0 := yoga.NewNodeWithConfig(config)
  root_child0.StyleSetAlignSelf(yoga.AlignFlexStart)
  root_child0.StyleSetWidth(10)
  root_child0.StyleSetHeight(10)
  root.InsertChild(root_child0, 0)
  root.StyleSetDirection(yoga.DirectionLTR)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 0, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 10, root_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child0.LayoutHeight())

  root.StyleSetDirection(yoga.DirectionRTL)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionRTL)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 90, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 10, root_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child0.LayoutHeight())
}

func TestAlignSelfFlexEndOverrideFlexStart(t *testing.T) {
  config := yoga.ConfigNew()
  config.SetExperimentalFeatureEnabled(yoga.ExperimentalFeatureAbsolutePercentageAgainstPaddingEdge, true)

  root := yoga.NewNodeWithConfig(config)
  root.StyleSetAlignItems(yoga.AlignFlexStart)
  root.StyleSetPositionType(yoga.PositionTypeAbsolute)
  root.StyleSetWidth(100)
  root.StyleSetHeight(100)

  root_child0 := yoga.NewNodeWithConfig(config)
  root_child0.StyleSetAlignSelf(yoga.AlignFlexEnd)
  root_child0.StyleSetWidth(10)
  root_child0.StyleSetHeight(10)
  root.InsertChild(root_child0, 0)
  root.StyleSetDirection(yoga.DirectionLTR)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 90, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 10, root_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child0.LayoutHeight())

  root.StyleSetDirection(yoga.DirectionRTL)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionRTL)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 0, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 10, root_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child0.LayoutHeight())
}

func TestAlignSelfBaseline(t *testing.T) {
  config := yoga.ConfigNew()
  config.SetExperimentalFeatureEnabled(yoga.ExperimentalFeatureAbsolutePercentageAgainstPaddingEdge, true)

  root := yoga.NewNodeWithConfig(config)
  root.StyleSetFlexDirection(yoga.FlexDirectionRow)
  root.StyleSetPositionType(yoga.PositionTypeAbsolute)
  root.StyleSetWidth(100)
  root.StyleSetHeight(100)

  root_child0 := yoga.NewNodeWithConfig(config)
  root_child0.StyleSetAlignSelf(yoga.AlignBaseline)
  root_child0.StyleSetWidth(50)
  root_child0.StyleSetHeight(50)
  root.InsertChild(root_child0, 0)

  root_child1 := yoga.NewNodeWithConfig(config)
  root_child1.StyleSetAlignSelf(yoga.AlignBaseline)
  root_child1.StyleSetWidth(50)
  root_child1.StyleSetHeight(20)
  root.InsertChild(root_child1, 1)

  root_child1_child0 := yoga.NewNodeWithConfig(config)
  root_child1_child0.StyleSetWidth(50)
  root_child1_child0.StyleSetHeight(10)
  root_child1.InsertChild(root_child1_child0, 0)
  root.StyleSetDirection(yoga.DirectionLTR)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 0, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 50, root_child0.LayoutWidth())
  assert.EqualValues(t, 50, root_child0.LayoutHeight())

  assert.EqualValues(t, 50, root_child1.LayoutLeft())
  assert.EqualValues(t, 40, root_child1.LayoutTop())
  assert.EqualValues(t, 50, root_child1.LayoutWidth())
  assert.EqualValues(t, 20, root_child1.LayoutHeight())

  assert.EqualValues(t, 0, root_child1_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child1_child0.LayoutTop())
  assert.EqualValues(t, 50, root_child1_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child1_child0.LayoutHeight())

  root.StyleSetDirection(yoga.DirectionRTL)
  yoga.CalculateLayout(root,yoga.Undefined, yoga.Undefined, yoga.DirectionRTL)

  assert.EqualValues(t, 0, root.LayoutLeft())
  assert.EqualValues(t, 0, root.LayoutTop())
  assert.EqualValues(t, 100, root.LayoutWidth())
  assert.EqualValues(t, 100, root.LayoutHeight())

  assert.EqualValues(t, 50, root_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child0.LayoutTop())
  assert.EqualValues(t, 50, root_child0.LayoutWidth())
  assert.EqualValues(t, 50, root_child0.LayoutHeight())

  assert.EqualValues(t, 0, root_child1.LayoutLeft())
  assert.EqualValues(t, 40, root_child1.LayoutTop())
  assert.EqualValues(t, 50, root_child1.LayoutWidth())
  assert.EqualValues(t, 20, root_child1.LayoutHeight())

  assert.EqualValues(t, 0, root_child1_child0.LayoutLeft())
  assert.EqualValues(t, 0, root_child1_child0.LayoutTop())
  assert.EqualValues(t, 50, root_child1_child0.LayoutWidth())
  assert.EqualValues(t, 10, root_child1_child0.LayoutHeight())
}