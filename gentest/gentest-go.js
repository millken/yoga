/**
 * Copyright (c) 2014-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

function toValueGo(value) {
    return value.toString().replace('px','').replace('%','');
  }
  
  function toMethodName(value) {
    if (value.indexOf('%') >= 0){
      return 'Percent';
    } else if(value.indexOf('Auto') >= 0) {
      return 'Auto';
    }
    return '';
  }
  
  function toExportName(name) {
    name = name.replace(/(\_\w)/g, function(m) { return m[1].toUpperCase(); });
    if (name.length > 0) {
      name = name[0].toUpperCase() + name.substring(1);
    }
    return name;
  }
  
  var GoEmitter = function() {
    Emitter.call(this, 'go', '  ');
  };
  
  GoEmitter.prototype = Object.create(Emitter.prototype, {
    constructor:{value:GoEmitter},
  
    emitPrologue:{
        value: function () {
            this.push([
                'package tests',
                '',
                'import (',
                '  "testing"',

                ' "github.com/millken/yoga"',
                ' "github.com/stretchr/testify/assert"',
                ')',
              ]);
          },
    },
  
    emitTestPrologue:{value:function(name, experiments) {
      this.push('func Test' + toExportName(name) + '(t *testing.T) {');
      this.pushIndent();
  
      this.push('config := yoga.ConfigNew()')
      for (var i in experiments) {
        this.push('config.SetExperimentalFeatureEnabled(Experiment' + experiments[i] +', true)');
      }
      this.push('');
    }},
  
    emitTestTreePrologue:{value:function(nodeName) {
      this.push(nodeName + ' := yoga.NewNodeWithConfig(config)');
    }},
  
    emitTestEpilogue:{value:function(experiments) {
      this.popIndent();
      this.push('}');
      this.push('');
    }},
  
    emitEpilogue:{value:function(lines) {}},
  
    AssertEQ:{value:function(v0, v1) {
      this.push('assert.EqualValues(t, ' + v0 + ', ' + v1 + ')');
    }},
  
    YGAlignAuto:{value:'AlignAuto'},
    YGAlignCenter:{value:'AlignCenter'},
    YGAlignFlexEnd:{value:'AlignFlexEnd'},
    YGAlignFlexStart:{value:'AlignFlexStart'},
    YGAlignStretch:{value:'AlignStretch'},
    YGAlignSpaceBetween:{value:'AlignSpaceBetween'},
    YGAlignSpaceAround:{value:'AlignSpaceAround'},
    YGAlignBaseline:{value:'AlignBaseline'},
  
    YGDirectionInherit:{value:'DirectionInherit'},
    YGDirectionLTR:{value:'DirectionLTR'},
    YGDirectionRTL:{value:'DirectionRTL'},
  
    YGEdgeBottom:{value:'EdgeBottom'},
    YGEdgeEnd:{value:'EdgeEnd'},
    YGEdgeLeft:{value:'EdgeLeft'},
    YGEdgeRight:{value:'EdgeRight'},
    YGEdgeStart:{value:'EdgeStart'},
    YGEdgeTop:{value:'EdgeTop'},
  
    YGFlexDirectionColumn:{value:'FlexDirectionColumn'},
    YGFlexDirectionColumnReverse:{value:'FlexDirectionColumnReverse'},
    YGFlexDirectionRow:{value:'FlexDirectionRow'},
    YGFlexDirectionRowReverse:{value:'FlexDirectionRowReverse'},
  
    YGJustifyCenter:{value:'JustifyCenter'},
    YGJustifyFlexEnd:{value:'JustifyFlexEnd'},
    YGJustifyFlexStart:{value:'JustifyFlexStart'},
    YGJustifySpaceAround:{value:'JustifySpaceAround'},
    YGJustifySpaceBetween:{value:'JustifySpaceBetween'},
    YGJustifySpaceEvenly:{value:'JustifySpaceEvenly'},
  
    YGOverflowHidden:{value:'OverflowHidden'},
    YGOverflowVisible:{value:'OverflowVisible'},
  
    YGPositionTypeAbsolute:{value:'PositionAbsolute'},
    YGPositionTypeRelative:{value:'PositionRelative'},
  
    YGUndefined:{value:'Undefined'},
  
    YGDisplayFlex:{value:'DisplayFlex'},
    YGDisplayNone:{value:'DisplayNone'},
    YGAuto:{value:'Auto'},
  
  
    YGWrapNoWrap:{value:'WrapNone'},
    YGWrapWrap:{value:'WrapWrap'},
    YGWrapWrapReverse:{value: 'WrapReverse'},
  
    YGNodeCalculateLayout:{value:function(node, dir, experiments) {
      this.push(node + '.StyleSetDirection(' + dir + ')');
      this.push(node + '.CalculateLayout(Undefined, Undefined, ' + dir +')');
    }},
  
    YGNodeInsertChild:{value:function(parentName, nodeName, index) {
      this.push(parentName + '.InsertChild(' + nodeName + ', ' + index + ')');
    }},
  
    YGNodeLayoutGetLeft:{value:function(nodeName) {
      return nodeName + '.LayoutLeft()';
    }},
  
    YGNodeLayoutGetTop:{value:function(nodeName) {
      return nodeName + '.LayoutTop()';
    }},
  
    YGNodeLayoutGetWidth:{value:function(nodeName) {
      return nodeName + '.LayoutWidth()';
    }},
  
    YGNodeLayoutGetHeight:{value:function(nodeName) {
      return nodeName + '.LayoutHeight()';
    }},
  
    YGNodeStyleSetAlignContent:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetAlignContent(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetAlignItems:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetAlignItems(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetAlignSelf:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetAlignSelf(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetBorder:{value:function(nodeName, edge, value) {
      this.push(nodeName + '.StyleSetBorder(' + edge + ', ' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetDirection:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetDirection(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetDisplay:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetDisplay(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetFlexBasis:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetFlexBasis' + toMethodName(value) + '(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetFlexDirection:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetFlexDirection(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetFlexGrow:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetFlexGrow(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetFlexShrink:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetFlexShrink(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetFlexWrap:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetFlexWrap(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetHeight:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetHeight' + toMethodName(value) + '(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetJustifyContent:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetJustifyContent(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetMargin:{value:function(nodeName, edge, value) {
      var valueStr = toValueGo(value);
      if (valueStr != 'Auto') {
        valueStr = ', ' + valueStr + '';
      } else {
        valueStr = '';
      }
  
      this.push(nodeName + '.StyleSetMargin' + toMethodName(value) + '(' + edge + valueStr + ')');
    }},
  
    YGNodeStyleSetMaxHeight:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetMaxHeight' + toMethodName(value) + '(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetMaxWidth:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetMaxWidth' + toMethodName(value) + '(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetMinHeight:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetMinHeight' + toMethodName(value) + '(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetMinWidth:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetMinWidth' + toMethodName(value) + '(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetOverflow:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetOverflow(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetPadding:{value:function(nodeName, edge, value) {
      this.push(nodeName + '.StyleSetPadding' + toMethodName(value) + '(' + edge + ', ' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetPosition:{value:function(nodeName, edge, value) {
      this.push(nodeName + '.StyleSetPosition' + toMethodName(value) + '(' + edge + ', ' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetPositionType:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetPositionType(' + toValueGo(value) + ')');
    }},
  
    YGNodeStyleSetWidth:{value:function(nodeName, value) {
      this.push(nodeName + '.StyleSetWidth' + toMethodName(value) + '(' + toValueGo(value) + ')');
    }},
  });