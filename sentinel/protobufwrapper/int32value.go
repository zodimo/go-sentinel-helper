package protobufwrapper

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// Int32Value

// 1. Sentinel
var int32ValueUnspecified = &wrapperspb.Int32Value{}

// 2. IsSpecified
func isSpecifiedInt32Value(v *wrapperspb.Int32Value) bool {
	return v != nil && v != int32ValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedInt32Value(v *wrapperspb.Int32Value) bool {
	return !isSpecifiedInt32Value(v)
}

// 4. TakeOrElse
func takeOrElseInt32Value(v, def *wrapperspb.Int32Value) *wrapperspb.Int32Value {
	if !isSpecifiedInt32Value(v) {
		return def
	}
	return v
}

// 5. Merge
func mergeInt32Value(a, b *wrapperspb.Int32Value) *wrapperspb.Int32Value {
	a = coalesceInt32Value(a, int32ValueUnspecified)
	b = coalesceInt32Value(b, int32ValueUnspecified)

	if a == int32ValueUnspecified {
		return b
	}
	if b == int32ValueUnspecified {
		return a
	}

	return &wrapperspb.Int32Value{Value: b.Value}
}

// 6. String
func stringInt32Value(v *wrapperspb.Int32Value) string {
	if !isSpecifiedInt32Value(v) {
		return "Int32Value{Unspecified}"
	}
	return fmt.Sprintf("Int32Value{%d}", v.Value)
}

// 7. Same
func sameInt32Value(a, b *wrapperspb.Int32Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == int32ValueUnspecified
	}
	if b == nil {
		return a == int32ValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func semanticEqualInt32Value(a, b *wrapperspb.Int32Value) bool {
	a = coalesceInt32Value(a, int32ValueUnspecified)
	b = coalesceInt32Value(b, int32ValueUnspecified)
	return a.Value == b.Value
}

// 9. Equal
func equalInt32Value(a, b *wrapperspb.Int32Value) bool {
	if sameInt32Value(a, b) {
		return true
	}
	return semanticEqualInt32Value(a, b)
}

// 10. Coalesce
func coalesceInt32Value(ptr, def *wrapperspb.Int32Value) *wrapperspb.Int32Value {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy
func copyInt32Value(v *wrapperspb.Int32Value, options ...func(*wrapperspb.Int32Value)) *wrapperspb.Int32Value {
	if !isSpecifiedInt32Value(v) {
		return int32ValueUnspecified
	}
	out := &wrapperspb.Int32Value{Value: v.Value}
	for _, option := range options {
		option(out)
	}
	return out
}

// 12. Unspecified
func unspecifiedInt32Value() *wrapperspb.Int32Value {
	return int32ValueUnspecified
}

// Int32ValueCompanion is the ONLY public companion of this type.
var Int32ValueCompanion = companions.NewStructCompanion(
	unspecifiedInt32Value,
	stringInt32Value,
	isSpecifiedInt32Value,
	isUnspecifiedInt32Value,
	takeOrElseInt32Value,
	mergeInt32Value,
	sameInt32Value,
	semanticEqualInt32Value,
	equalInt32Value,
	coalesceInt32Value,
	copyInt32Value,
)
