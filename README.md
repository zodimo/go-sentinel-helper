# Go Sentinel Helper

Utilities and patterns for handling "Unspecified" state in Go, enabling declarative zero-allocation semantics.

## Overview

This library provides a standard way to represent "unspecified" or "optional" values without using pointers (which cause heap allocations) or wrapping types (which add ergonomic friction).

It implements the **Unspecified Sentinel Pattern**, allowing you to distinguish between a "zero value" (like `0` or `""`) and a truly "unspecified" value (which should inherit a default or be merged).

Every helper package exposes **exactly one public value: its Companion** — a single
generic value that bundles the ten canonical operations (`Unspecified`, `String`,
`IsSpecified`, `IsUnspecified`, `TakeOrElse`, `Merge`, `Same`, `SemanticEqual`, `Equal`,
and `Copy`). The underlying helper funcs are package-private.

## Installation

```bash
go get github.com/zodimo/go-sentinel-helper
```

## Documentation

For a complete reference on the pattern, contracts, and usage examples, please read:

👉 **[The Unspecified Pattern Reference](docs/sentinel_pattern.md)**

## Packages

The definitions are split by type to keep dependencies minimal:

| Package | Type | Companion | Sentinel access |
|---------|------|-----------|-----------------|
| [`sentinel/floatutils`](sentinel/floatutils) | `float32`, `float64` | `Float32Companion`, `Float64Companion` | `Companion.Unspecified()` |
| [`sentinel/intutils`](sentinel/intutils) | `int` | `IntValueCompanion` | `Companion.Unspecified()` |
| [`sentinel/stringutils`](sentinel/stringutils) | `string` | `StringsCompanion` | `Companion.Unspecified()` |
| [`sentinel/boolutils`](sentinel/boolutils) | `BooleanValue` | `BooleanValueCompanion` | `Companion.Unspecified()` |
| [`sentinel/protobufwrapper`](sentinel/protobufwrapper) | `wrapperspb.*Value` | `StringValueCompanion`, `Int32ValueCompanion`, ... (9 total) | `Companion.Unspecified()` |
| [`sentinel/companions`](sentinel/companions) | — | `ScalarCompanion[T]`, `StructCompanion[T]` | the companion machinery itself |

**The sentinel itself is never public.** The reserved value is unexported and reachable only
through `Companion.Unspecified()` — that is the single way to obtain it, for the library's own
packages and for any companion you define. What exactly the sentinel is (NaN, `math.MinInt`,
`"\x00unspecified"`, a singleton) is an internal choice; consult the package source if you
need to reason about it, never a constant.

Beyond the companion, a few packages export documented *extras*: `floatutils` adds
`IsInfinite`, `Float32Equals` and `Float64Equals`; `boolutils` keeps `IsTrue`, `IsFalse`,
`Bool` and `BoolOrElse` on the `BooleanValue` type.

## Quick Start

### Using a Companion Directly

```go
import "github.com/zodimo/go-sentinel-helper/sentinel/stringutils"

c := stringutils.StringsCompanion

c.IsSpecified("hello")            // true
c.IsSpecified("")                 // true  ("" is a meaningful value)
c.IsUnspecified("")               // false
c.Merge("a", "b")                 // "b"
c.TakeOrElse("", "default")       // ""    ("" is specified)
c.String("hi")                    // `StringValue{"hi"}`
c.Copy("hi")                      // "hi"
c.Unspecified()                   // the reserved sentinel value
```

### Defining a Sentinel Type With a Companion

Define your domain type, then register its ten package-private funcs on a companion.
This mirrors `sentinel/stringutils` — the reference implementation.

```go
import (
    "fmt"

    "github.com/zodimo/go-sentinel-helper/sentinel/companions"
    "github.com/zodimo/go-sentinel-helper/sentinel/floatutils"
)

// 1. Define your domain type
type Opacity float32

// 2. Define the Unspecified sentinel
// (zero-allocation: the constant lives in .rodata)
// No sentinel is exported; get it from the companion.
var opacityUnspecified = Opacity(floatutils.Float32Companion.Unspecified())

func isSpecifiedOpacity(o Opacity) bool   { return floatutils.Float32Companion.IsSpecified(float32(o)) }
func isUnspecifiedOpacity(o Opacity) bool { return floatutils.Float32Companion.IsUnspecified(float32(o)) }
func takeOrElseOpacity(o, def Opacity) Opacity {
    if !isSpecifiedOpacity(o) {
        return def
    }
    return o
}
func mergeOpacity(a, b Opacity) Opacity {
    if isSpecifiedOpacity(b) {
        return b
    }
    return a
}
func stringOpacity(o Opacity) string { return fmt.Sprintf("Opacity{%v}", float32(o)) }
func sameOpacity(a, b Opacity) bool  { return a == b }
func semanticEqualOpacity(a, b Opacity) bool { return a == b }
func equalOpacity(a, b Opacity) bool { return a == b }
func copyOpacity(o Opacity) Opacity  { return o }

// 3. Export ONLY the companion.
var OpacityCompanion = companions.NewScalarCompanion(
    func() Opacity { return opacityUnspecified },
    stringOpacity, isSpecifiedOpacity, isUnspecifiedOpacity,
    takeOrElseOpacity, mergeOpacity,
    sameOpacity, semanticEqualOpacity, equalOpacity, copyOpacity,
)

// 4. Use it!
func ApplyOpacity(o Opacity) {
    o = OpacityCompanion.TakeOrElse(o, 1.0) // default to fully opaque
    // ...
}
```

### Merging Values

```go
import "github.com/zodimo/go-sentinel-helper/sentinel/intutils"

func MergeConfigs(base, override int) int {
    // Returns override if specified, otherwise base
    return intutils.IntValueCompanion.Merge(base, override)
}
```

### Boolean Values (Tri-state)

For booleans, we cannot use a sentinel value (both `true` and `false` are valid). We use a
zero-allocation struct wrapper with a dedicated companion:

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

### Protobuf Wrappers (Struct Companions)

For `wrapperspb.*Value` pointers, use the matching StructCompanion. It is nil-safe and
provides `Coalesce` plus option-based `Copy`:

```go
import (
    "github.com/zodimo/go-sentinel-helper/sentinel/protobufwrapper"
    "google.golang.org/protobuf/types/known/wrapperspb"
)

c := protobufwrapper.StringValueCompanion

c.IsSpecified(nil)                                        // false
c.Merge(nil, &wrapperspb.StringValue{Value: "x"}).Value   // "x"
c.Copy(&wrapperspb.StringValue{Value: "x"}, func(v *wrapperspb.StringValue) {
    v.Value = "y"
}).Value                                                  // "y"
```


# Possible future extensions

## Validation

```go
import (
    "github.com/bufbuild/protovalidate-go"
)

var validator, _ = protovalidate.New()

func ValidateUserProfile(u UserProfile) error {
    if !IsSpecifiedUserProfile(u) {
        return nil // Or return a "required" error if the object is mandatory
    }
    return validator.Validate(u.data)
}
```


# License
Distributed under the MIT License. See `LICENSE` for more information.
