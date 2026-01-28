# Go Sentinel Helper

Utilities and patterns for handling "Unspecified" state in Go, enabling declarative zero-allocation semantics.

## Overview

This library provides a standard way to represent "unspecified" or "optional" values without using pointers (which cause heap allocations) or wrapping types (which add ergonomic friction).

It implements the **Unspecified Sentinel Pattern**, allowing you to distiguish between a "zero value" (like `0` or `""`) and a truly "unspecified" value (which should inherit a default or be merged).

## Installation

```bash
go get github.com/zodimo/go-sentinel-helper
```

## Documentation

For a complete reference on the pattern, contracts, and usage examples, please read:

👉 **[The Unspecified Pattern Reference](docs/sentinel_pattern.md)**

## Packages

The definitions are split by type to keep dependencies minimal:

| Package | Type | Sentinel Value |
|---------|------|----------------|
| [`sentinel/floatutils`](sentinel/floatutils) | `float32`, `float64` | `NaN` (Quiet) |
| [`sentinel/intutils`](sentinel/intutils) | `int` | `math.MinInt` |
| [`sentinel/stringutils`](sentinel/stringutils) | `string` | `"\x00unspecified"` |
| [`sentinel/boolutils`](sentinel/boolutils) | `BooleanValue` | `BooleanValueUnspecified` (Enum) |

## Quick Start

### Defining a Sentinel Type

```go
import "github.com/zodimo/go-sentinel-helper/sentinel/floatutils"

// 1. Define your domain type
type Opacity float32

// 2. Define the Unspecified sentinel
// (This is a zero-allocation constant in .rodata)
var OpacityUnspecified = Opacity(floatutils.Float32Unspecified)

// 3. Add the IsSpecified predicate
func (o Opacity) IsSpecified() bool {
    return floatutils.IsSpecified(o)
}

// 4. Use it!
func ApplyOpacity(o Opacity) {
    if !o.IsSpecified() {
        // Use default
        o = 1.0 
    }
    // ...
}
```

### Merging Values

```go
import "github.com/zodimo/go-sentinel-helper/sentinel/intutils"

func MergeConfigs(base, override int) int {
    // Returns override if specified, otherwise base
    return intutils.MergeIntValue(base, override)
}
```
### Boolean Values (Tri-state)

For booleans, we cannot use a sentinel value (both `true` and `false` are valid). We use a zero-allocation struct wrapper:

```go
import "github.com/zodimo/go-sentinel-helper/sentinel/boolutils"

type FeatureFlag struct {
    Enabled boolutils.BooleanValue
}

func NewFeature(enabled bool) FeatureFlag {
    return FeatureFlag{
        // Convert bool -> BooleanValue
        Enabled: boolutils.BooleanValueFrom(enabled), 
    }
}

func (f FeatureFlag) IsEnabled() bool {
    // defaults to false if unspecified
    return f.Enabled.Bool() 
}
```


# License
Distributed under the MIT License. See `LICENSE` for more information.
