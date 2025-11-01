#!/usr/bin/env python3
# Copyright (c) Meta Platforms, Inc. and affiliates.
#
# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

import os

ENUMS = {
    "Direction": ["Inherit", "LTR", "RTL"],
    "Unit": [
        "Undefined",
        "Point",
        "Percent",
        "Auto",
        "MaxContent",
        "FitContent",
        "Stretch",
    ],
    "FlexDirection": ["Column", "ColumnReverse", "Row", "RowReverse"],
    "Justify": [
        "FlexStart",
        "Center",
        "FlexEnd",
        "SpaceBetween",
        "SpaceAround",
        "SpaceEvenly",
    ],
    "Overflow": ["Visible", "Hidden", "Scroll"],
    "Align": [
        "Auto",
        "FlexStart",
        "Center",
        "FlexEnd",
        "Stretch",
        "Baseline",
        "SpaceBetween",
        "SpaceAround",
        "SpaceEvenly",
    ],
    "PositionType": ["Static", "Relative", "Absolute"],
    "Display": ["Flex", "None", "Contents"],
    "Wrap": ["NoWrap", "Wrap", "WrapReverse"],
    "BoxSizing": ["BorderBox", "ContentBox"],
    "MeasureMode": ["Undefined", "Exactly", "AtMost"],
    "Dimension": ["Width", "Height"],
    "Edge": [
        "Left",
        "Top",
        "Right",
        "Bottom",
        "Start",
        "End",
        "Horizontal",
        "Vertical",
        "All",
    ],
    "NodeType": ["Default", "Text"],
    "LogLevel": ["Error", "Warn", "Info", "Debug", "Verbose", "Fatal"],
    "ExperimentalFeature": [
        # Mimic web flex-basis behavior (experiment may be broken)
        "WebFlexBasis",
    ],
    "Gutter": ["Column", "Row", "All"],
    # Known incorrect behavior which can be enabled for compatibility
    "Errata": [
        # Default: Standards conformant mode
        ("None", 0),
        # Allows main-axis flex basis to be stretched without flexGrow being
        # set (previously referred to as "UseLegacyStretchBehaviour")
        ("StretchFlexBasis", 1 << 0),
        # Absolute position in a given axis will be relative to the padding
        # edge of the parent container instead of the content edge when a
        # specific inset (top/bottom/left/right) is not set.
        ("AbsolutePositionWithoutInsetsExcludesPadding", 1 << 1),
        # Absolute nodes will resolve percentages against the inner size of
        # their containing node, not the padding box
        ("AbsolutePercentAgainstInnerSize", 1 << 2),
        # Enable all incorrect behavior (preserve compatibility)
        ("All", 0x7FFFFFFF),
        # Enable all errata except for "StretchFlexBasis" (Defaults behavior
        # before Yoga 2.0)
        ("Classic", 0x7FFFFFFF & (~(1 << 0))),
    ],
}

DO_NOT_STRIP = ["LogLevel"]

BITSET_ENUMS = ["Errata"]


def get_license(ext):
    return f"""{"/**" if ext == "js" else "/*"}
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

// @{"generated"} by enums.py
{"// clang-format off" if ext == "cpp" else ""}
"""


def _format_name(symbol, delimiter=None, transform=None):
    symbol = str(symbol)
    out = ""
    for i in range(0, len(symbol)):
        c = symbol[i]
        if str.istitle(c) and i != 0 and not str.istitle(symbol[i - 1]):
            out += delimiter or ""
        if transform is None:
            out += c
        else:
            out += getattr(c, transform)()
    return out


def to_java_upper(symbol):
    return _format_name(symbol, "_", "upper")


def to_hyphenated_lower(symbol):
    return _format_name(symbol, "-", "lower")


root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

# Write Go language enum file
with open(root + "/enums_gen.go", "w") as f:
    f.write(get_license("go").strip() + "\n\n")
    f.write("package yoga\n\n")
    
    # Required imports
    f.write("import (\n")
    f.write("\t\"fmt\"\n")
    f.write(")\n\n")
    
    # Generate each enum type
    for enum_name, values in sorted(ENUMS.items()):
        # Write type definition
        if enum_name in BITSET_ENUMS:
            f.write(f"// {enum_name} represents the enum type for {enum_name}\n")
            f.write(f"type {enum_name} uint32\n\n")
        else:
            f.write(f"// {enum_name} represents the enum type for {enum_name}\n")
            f.write(f"type {enum_name} int\n\n")
        
        # Write constants
        f.write("const (\n")
        if enum_name in BITSET_ENUMS:
            # For bitset enums, we need to explicitly set values
            for value in values:
                if isinstance(value, tuple):
                    f.write(f"\t{enum_name}{value[0]} {enum_name} = {value[1]}\n")
                else:
                    f.write(f"\t{enum_name}{value} {enum_name} = {values.index(value)}\n")
        else:
            # Use iota to automatically assign values for other enums
            for i, value in enumerate(values):
                if isinstance(value, tuple):
                    name = value[0]
                    val = value[1]
                    if i == 0:
                        f.write(f"\t{enum_name}{name} {enum_name} = {val}\n")
                    else:
                        f.write(f"\t{enum_name}{name} {enum_name} = {val}\n")
                else:
                    if i == 0:
                        f.write(f"\t{enum_name}{value} {enum_name} = iota\n")
                    else:
                        f.write(f"\t{enum_name}{value}\n")
        f.write(")\n\n")
        
        # Write String method
        f.write(f"// String returns the string representation of {enum_name}\n")
        f.write(f"func (e {enum_name}) String() string {{\n")
        f.write(f"\tswitch e {{\n")
        for value in values:
            if isinstance(value, tuple):
                name = value[0]
                f.write(f"\tcase {enum_name}{name}:\n")
                f.write(f"\t\treturn \"{to_hyphenated_lower(name)}\"\n")
            else:
                f.write(f"\tcase {enum_name}{value}:\n")
                f.write(f"\t\treturn \"{to_hyphenated_lower(value)}\"\n")
        f.write("\tdefault:\n")
        f.write("\t\treturn \"unknown\"\n")
        f.write("\t}\n")
        f.write("}\n\n")
        
        # Write FromString method
        f.write(f"// {enum_name}FromString parses {enum_name} enum value from string\n")
        f.write(f"func {enum_name}FromString(s string) ({enum_name}, error) {{\n")
        f.write(f"\tswitch s {{\n")
        for value in values:
            if isinstance(value, tuple):
                name = value[0]
                f.write(f"\tcase \"{to_hyphenated_lower(name)}\":\n")
                f.write(f"\t\treturn {enum_name}{name}, nil\n")
            else:
                f.write(f"\tcase \"{to_hyphenated_lower(value)}\":\n")
                f.write(f"\t\treturn {enum_name}{value}, nil\n")
        f.write("\tdefault:\n")
        f.write(f"\t\treturn 0, fmt.Errorf(\"unknown {enum_name}: %s\", s)\n")
        f.write("\t}\n")
        f.write("}\n\n")
