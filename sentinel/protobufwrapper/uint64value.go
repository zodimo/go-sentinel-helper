package protobufwrapper

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// UInt64Value

// 1. Sentinel
var uint64ValueUnspecified = &wrapperspb.UInt64Value{}

// 2. IsSpecified
func isSpecifiedUInt64Value(v *wrapperspb.UInt64Value) bool {
	return v != nil && v != uint64ValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedUInt64Value(v *wrapperspb.UInt64Value) bool {
	return !isSpecifiedUInt64Value(v)
}

// 4. TakeOrElse
func takeOrElseUInt64Value(v, def *wrapperspb.UInt64Value) *wrapperspb.UInt64Value {
	if !isSpecifiedUInt64Value(v) {
		return def
	}
	return v
}

// 5. Merge
func mergeUInt64Value(a, b *wrapperspb.UInt64Value) *wrapperspb.UInt64Value {
	a = coalesceUInt64Value(a, uint64ValueUnspecified)
	b = coalesceUInt64Value(b, uint64ValueUnspecified)

	if a == uint64ValueUnspecified {
		return b
	}
	if b == uint64ValueUnspecified {
		return a
	}

	return &wrapperspb.UInt64Value{Value: b.Value}
}

// 6. String
func stringUInt64Value(v *wrapperspb.UInt64Value) string {
	if !isSpecifiedUInt64Value(v) {
		return "UInt64Value{Unspecified}"
	}
	return fmt.Sprintf("UInt64Value{%d}", v.Value)
}

// 7. Same
func sameUInt64Value(a, b *wrapperspb.UInt64Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == uint64ValueUnspecified
	}
	if b == nil {
		return a == uint64ValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func semanticEqualUInt64Value(a, b *wrapperspb.UInt64Value) bool {
	a = coalesceUInt64Value(a, uint64ValueUnspecified)
	b = coalesceUInt64Value(b, uint64ValueUnspecified)
	return a.Value == b.Value
}

// 9. Equal
func equalUInt64Value(a, b *wrapperspb.UInt64Value) bool {
	if sameUInt64Value(a, b) {
		return true
	}
	return semanticEqualUInt64Value(a, b)
}

// 10. Coalesce
func coalesceUInt64Value(ptr, def *wrapperspb.UInt64Value) *wrapperspb.UInt64Value {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy
func copyUInt64Value(v *wrapperspb.UInt64Value, options ...func(*wrapperspb.UInt64Value)) *wrapperspb.UInt64Value {
	if !isSpecifiedUInt64Value(v) {
		return uint64ValueUnspecified
	}
	out := &wrapperspb.UInt64Value{Value: v.Value}
	for _, option := range options {
		option(out)
	}
	return out
}

// 12. Unspecified
func unspecifiedUInt64Value() *wrapperspb.UInt64Value {
	return uint64ValueUnspecified
}

// UInt64ValueCompanion is the ONLY public companion of this type.
var UInt64ValueCompanion = companions.NewStructCompanion(
	unspecifiedUInt64Value,
	stringUInt64Value,
	isSpecifiedUInt64Value,
	isUnspecifiedUInt64Value,
	takeOrElseUInt64Value,
	mergeUInt64Value,
	sameUInt64Value,
	semanticEqualUInt64Value,
	equalUInt64Value,
	coalesceUInt64Value,
	copyUInt64Value,
)
