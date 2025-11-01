# Yoga Layout for Go

[![Go Test](https://github.com/millken/yoga/actions/workflows/test.yml/badge.svg)](https://github.com/millken/yoga/actions/workflows/test.yml)

Go bindings for Facebook's Yoga Layout library - a cross-platform layout engine implementing Flexbox.

## Features

- 🎯 **Complete Flexbox Implementation** - Full support for CSS Flexbox properties
- 📱 **Cross-Platform** - Works on all platforms supported by Go
- 🔧 **Simple API** - Clean, idiomatic Go interface
- 🎨 **Layout Debugging** - Built-in support for generating layout representations
- ⚡ **High Performance** - Direct C bindings for optimal performance

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/millken/yoga"
)

func main() {
    // Create a root node
    root := yoga.NewNode()
    defer root.Destroy()

    // Set flexbox properties
    root.SetWidth(375)
    root.SetHeight(667)
    root.SetFlexDirection(yoga.FlexDirectionColumn)

    // Create a child node
    child := yoga.NewNode()
    defer child.Destroy()
    child.SetFlexGrow(1)
    child.SetHeight(100)

    // Add child to parent
    root.InsertChild(child, 0)

    // Calculate layout
    root.CalculateLayout(yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)

    // Print computed layout
    fmt.Printf("Child width: %.2f, height: %.2f\n",
        child.GetComputedWidth(), child.GetComputedHeight())

    // Generate layout representation for yogalayout.dev
    layout := root.ToYogaLayout()
    fmt.Println(layout)
}
```

## Layout Debugging

The library provides built-in support for generating layout representations that can be used with [yogalayout.dev playground](https://yogalayout.dev/playground):

```go
root := yoga.NewNode()
defer root.Destroy()

root.SetWidth(200)
root.SetHeight(200)

child := yoga.NewNode()
defer child.Destroy()
child.SetFlexGrow(1)
child.SetHeight(100)

root.InsertChild(child, 0)
root.CalculateLayout(yoga.Undefined, yoga.Undefined, yoga.DirectionLTR)

layout := root.ToYogaLayout()
fmt.Println(layout)
```

Output:
```jsx
<Layout config={{useWebDefaults: false}}>
    <Node style={{width: 200, height: 200}}>
    <Node style={{height: 100, flex: 1}} />
  </Node>
</Layout>
```

## Installation

```bash
go get github.com/millken/yoga
```

## Documentation

- [Yoga Layout Documentation](https://www.yogalayout.dev/docs/about-yoga) - Official Yoga Layout documentation
- [API Reference](https://pkg.go.dev/github.com/millken/yoga) - Go package documentation

## Status

**⚠️ Development Version** - This is a development version and should not be used in production environments.

## License

MIT License - see LICENSE file for details.