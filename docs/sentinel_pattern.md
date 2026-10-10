# The Unspecified Pattern: Complete Go Reference  
*Declarative composition with zero-allocation semantics — powered by Companions*

---

> **AGENT CONTRACT**: Every sentinel helper package exports its **Companion(s)** as its
> public surface. The ten canonical operations live behind that value, and every helper used
> by a domain type is registered with the companion at construction time.
>
> 1. `Unspecified() T` – the sentinel (or its singleton)
> 2. `IsSpecified(T) bool` – predicate
> 3. `IsUnspecified(T) bool` – inverse predicate
> 4. `TakeOrElse(src, def T) T` – 2-param fallback
> 5. `Merge(a, b T) T` – composition merge
> 6. `String(T) string` – stringification
> 7. `Same(a, b T) bool` – identity
> 8. `SemanticEqual(a, b T) bool` – semantic equality
> 9. `Equal(a, b T) bool` – composed equality (`Same` fast-path, else `SemanticEqual`)
> 10. `Copy(...)` – copy (value types: identity; struct types: functional options)

> **Scalar companions** add `Copy(s T) T`; **Struct companions** add `Coalesce(ptr, def T) T`
> and `Copy(s T, options ...func(T)) T`. All verification commands are in Section 9.
>
> A package may export **more than one** companion when the underlying type is generic over
> widths — `floatutils` exports both `Float32Companion` and `Float64Companion`. Sentinel
> *values* are **never** exported — each lives as an unexported package-level symbol and is
> reachable only through `Companion.Unspecified()`. This holds for every package built on the
> pattern, including one you define yourself: do not add an exported `XUnspecified` var, const,
> or func as a convenience. Genuine domain predicates that are **not** among the ten (e.g.
> `IsInfinite`, `IsTrue`) may also be exported — as a package-level func or a method on the
> value type — and must be documented as extras.

---

## Comments on Merge and Copy

- They share the same "prefer incoming specified value" logic but at different granularity.
- `Merge` is used for **composition**; `Copy` is used for **mutation**.
- The companion is the *only* public surface. The ten funcs are package-private and are wired
  into the companion — callers never see them.

---

## 0. Quick Glance Cheat-Sheet

| Type Family | Go Declaration | Companion Kind | Sentinel (unexported) | Is-Specified Check | Alloc/Cost | Use-Case |
|-------------|----------------|----------------|-----------------------|--------------------|------------|----------|
| **1-A Primitive** | `type Dp float32` | Scalar | `var dpUnspecified = Dp(NaN)` | `c.IsSpecified(d)` | **0 B / 2 ns** | dp, simple units |
| **1-B Packed Primitive** | `type Color uint64` | Scalar (+`Copy` options) | `const colorUnspecified = 0` | `c.IsSpecified(col)` | **0 B / 2 ns** | color with copy options |
| **1-C Complex** | `type TextStyle struct{...}` | Struct (pointer) | `var textStyleUnspecified = &TextStyle{...}` | `c.IsSpecified(s)` | **0 B if nil / 24 B if custom** | style, modifier |
| **1-D Type-Safe Wrapper** | `type TextUnit struct{ packed int64 }` | Scalar | `var textUnitUnspecified = TextUnit{...}` | `c.IsSpecified(tu)` | **0 B / 2 ns** | packed unit with type safety |

Every sentinel above is **unexported**. Callers obtain it only via `c.Unspecified()`; the names
are shown lowercase to make clear there is no exported `DpUnspecified` / `ColorUnspecified` /
`TextStyleUnspecified` / `TextUnitUnspecified` to reach for.

---

## 1. The Companion

A **Companion** is a single generic value that bundles the ten canonical operations for one
sentinel type. It is defined in `sentinel/companions`:

```go
// sentinel/companions/companion.go
package companions

type Companion[T any] interface {
	Unspecified() T
	String(T) string
	IsSpecified(T) bool
	IsUnspecified(T) bool
	TakeOrElse(src, def T) T
	Merge(a, b T) T
	Same(a, b T) bool
	SemanticEqual(a, b T) bool
	Equal(a, b T) bool
}

// ScalarCompanion is for value types (no nil possible).
type ScalarCompanion[T any] interface {
	Companion[T]
	Copy(s T) T
}

// StructCompanion is for pointer types (nil is possible).
type StructCompanion[T any] interface {
	Companion[T]
	Coalesce(ptr, def T) T
	Copy(s T, options ...func(T)) T
}
```

Each helper package constructs its companion once, at package init, via
`NewScalarCompanion` / `NewStructCompanion`. The ten funcs are unexported; **only the
companion is exported**.

### 1.1 The canonical worked example — `stringutils`

This is the reference implementation. Note the ten numbered funcs are **unexported** and the
only export is `StringsCompanion`.

```go
// sentinel/stringutils/strings.go
package stringutils

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
)

type StringValue = string

// 1. Sentinel — "\x00unspecified" is used because "" is a meaningful value.
const stringValueUnspecified StringValue = "\x00unspecified"

// 2. IsSpecified — predicate (package-level, unexported)
func isSpecifiedString(s StringValue) bool {
	return s != stringValueUnspecified
}

// IsUnspecified — inverse predicate
func isUnspecifiedString(s StringValue) bool {
	return s == stringValueUnspecified
}

// 3. TakeOrElse — 2-param fallback
func takeOrElseString(a, b StringValue) StringValue {
	if a != stringValueUnspecified {
		return a
	}
	return b
}

// 4. Merge — composition merge (prefers incoming specified values)
func mergeString(a, b StringValue) StringValue {
	if b != stringValueUnspecified {
		return b
	}
	return a
}

// 5. String — stringification
func stringString(s StringValue) string {
	if s == stringValueUnspecified {
		return "StringValue{Unspecified}"
	}
	return fmt.Sprintf("StringValue{%q}", s)
}

// 6. Coalesce — N/A for value types (no nil possible)

// 7. Same — identity
func sameString(a, b StringValue) bool {
	return a == b
}

// 8. SemanticEqual — semantic equality (same as Same for strings)
func semanticEqualString(a, b StringValue) bool {
	return a == b
}

// 9. Equal — equality check
func equalString(a, b StringValue) bool {
	return a == b
}

// 10. Copy — identity for immutable value types
func copyString(s StringValue) StringValue {
	return s
}

// The ONLY public symbol: the companion.
var StringsCompanion = companions.NewScalarCompanion(
	func() StringValue { return stringValueUnspecified },
	stringString,
	isSpecifiedString,
	isUnspecifiedString,
	takeOrElseString,
	mergeString,
	sameString,
	semanticEqualString,
	equalString,
	copyString,
)
```

**Calling it**:

```go
import "github.com/zodimo/go-sentinel-helper/sentinel/stringutils"

stringutils.StringsCompanion.IsSpecified("hello")       // true
stringutils.StringsCompanion.IsSpecified("")            // true  ("" is meaningful)
stringutils.StringsCompanion.IsUnspecified("")          // false
stringutils.StringsCompanion.Merge("a", "b")            // "b"
stringutils.StringsCompanion.TakeOrElse("", "def")      // ""    ("" is specified)
stringutils.StringsCompanion.String("hi")               // `StringValue{"hi"}`
stringutils.StringsCompanion.Copy("hi")                 // "hi"
```

> **Why the sentinel is private:** downstream code cannot reference
> `stringValueUnspecified` directly. It must go through the companion —
> `StringsCompanion.Unspecified()` returns it. This keeps the sentinel an
> implementation detail and guarantees every consumer sees the *same* value.

---

## 2. The Four Sentinel Patterns

Choose **once per type** — never mix. Each maps to either a **scalar** companion
(value types) or a **struct** companion (pointer types).

### 2-A Primitive / Value Types (zero-allocation sentinel)

The value is stored in a primitive. The sentinel is a bit-pattern that is impossible in the
domain.

```go
// unit.Dp

import (
	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	"github.com/zodimo/go-sentinel-helper/sentinel/floatutils"
)

type Dp float32

// floatutils exposes concrete width companions; delegate to the float32 one.
var f32 = floatutils.Float32Companion

// 1. Sentinel — unexported; only reachable via DpCompanion.Unspecified()
var dpUnspecified = Dp(f32.Unspecified())

// 2. IsSpecified
func isSpecifiedDp(d Dp) bool { return f32.IsSpecified(float32(d)) }

// 3. IsUnspecified
func isUnspecifiedDp(d Dp) bool { return f32.IsUnspecified(float32(d)) }

// 4. TakeOrElse
func takeOrElseDp(d, def Dp) Dp {
	if !isSpecifiedDp(d) {
		return def
	}
	return d
}

// 5. Merge
func mergeDp(a, b Dp) Dp {
	if isSpecifiedDp(b) {
		return b
	}
	return a
}

// 6. String
func stringDp(d Dp) string { return fmt.Sprintf("Dp{%v}", float32(d)) }

// 7. Same
func sameDp(a, b Dp) bool { return float32(a) == float32(b) }

// 8. SemanticEqual
func semanticEqualDp(a, b Dp) bool { return f32.SemanticEqual(float32(a), float32(b)) }

// 9. Equal
func equalDp(a, b Dp) bool { return semanticEqualDp(a, b) }

// 10. Copy
func copyDp(d Dp) Dp { return d }

// The only export:
var DpCompanion = companions.NewScalarCompanion(
	func() Dp { return dpUnspecified },
	stringDp, isSpecifiedDp, isUnspecifiedDp,
	takeOrElseDp, mergeDp, sameDp, semanticEqualDp, equalDp, copyDp,
)
```

**Sentinel choice guide**
- `0` → OK when `0` is rare
- `Float32Companion.Unspecified()` / `Float64Companion.Unspecified()` (NaN) → use when `0` is common
- `IntValueCompanion.Unspecified()` (`math.MinInt`) → for integers
- `iota + sentinel` → for enums
- `"\x00unspecified"` → for strings when `""` is meaningful (`stringutils`)

**Extras beyond the ten**: `floatutils` also exports the comparison helper `IsInfinite`
(a package-level generic func) and `Float32Equals`/`Float64Equals`. These are *not* part of
the companion contract; they are documented extras.

**Guarantees**: lives in `.rodata`, copied into register/stack → **zero heap bytes**.

---

### 2-B Packed Primitive / Packed Value Types (zero-allocation sentinel)

Same as 2-A, but `Copy` accepts functional options backed by sentinel fields, so "set to
zero" and "no change" are distinguishable.

```go
type ColorCopyOptions struct {
	Alpha, Red, Green, Blue float32
}
type ColorCopyOption func(*ColorCopyOptions) ColorCopyOptions

func CopyWithAlpha(alpha float32) ColorCopyOption {
	return func(o *ColorCopyOptions) ColorCopyOptions {
		o.Alpha = alpha
		return *o
	}
}
// ...CopyWithRed/Green/Blue identical...

// Copy creates a new color with modified components.
func (c Color) Copy(opts ...ColorCopyOption) Color {
	f32 := floatutils.Float32Companion
	o := ColorCopyOptions{
		Alpha: f32.Unspecified(),
		Red:   f32.Unspecified(),
		Green: f32.Unspecified(),
		Blue:  f32.Unspecified(),
	}
	for _, opt := range opts {
		opt(&o)
	}

	id := c.ColorSpaceId()
	space := colorspace.Get(id)
	return NewColor(
		f32.TakeOrElse(o.Alpha, c.Alpha()),
		f32.TakeOrElse(o.Red, c.Red()),
		f32.TakeOrElse(o.Green, c.Green()),
		f32.TakeOrElse(o.Blue, c.Blue()),
		space,
	)
}
```

The remaining nine operations follow 2-A. Because `Color` remains a value type, it is still
registered with a **ScalarCompanion** — `Copy` is the only operation whose shape changes.

---

### 2-C Complex Object Types (nullable pointer + singleton)

See Section 3 for the full `TextStyle` example. Registered with a **StructCompanion**.

---

### 2-D Type-Safe Struct Wrappers (zero-allocation, compile-time safety)

Use this pattern to **prevent implicit conversions** while keeping **zero-allocation value
semantics**. This is for single-field struct wrappers around packed primitives.

**When to use**:
- The underlying type is a primitive (`int64`, `uint64`) but allows unwanted implicit conversions
- You want `Foo(24)` to be a compile error, forcing users through `NewFoo(24)`
- The type is semantically atomic (not a collection of optional fields)
- Zero-allocation is critical

**Example: TextUnit**

```go
// TextUnit packs unit type (Sp/Em) and float value into 64 bits.
// Struct wrapper prevents `TextUnit(24)` — must use `Sp(24)` or `Em(1.5)`.
type TextUnit struct {
	packed int64
}

// 1. Sentinel — value, not pointer; unexported (reach it via TextUnitCompanion.Unspecified())
var textUnitUnspecified = TextUnit{packed: packRaw(int64(TextUnitTypeUnspecified), float32(math.NaN()))}

func packRaw(unitType int64, v float32) int64 {
	valBits := int64(math.Float32bits(v)) & 0xFFFFFFFF
	return unitType | valBits
}

// Constructors — the ONLY way to create valid TextUnits
func Sp(value float32) TextUnit { return TextUnit{packed: packRaw(int64(TextUnitTypeSp), value)} }
func Em(value float32) TextUnit { return TextUnit{packed: packRaw(int64(TextUnitTypeEm), value)} }

// 2. IsSpecified
func isSpecifiedTextUnit(tu TextUnit) bool {
	return tu.rawType() != int64(TextUnitTypeUnspecified)
}

// 3. IsUnspecified
func isUnspecifiedTextUnit(tu TextUnit) bool { return !isSpecifiedTextUnit(tu) }

// 4. TakeOrElse
func takeOrElseTextUnit(tu, def TextUnit) TextUnit {
	if isSpecifiedTextUnit(tu) {
		return tu
	}
	return def
}

// 5. Merge — whole-value replacement (atomic type)
func mergeTextUnit(a, b TextUnit) TextUnit {
	if isSpecifiedTextUnit(b) {
		return b
	}
	return a
}

// 6. String
func stringTextUnit(tu TextUnit) string {
	if !isSpecifiedTextUnit(tu) {
		return "TextUnit{Unspecified}"
	}
	return fmt.Sprintf("TextUnit{%v %s}", tu.Value(), tu.Type())
}

// 7-9. Equality
func sameTextUnit(a, b TextUnit) bool { return a.packed == b.packed }
func semanticEqualTextUnit(a, b TextUnit) bool {
	if a.Type() != b.Type() {
		return false
	}
	if !isSpecifiedTextUnit(a) {
		return !isSpecifiedTextUnit(b)
	}
	return floatutils.Float32Equals(a.Value(), b.Value(), floatutils.Float32EqualityThreshold)
}
func equalTextUnit(a, b TextUnit) bool {
	if sameTextUnit(a, b) {
		return true
	}
	return semanticEqualTextUnit(a, b)
}

// 10. Copy — identity for immutable value types
func copyTextUnit(tu TextUnit) TextUnit { return tu }

var TextUnitCompanion = companions.NewScalarCompanion(
	func() TextUnit { return textUnitUnspecified },
	stringTextUnit, isSpecifiedTextUnit, isUnspecifiedTextUnit,
	takeOrElseTextUnit, mergeTextUnit, sameTextUnit, semanticEqualTextUnit, equalTextUnit, copyTextUnit,
)
```

> [!NOTE]
> **Merge and Copy are implementation-dependent.** The packed value may contain multiple
> logical components. Choose the pattern that matches your type's semantics:
>
> **Option A: Whole-value replacement** (atomic types like `TextUnit`)
> ```go
> func mergeTextUnit(a, b TextUnit) TextUnit {
>     if isSpecifiedTextUnit(b) { return b }
>     return a
> }
> ```
>
> **Option B: Component overlay** (struct option types like `Color`, see Section 3.1)
> ```go
> func copyColor(c Color, opts ...func(*Color)) Color { /* overlay specified fields */ }
> ```

**Key differences from 2-A/2-B**:
- Uses `struct { packed T }` instead of `type T primitive`
- Constructors are **mandatory** — no implicit conversion from primitives
- Sentinel is a struct value, not a primitive constant

**Key differences from 2-C**:
- Passed by **value**, not pointer → still a **ScalarCompanion**
- Sentinel is a **value**, not a pointer singleton
- `Coalesce` is not needed (no nil possible)

**Guarantees**: Same as 2-A — lives in registers/stack → **zero heap bytes**.

---

## 2.1 The `Copy` Pattern (Functional Options)

Use **functional options** backed by sentinel values to implement strictly-typed partial
updates. This resolves the ambiguity between "set to zero" and "no change":
`Shadow.Copy(WithBlurRadius(0))` vs `Shadow.Copy()`.

For a **StructCompanion**, `Copy` receives the options as `...func(T)`:

```go
type Shadow struct {
	Color      Color
	Offset     Offset
	BlurRadius float32
}

// Options mutate a scratch copy; unspecified fields fall back to the source.
func copyShadow(s *Shadow, options ...func(*Shadow)) *Shadow {
	opt := *shadowUnspecified // start from the sentinel values
	for _, option := range options {
		option(&opt)
	}
	return &Shadow{
		Color:      opt.Color.TakeOrElse(s.Color),
		Offset:     opt.Offset.TakeOrElse(s.Offset),
		BlurRadius: floatutils.Float32Companion.TakeOrElse(opt.BlurRadius, s.BlurRadius),
	}
}
```

Register `copyShadow` as the `Copy` func of `ShadowCompanion`; call it as
`ShadowCompanion.Copy(s, func(o *Shadow) { o.BlurRadius = 0 })`.

---

## 3. Semantic Levels of "Unspecified"

| Level | Value | Meaning | Typical Use |
|-------|-------|---------|-------------|
| **Object absent** | `nil` | "No style provided" | parameter default → coalesced to unspecified |
| **Object present, all fields deferred** | `textStyleUnspecified` (unexported) | "Use theme for everything" | partial merge base |
| **Field absent** | `dpUnspecified` (inside struct, unexported) | "Use ambient for this field" | field-level override |

Uppercase spellings (`TextStyleUnspecified`) name the concept here, not exported identifiers —
the only public way to reach any sentinel is `XCompanion.Unspecified()`.

**Example** (callers pass `nil` or a real style; they reach the singleton through the companion,
never through a constant):
```go
Text("hi")                                                      // nil → use theme 100 %
Text("hi", WithTextStyle(TextStyleCompanion.Unspecified()))     // singleton → use theme 100 %
Text("hi", WithTextStyle(&TextStyle{fontSize: 20}))             // partial → theme color + 20 sp
```

---

## 3.1 Complex Types — Full `TextStyle` Example (StructCompanion)

Complex types are pointers. The companion wraps a **singleton sentinel** and every helper is
nil-safe. Business code **never writes `== nil`** outside these helpers.

```go
type TextStyle struct {
	color      Color
	fontSize   Dp
	fontWeight FontWeight
}

// 1. Sentinel — a singleton, unexported (reach it via TextStyleCompanion.Unspecified())
var textStyleUnspecified = &TextStyle{
	color:      colorUnspecified,
	fontSize:   dpUnspecified,
	fontWeight: fontWeightUnspecified,
}

// 2. IsSpecified
func isSpecifiedTextStyle(s *TextStyle) bool {
	return s != nil && s != textStyleUnspecified
}

// 3. IsUnspecified
func isUnspecifiedTextStyle(s *TextStyle) bool { return !isSpecifiedTextStyle(s) }

// 4. TakeOrElse
func takeOrElseTextStyle(s, def *TextStyle) *TextStyle {
	if !isSpecifiedTextStyle(s) {
		return def
	}
	return s
}

// 5. Merge — nil-safe, field-by-field
func mergeTextStyle(a, b *TextStyle) *TextStyle {
	a = coalesceTextStyle(a, textStyleUnspecified)
	b = coalesceTextStyle(b, textStyleUnspecified)

	if a == textStyleUnspecified { return b }
	if b == textStyleUnspecified { return a }

	return &TextStyle{
		color:      mergeColor(b.color, a.color),
		fontSize:   mergeDp(b.fontSize, a.fontSize),
		fontWeight: mergeFontWeight(b.fontWeight, a.fontWeight),
	}
}

// 6. String
func stringTextStyle(s *TextStyle) string {
	if !isSpecifiedTextStyle(s) {
		return "TextStyle{Unspecified}"
	}
	return fmt.Sprintf(
		"TextStyle{color: %s, fontSize: %s, fontWeight: %s}",
		s.color, s.fontSize, s.fontWeight,
	)
}

// 7. Same — identity
func sameTextStyle(a, b *TextStyle) bool {
	if a == nil && b == nil { return true }
	if a == nil { return b == textStyleUnspecified }
	if b == nil { return a == textStyleUnspecified }
	return a == b
}

// 8. SemanticEqual — nil-safe, field-by-field
func semanticEqualTextStyle(a, b *TextStyle) bool {
	a = coalesceTextStyle(a, textStyleUnspecified)
	b = coalesceTextStyle(b, textStyleUnspecified)
	return semanticEqualColor(a.color, b.color) &&
		semanticEqualDp(a.fontSize, b.fontSize) &&
		semanticEqualFontWeight(a.fontWeight, b.fontWeight)
}

// 9. Equal — Same fast-path, else SemanticEqual
func equalTextStyle(a, b *TextStyle) bool {
	if sameTextStyle(a, b) { return true }
	return semanticEqualTextStyle(a, b)
}

// 10. Coalesce — nil is a valid input, coalesced to the singleton
func coalesceTextStyle(ptr, def *TextStyle) *TextStyle {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy — functional options, nil-safe
func copyTextStyle(s *TextStyle, options ...func(*TextStyle)) *TextStyle {
	if !isSpecifiedTextStyle(s) {
		return textStyleUnspecified
	}
	opt := *textStyleUnspecified
	for _, option := range options {
		option(&opt)
	}
	return &TextStyle{
		color:      takeOrElseColor(opt.color, s.color),
		fontSize:   takeOrElseDp(opt.fontSize, s.fontSize),
		fontWeight: takeOrElseFontWeight(opt.fontWeight, s.fontWeight),
	}
}

// The only export: a StructCompanion.
var TextStyleCompanion = companions.NewStructCompanion(
	func() *TextStyle { return textStyleUnspecified },
	stringTextStyle,
	isSpecifiedTextStyle,
	isUnspecifiedTextStyle,
	takeOrElseTextStyle,
	mergeTextStyle,
	sameTextStyle,
	semanticEqualTextStyle,
	equalTextStyle,
	coalesceTextStyle,
	copyTextStyle,
)
```

**Memory layout**: pointer on stack (8 B) → singleton in `.data` → **zero heap bytes at call
site**.

---

## 4. Sentinel Patterns for `int`, `string`, `bool`

When the type cannot hold an actual `NaN`, reserve a **valid bit-pattern** that is
**semantically impossible** in your domain.

### 4-A `int` (or `int32`, `int64`, `uint`, ...) — see `intutils`

```go
// sentinel/intutils/ints.go
type IntValue = int
const intValueUnspecified IntValue = math.MinInt   // 0x800...0 (unexported)

// The ten funcs are private; only the companion is exported.
var IntValueCompanion = companions.NewScalarCompanion(/* the ten funcs */)
```

Use it as `intutils.IntValueCompanion.IsSpecified(v)`,
`intutils.IntValueCompanion.Merge(a, b)`, etc. `0` is valid → `math.MinInt` is the sentinel.
Zero-allocation: the constant is a register literal.

### 4-B `string` (`stringutils`)

```go
const stringValueUnspecified StringValue = "\x00unspecified"  // impossible in user data
```

If **empty string is never meaningful**, you may instead alias the zero value:
`const sUnspecified = ""`. Both are **zero-allocation**: the constant lives in `.rodata`, the
field stores a pointer to that data. `stringutils` chooses `"\x00unspecified"` precisely so
that `""` stays meaningful.

### 4-C `bool` (`boolutils`, Pattern 2-D)

Both `true` and `false` are valid, so a bool sentinel is impossible. Promote to a tri-state
enum inside a zero-allocation struct wrapper:

```go
type BooleanValue struct{ value booleanTriValue } // booleanTriValue is an int enum

const (
	booleanTriValueUnspecified booleanTriValue = iota
	booleanTriValueTrue
	booleanTriValueFalse
)

var booleanValueUnspecified = BooleanValue{value: booleanTriValueUnspecified} // unexported
func BooleanValueFrom(b bool) BooleanValue { /* -> True/False */ }
```

`boolutils` registers a **ScalarCompanion** (`BooleanValueCompanion`) over the ten funcs,
and *additionally* keeps a few domain-specific **methods** on the value type
(`IsTrue`, `IsFalse`, `Bool`, `BoolOrElse`, `Format`) because that is the reason the wrapper
exists. Canonical operations go through the companion:
`boolutils.BooleanValueCompanion.TakeOrElse(a, b)`, `.Merge(a, b)`, `.Copy(v)`.

---

## 5. Copy-Paste Templates

```go
// Scalar companion over a primitive
type MyUnit float32
var myUnitUnspecified = MyUnit(floatutils.Float32Companion.Unspecified())

func isSpecifiedMyUnit(u MyUnit) bool   { return floatutils.Float32Companion.IsSpecified(float32(u)) }
func isUnspecifiedMyUnit(u MyUnit) bool { return floatutils.Float32Companion.IsUnspecified(float32(u)) }
func takeOrElseMyUnit(u, def MyUnit) MyUnit {
	if !isSpecifiedMyUnit(u) { return def }
	return u
}
func mergeMyUnit(a, b MyUnit) MyUnit {
	if isSpecifiedMyUnit(b) { return b }
	return a
}
func stringMyUnit(u MyUnit) string { return fmt.Sprintf("MyUnit{%v}", float32(u)) }
func sameMyUnit(a, b MyUnit) bool  { return float32(a) == float32(b) }
func semanticEqualMyUnit(a, b MyUnit) bool { return sameMyUnit(a, b) }
func equalMyUnit(a, b MyUnit) bool  { return sameMyUnit(a, b) }
func copyMyUnit(u MyUnit) MyUnit    { return u }

// The ONLY export:
var MyUnitCompanion = companions.NewScalarCompanion(
	func() MyUnit { return myUnitUnspecified },
	stringMyUnit, isSpecifiedMyUnit, isUnspecifiedMyUnit,
	takeOrElseMyUnit, mergeMyUnit, sameMyUnit, semanticEqualMyUnit, equalMyUnit, copyMyUnit,
)

// Struct companion over a pointer type — sentinel unexported
type MyStyle struct { field1 Color; field2 Dp }
var myStyleUnspecified = &MyStyle{field1: colorUnspecified, field2: dpUnspecified}

func isSpecifiedMyStyle(s *MyStyle) bool { return s != nil && s != myStyleUnspecified }
func coalesceMyStyle(ptr, def *MyStyle) *MyStyle {
	if ptr == nil { return def }
	return ptr
}
// ...remaining funcs...
// The ONLY export:
var MyStyleCompanion = companions.NewStructCompanion(/* 11 funcs */)
```

---

## 5.1 Protobuf Wrappers (Struct Companions)

`sentinel/protobufwrapper` wraps the nine `wrapperspb.*Value` pointer types. Each is a
**StructCompanion** and its sentinel is the private empty singleton; nil is coalesced to it:

```go
c := protobufwrapper.StringValueCompanion

c.IsSpecified(nil)                     // false → coalesced to Unspecified
c.Coalesce(nil, c.Unspecified())       // the singleton
c.Merge(nil, &wrapperspb.StringValue{Value: "x"}).Value // "x"
c.Copy(&wrapperspb.StringValue{Value: "x"}, func(v *wrapperspb.StringValue) {
    v.Value = "y"
}).Value                                              // "y"  (deep copy for bytes)
```

---

## 6. Performance Contract

- **Scalar sentinel**: 0 heap bytes, 2 ns, inlined by compiler
- **Struct nil**: 0 heap bytes, 5 ns
- **Struct custom**: 24 B struct alloc, 20 ns (paid only when the user creates a new style)

The companion is a package-level interface value; its methods are indirect calls. Keep the
**functions** (not the interface value) in hot paths when inlining matters.

Verify with:
```bash
go test -bench=. -gcflags="-m" 2>&1 | grep -E "does not escape|inlined"
```

---

## 7. Anti-Patterns (reject on sight)

```go
❌ type Color interface { IsSpecified() bool }        // interface forces alloc
❌ var UnspecifiedColor *Color = nil                  // nil interface → panic
❌ func (ts *TextStyle) IsSpecified() bool            // method on nil receiver
❌ func (ts *TextStyle) TakeOrElse(def *TextStyle) *TextStyle // REJECT: method on nil receiver
❌ func TakeOrElseColor(block func() Color) Color     // lambda escapes in Go
❌ type T struct { isSpecified bool }                 // flag field = double mem
❌ exporting the ten funcs                            // only the companion may be exported
```

---

## 8. Function Signatures (Principle of Least Parameters)

### Companion surface (2-param fallback)
```go
// Exposed through the companion — clear, no lambda, no 3rd param.
TextStyleCompanion.TakeOrElse(style, defaultStyle) // -> *TextStyle
```

### Private 3-param generic helper
```go
// takeOrElse: generic 3-param helper used INSIDE helpers to merge fields.
func takeOrElse[T comparable](a, b, sentinel T) T {
	if a != sentinel {
		return a
	}
	return b
}
```

**Rule**: the companion always exposes a **2-parameter** `TakeOrElse`; the 3-param generic is
an **implementation detail**.

---

## 9. Package-Level Contract (Machine-Readable)

```go
// COMPANION_PACKAGE_CONTRACT
// Every sentinel helper package MUST export exactly ONE public value: its Companion.
//
//   var XCompanion companions.ScalarCompanion[T]  // for value types
//   var XCompanion companions.StructCompanion[T]  // for pointer types
//
// The companion MUST be constructed from these ten package-private funcs:
//   func xUnspecified() T                     // sentinel / singleton
//   func stringX(T) string                    // stringification
//   func isSpecifiedX(T) bool                 // predicate
//   func isUnspecifiedX(T) bool               // inverse predicate
//   func takeOrElseX(src, def T) T            // fallback
//   func mergeX(a, b T) T                     // composition
//   func sameX(a, b T) bool                   // identity
//   func semanticEqualX(a, b T) bool          // semantic equality
//   func equalX(a, b T) bool                  // composed equality
//   func copyX(...) ...                       // scalar: Copy(T) T
//                                              // struct: Copy(T, ...func(T)) T
// Plus, for StructCompanion only:
//   func coalesceX(ptr, def T) T              // nil coalescing
//
// Additional symbols are allowed but MUST be documented and package-private.
// All sentinels MUST be compile-time constants or package-level variables.
// No function may accept func() parameters in hot paths (forces heap escape).
// No method may have a pointer receiver that checks for nil (anti-pattern).
// END_CONTRACT
```

---

## 10. Verification Commands

```bash
# Check for heap escapes (must show "does not escape")
go test -gcflags="-m" -run=ExampleUsage 2>&1 | grep -E "does not escape|escapes to heap"

# Check for inlining (must show "can inline" for hot paths)
go test -gcflags="-m" -run=ExampleUsage 2>&1 | grep "can inline"

# Benchmark allocation counts (should be 0 or 1 per operation)
go test -bench=. -benchmem

# Build with escape analysis for the entire module
go build -gcflags="-m" ./... 2>&1 | grep "sentinel/"
```

---

**Document Version**: 3.0  
**Last Updated**: 2026-10-09
