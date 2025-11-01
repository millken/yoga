/**
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/* global Emitter:readable */

const GoEmitter = function () {
  Emitter.call(this, 'go', '  ');
};

function toValueGo(value) {
  value = value.toString().replace('px','').replace('%','');
  if (value.match(/^[0-9.e+-]+px$/i)) return parseFloat(value);
  if (value.match(/^[0-9.e+-]+%/i)) return JSON.stringify(value);
  if (value == 'Yoga.AUTO') return '"auto"';
  if (value == 'max-content') return '"max-content"';
  if (value == 'fit-content') return '"fit-content"';
  if (value == 'stretch') return '"stretch"';
  return value;
}

function toExportName(name) {
  name = name.replace(/(\_\w)/g, function(m) { return m[1].toUpperCase(); });
  if (name.length > 0) {
    name = name[0].toUpperCase() + name.substring(1);
  }
  return name;
}

GoEmitter.prototype = Object.create(Emitter.prototype, {
  constructor: {value: GoEmitter},

  emitPrologue: {
    value: function () {
      this.push([
        'package yoga_test',
        '',
        'import (',
        '  "testing"',
        ' "github.com/millken/yoga"',
        ' "github.com/dnsoa/go/assert"',
        ')',
        '',
      ]);
    },
  },

  emitTestPrologue:{value:function(name, experiments, disabled) {
    this.push('func Test' + toExportName(name) + '(t *testing.T) {');
    this.pushIndent();

    if (disabled) {
      this.push('t.Skip()');
      this.push('');
    }

    this.push('config := yoga.NewConfig()')
    for (var i in experiments) {
      this.push('config.SetExperimentalFeatureEnabled(yoga.ExperimentalFeature' + experiments[i] +', true)');
    }
    this.push('');
  }},

  emitTestTreePrologue: {
    value: function (nodeName) {
      this.push(nodeName + ' := yoga.NewNodeWithConfig(config)');
    },
  },

  emitTestEpilogue: {
    value: function (_experiments) {
      this.push('config.Destroy()');
      this.popIndent();
      this.push('}');
      this.push('');
    },
  },

  emitEpilogue: {
    value: function () {
      this.push('');
    },
  },

  AssertEQ: {
    value: function (v0, v1) {
      this.push('assert.Equal(t, ' + v0 + ', ' + v1 + ')');
    },
  },

  YGAlignAuto: {value: 'yoga.AlignAuto'},
  YGAlignCenter: {value: 'yoga.AlignCenter'},
  YGAlignFlexEnd: {value: 'yoga.AlignFlexEnd'},
  YGAlignFlexStart: {value: 'yoga.AlignFlexStart'},
  YGAlignStretch: {value: 'yoga.AlignStretch'},
  YGAlignSpaceBetween: {value: 'yoga.AlignSpaceBetween'},
  YGAlignSpaceAround: {value: 'yoga.AlignSpaceAround'},
  YGAlignSpaceEvenly: {value: 'yoga.AlignSpaceEvenly'},
  YGAlignBaseline: {value: 'yoga.AlignBaseline'},

  YGDirectionInherit: {value: 'yoga.DirectionInherit'},
  YGDirectionLTR: {value: 'yoga.DirectionLTR'},
  YGDirectionRTL: {value: 'yoga.DirectionRTL'},

  YGEdgeBottom: {value: 'yoga.EdgeBottom'},
  YGEdgeEnd: {value: 'yoga.EdgeEnd'},
  YGEdgeLeft: {value: 'yoga.EdgeLeft'},
  YGEdgeRight: {value: 'yoga.EdgeRight'},
  YGEdgeStart: {value: 'yoga.EdgeStart'},
  YGEdgeTop: {value: 'yoga.EdgeTop'},

  YGGutterAll: {value: 'yoga.GutterAll'},
  YGGutterColumn: {value: 'yoga.GutterColumn'},
  YGGutterRow: {value: 'yoga.GutterRow'},

  YGFlexDirectionColumn: {value: 'yoga.FlexDirectionColumn'},
  YGFlexDirectionColumnReverse: {value: 'yoga.FlexDirectionColumnReverse'},
  YGFlexDirectionRow: {value: 'yoga.FlexDirectionRow'},
  YGFlexDirectionRowReverse: {value: 'yoga.FlexDirectionRowReverse'},

  YGJustifyCenter: {value: 'yoga.JustifyCenter'},
  YGJustifyFlexEnd: {value: 'yoga.JustifyFlexEnd'},
  YGJustifyFlexStart: {value: 'yoga.JustifyFlexStart'},
  YGJustifySpaceAround: {value: 'yoga.JustifySpaceAround'},
  YGJustifySpaceBetween: {value: 'yoga.JustifySpaceBetween'},
  YGJustifySpaceEvenly: {value: 'yoga.JustifySpaceEvenly'},

  YGOverflowHidden: {value: 'yoga.OverflowHidden'},
  YGOverflowVisible: {value: 'yoga.OverflowVisible'},
  YGOverflowScroll: {value: 'yoga.OverflowScroll'},

  YGPositionTypeAbsolute: {value: 'yoga.PositionTypeAbsolute'},
  YGPositionTypeRelative: {value: 'yoga.PositionTypeRelative'},
  YGPositionTypeStatic: {value: 'yoga.PositionTypeStatic'},

  YGAuto: {value: 'auto'},
  YGUndefined: {value: 'undefined'},

  YGWrapNoWrap: {value: 'yoga.WrapNoWrap'},
  YGWrapWrap: {value: 'yoga.WrapWrap'},
  YGWrapWrapReverse: {value: 'yoga.WrapWrapReverse'},

  YGDisplayFlex: {value: 'yoga.DisplayFlex'},
  YGDisplayNone: {value: 'yoga.DisplayNone'},
  YGDisplayContents: {value: 'yoga.DisplayContents'},

  YGBoxSizingBorderBox: {value: 'yoga.BoxSizingBorderBox'},
  YGBoxSizingContentBox: {value: 'yoga.BoxSizingContentBox'},

  YGMaxContent: {value: 'max-content'},
  YGFitContent: {value: 'fit-content'},
  YGStretch: {value: 'stretch'},

  YGNodeCalculateLayout: {
    value: function (node, dir, _experiments) {
      this.push(node + '.CalculateLayout(yoga.Undefined, yoga.Undefined, ' + dir + ')');
    },
  },

  YGNodeInsertChild: {
    value: function (parentName, nodeName, index) {
      this.push(parentName + '.InsertChild(' + nodeName + ', ' + index + ')');
    },
  },

  YGNodeLayoutGetLeft: {
    value: function (nodeName) {
      return nodeName + '.GetComputedLeft()';
    },
  },

  YGNodeLayoutGetTop: {
    value: function (nodeName) {
      return nodeName + '.GetComputedTop()';
    },
  },

  YGNodeLayoutGetWidth: {
    value: function (nodeName) {
      return nodeName + '.GetComputedWidth()';
    },
  },

  YGNodeLayoutGetHeight: {
    value: function (nodeName) {
      return nodeName + '.GetComputedHeight()';
    },
  },

  YGNodeLayoutGetDirection: {
    value: function (nodeName) {
      return nodeName + '.GetLayoutDirection()';
    },
  },

  YGNodeLayoutGetRawWidth: {
    value: function (nodeName) {
      return nodeName + '.GetRawWidth()';
    },
  },

  YGNodeLayoutGetRawHeight: {
    value: function (nodeName) {
      return nodeName + '.GetRawHeight()';
    },
  },

  YGNodeStyleGetFlex: {
    value: function (nodeName) {
      return nodeName + '.GetFlex()';
    },
  },

  YGNodeStyleSetAlignContent: {
    value: function (nodeName, value) {
      this.push(
        nodeName + '.SetAlignContent(' + toValueGo(value) + ')',
      );
    },
  },

  YGNodeStyleSetAlignItems: {
    value: function (nodeName, value) {
      this.push(nodeName + '.SetAlignItems(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetAlignSelf: {
    value: function (nodeName, value) {
      this.push(nodeName + '.SetAlignSelf(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetAspectRatio: {
    value: function (nodeName, value) {
      this.push(
        nodeName + '.SetAspectRatio(' + toValueGo(value) + ')',
      );
    },
  },

  YGNodeStyleSetBorder: {
    value: function (nodeName, edge, value) {
      this.push(
        nodeName +
          '.SetBorder(' +
          toValueGo(edge) +
          ', ' +
          toValueGo(value) +
          ')',
      );
    },
  },

  YGNodeStyleSetDirection: {
    value: function (nodeName, value) {
      this.push(nodeName + '.SetDirection(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetDisplay: {
    value: function (nodeName, value) {
      this.push(nodeName + '.SetDisplay(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetFlexBasis: {
    value: function (nodeName, value) {
      if (value == 'auto') {
        this.push(nodeName + '.SetFlexBasisAuto()');
        return;
      }else if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetFlexBasisPercent(' + toValueGo(value) + ')',
        );
        return 
      }else if (value == 'max-content') {
        this.push(
          nodeName + '.SetFlexBasisMaxContent()',
        );
        return
      }else if (value == 'fit-content') {
        this.push(
          nodeName + '.SetFlexBasisFitContent()',
        );
        return 
      }else if (value == 'stretch') {
        this.push(
          nodeName + '.SetFlexBasisStretch()',
        );
        return
      }
      this.push(nodeName + '.SetFlexBasis(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetFlexDirection: {
    value: function (nodeName, value) {
      this.push(
        nodeName + '.SetFlexDirection(' + toValueGo(value) + ')',
      );
    },
  },

  YGNodeStyleSetFlexGrow: {
    value: function (nodeName, value) {
      this.push(nodeName + '.SetFlexGrow(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetFlexShrink: {
    value: function (nodeName, value) {
      this.push(nodeName + '.SetFlexShrink(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetFlexWrap: {
    value: function (nodeName, value) {
      this.push(nodeName + '.SetFlexWrap(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetHeight: {
    value: function (nodeName, value) {
      if (value == 'auto') {
        this.push(nodeName + '.SetHeightAuto()');
        return;
      }else if (value == 'max-content') {
        this.push(
          nodeName + '.SetHeightMaxContent()',
        );
        return
      }else if (value == 'fit-content') {
        this.push(
          nodeName + '.SetHeightFitContent()',
        );
        return
      }else if (value == 'stretch') {
        this.push(
          nodeName + '.SetHeightStretch()',
        );
        return
      }else if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetHeightPercent(' + toValueGo(value) + ')',
        );
        return
      }
      this.push(nodeName + '.SetHeight(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetJustifyContent: {
    value: function (nodeName, value) {
      this.push(
        nodeName + '.SetJustifyContent(' + toValueGo(value) + ')',
      );
    },
  },

  YGNodeStyleSetMargin: {
    value: function (nodeName, edge, value) {
      if (value == 'auto') {
        this.push(
          nodeName + '.SetMarginAuto(' + toValueGo(edge) + ')',
        );
        return;
      }else if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetMarginPercent(' + toValueGo(edge) + ', ' + toValueGo(value) + ')',
        );
        return
      }
      this.push(
        nodeName +
          '.SetMargin(' +
          toValueGo(edge) +
          ', ' +
          toValueGo(value) +
          ')',
      );
    },
  },

  YGNodeStyleSetMaxHeight: {
    value: function (nodeName, value) {
      if (value == 'max-content') {
        this.push(
          nodeName + '.SetMaxHeightMaxContent()',
        );
        return
      }else if (value == 'fit-content') {
        this.push(
          nodeName + '.SetMaxHeightFitContent()',
        );
        return
      }else if (value == 'stretch') {
        this.push(
          nodeName + '.SetMaxHeightStretch()',
        );
        return 
      }else if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetMaxHeightPercent(' + toValueGo(value) + ')',
        );
        return
      }
      this.push(nodeName + '.SetMaxHeight(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetMaxWidth: {
    value: function (nodeName, value) {
      if (value == 'auto') {
        this.push(nodeName + '.SetMaxWidthAuto()');
        return;
      }else if (value == 'max-content') {
        this.push(
          nodeName + '.SetMaxWidthMaxContent()',
        );
        return
      }else if (value == 'fit-content') {
        this.push(
          nodeName + '.SetMaxWidthFitContent()',
        );
        return
      }else if (value == 'stretch') {
        this.push(
          nodeName + '.SetMaxWidthStretch()',
        );
        return
      }else if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetMaxWidthPercent(' + toValueGo(value) + ')',
        );
        return
      }
      this.push(nodeName + '.SetMaxWidth(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetMinHeight: {
    value: function (nodeName, value) {
      if (value == 'max-content') {
        this.push(
          nodeName + '.SetMinHeightMaxContent()',
        );
        return
      }else if (value == 'fit-content') {
        this.push(
          nodeName + '.SetMinHeightFitContent()',
        );
        return
      }else if (value == 'stretch') {
        this.push(
          nodeName + '.SetMinHeightStretch()',
        );
        return
      }else if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetMinHeightPercent(' + toValueGo(value) + ')',
        );
        return
      }
      this.push(nodeName + '.SetMinHeight(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetMinWidth: {
    value: function (nodeName, value) {
      if (value == 'max-content') {
        this.push(
          nodeName + '.SetMinWidthMaxContent()',
        );
        return
      }else if (value == 'fit-content') {
        this.push(
          nodeName + '.SetMinWidthFitContent()',
        );
        return
      }else if (value == 'stretch') {
        this.push(
          nodeName + '.SetMinWidthStretch()',
        );
        return
      }else if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetMinWidthPercent(' + toValueGo(value) + ')',
        );
        return
      }
      this.push(nodeName + '.SetMinWidth(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetOverflow: {
    value: function (nodeName, value) {
      this.push(nodeName + '.SetOverflow(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetPadding: {
    value: function (nodeName, edge, value) {
      if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetPaddingPercent(' + toValueGo(edge) + ', ' + toValueGo(value) + ')',
        );
        return
      }
      // Note: Padding doesn't support auto, max-content, fit-content, stretch in current API
      // These values are treated as regular float values or ignored
      if (value == 'auto' || value == 'max-content' || value == 'fit-content' || value == 'stretch') {
        this.push(
          `// Skipping padding set for unsupported value: ${value} on edge ${toValueGo(edge)}`,
        );
        return;
      }
      this.push(
        nodeName +
          '.SetPadding(' +
          toValueGo(edge) +
          ', ' +
          toValueGo(value) +
          ')',
      );
    },
  },

  YGNodeStyleSetPosition: {
    value: function (nodeName, edge, value) {
      const valueStr = toValueGo(value);

      if (value == 'auto') {
        this.push(
          nodeName + '.SetPositionAuto(' + toValueGo(edge) + ')',
        );
      }else if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetPositionPercent(' + toValueGo(edge) + ', ' + valueStr + ')',
        );
      } else {
        this.push(
          nodeName +
            '.SetPosition(' +
            toValueGo(edge) +
            ', ' +
            valueStr +
            ')',
        );
      }
    },
  },

  YGNodeStyleSetPositionType: {
    value: function (nodeName, value) {
      this.push(
        nodeName + '.SetPositionType(' + toValueGo(value) + ')',
      );
    },
  },

  YGNodeStyleSetWidth: {
    value: function (nodeName, value) {
      if (value == 'auto') {
        this.push(nodeName + '.SetWidthAuto()');
        return;
      }else if (value == 'max-content') {
        this.push(
          nodeName + '.SetWidthMaxContent()',
        );
        return
      }else if (value == 'fit-content') {
        this.push(
          nodeName + '.SetWidthFitContent()',
        );
        return
      }else if (value == 'stretch') {
        this.push(
          nodeName + '.SetWidthStretch()',
        );
        return
      }else if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName + '.SetWidthPercent(' + toValueGo(value) + ')',
        );
        return
      }
      this.push(nodeName + '.SetWidth(' + toValueGo(value) + ')');
    },
  },

  YGNodeStyleSetGap: {
    value: function (nodeName, gap, value) {
      if (value.match(/^[0-9.e+-]+%$/i)) {
        this.push(
          nodeName +
            '.SetGapPercent(' +
            toValueGo(gap) +
            ', ' +
            toValueGo(value) +
            ')',
        );
        return
      }
      this.push(
        nodeName +
          '.SetGap(' +
          toValueGo(gap) +
          ', ' +
          toValueGo(value) +
          ')',
      );
    },
  },

  YGNodeStyleSetBoxSizing: {
    value: function (nodeName, value) {
      this.push(nodeName + '.SetBoxSizing(' + toValueGo(value) + ')');
    },
  },

  YGNodeSetMeasureFunc: {
    value: function (nodeName, innerText, flexDirection) {
      this.push(`${nodeName}.SetMeasureFunc(intrinsicSizeMeasureFunc("${innerText}", ${flexDirection}));`);
    },
  },
});
