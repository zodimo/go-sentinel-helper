package protobufwrapper

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// UInt32Value

// 1. Sentinel
var uint32ValueUnspecified = &wrapperspb.UInt32Value{}

// 2. IsSpecified
func isSpecifiedUInt32Value(v *wrapperspb.UInt32Value) bool {
	return v != nil && v != uint32ValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedUInt32Value(v *wrapperspb.UInt32Value) bool {
	return !isSpecifiedUInt32Value(v)
}

// 4. TakeOrElse
func takeOrElseUInt32Value(v, def *wrapperspb.UInt32Value) *wrapperspb.UInt32Value {
	if !isSpecifiedUInt32Value(v) {
		return def
	}
	return v
}

// 5. Merge
func mergeUInt32Value(a, b *wrapperspb.UInt32Value) *wrapperspb.UInt32Value {
	a = coalesceUInt32Value(a, uint32ValueUnspecified)
	b = coalesceUInt32Value(b, uint32ValueUnspecified)

	if a == uint32ValueUnspecified {
		return b
	}
	if b == uint32ValueUnspecified {
		return a
	}

	return &wrapperspb.UInt32Value{Value: b.Value}
}

// 6. String
func stringUInt32Value(v *wrapperspb.UInt32Value) string {
	if !isSpecifiedUInt32Value(v) {
		return "UInt32Value{Unspecified}"
	}
	return fmt.Sprintf("UInt32Value{%d}", v.Value)
}

// 7. Same
func sameUInt32Value(a, b *wrapperspb.UInt32Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == uint32ValueUnspecified
	}
	if b == nil {
		return a == uint32ValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func semanticEqualUInt32Value(a, b *wrapperspb.UInt32Value) bool {
	a = coalesceUInt32Value(a, uint32ValueUnspecified)
	b = coalesceUInt32Value(b, uint32ValueUnspecified)
	return a.Value == b.Value
}

// 9. Equal
func equalUInt32Value(a, b *wrapperspb.UInt32Value) bool {
	if sameUInt32Value(a, b) {
		return true
	}
	return semanticEqualUInt32Value(a, b)
}

// 10. Coalesce
func coalesceUInt32Value(ptr, def *wrapperspb.UInt32Value) *wrapperspb.UInt32Value {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy
func copyUInt32Value(v *wrapperspb.UInt32Value, options ...func(*wrapperspb.UInt32Value)) *wrapperspb.UInt32Value {
	if !isSpecifiedUInt32Value(v) {
		return uint32ValueUnspecified
	}
	out := &wrapperspb.UInt32Value{Value: v.Value}
	for _, option := range options {
		option(out)
	}
	return out
}

// 12. Unspecified
func unspecifiedUInt32Value() *wrapperspb.UInt32Value {
	return uint32ValueUnspecified
}

// UInt32ValueCompanion is the ONLY public companion of this type.
var UInt32ValueCompanion = companions.NewStructCompanion(
	unspecifiedUInt32Value,
	stringUInt32Value,
	isSpecifiedUInt32Value,
	isUnspecifiedUInt32Value,
	takeOrElseUInt32Value,
	mergeUInt32Value,
	sameUInt32Value,
	semanticEqualUInt32Value,
	equalUInt32Value,
	coalesceUInt32Value,
	copyUInt32Value,
)
