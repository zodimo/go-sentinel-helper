package protobufwrapper

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// Int64Value

// 1. Sentinel
var int64ValueUnspecified = &wrapperspb.Int64Value{}

// 2. IsSpecified
func isSpecifiedInt64Value(v *wrapperspb.Int64Value) bool {
	return v != nil && v != int64ValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedInt64Value(v *wrapperspb.Int64Value) bool {
	return !isSpecifiedInt64Value(v)
}

// 4. TakeOrElse
func takeOrElseInt64Value(v, def *wrapperspb.Int64Value) *wrapperspb.Int64Value {
	if !isSpecifiedInt64Value(v) {
		return def
	}
	return v
}

// 5. Merge
func mergeInt64Value(a, b *wrapperspb.Int64Value) *wrapperspb.Int64Value {
	a = coalesceInt64Value(a, int64ValueUnspecified)
	b = coalesceInt64Value(b, int64ValueUnspecified)

	if a == int64ValueUnspecified {
		return b
	}
	if b == int64ValueUnspecified {
		return a
	}

	return &wrapperspb.Int64Value{Value: b.Value}
}

// 6. String
func stringInt64Value(v *wrapperspb.Int64Value) string {
	if !isSpecifiedInt64Value(v) {
		return "Int64Value{Unspecified}"
	}
	return fmt.Sprintf("Int64Value{%d}", v.Value)
}

// 7. Same
func sameInt64Value(a, b *wrapperspb.Int64Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == int64ValueUnspecified
	}
	if b == nil {
		return a == int64ValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func semanticEqualInt64Value(a, b *wrapperspb.Int64Value) bool {
	a = coalesceInt64Value(a, int64ValueUnspecified)
	b = coalesceInt64Value(b, int64ValueUnspecified)
	return a.Value == b.Value
}

// 9. Equal
func equalInt64Value(a, b *wrapperspb.Int64Value) bool {
	if sameInt64Value(a, b) {
		return true
	}
	return semanticEqualInt64Value(a, b)
}

// 10. Coalesce
func coalesceInt64Value(ptr, def *wrapperspb.Int64Value) *wrapperspb.Int64Value {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy
func copyInt64Value(v *wrapperspb.Int64Value, options ...func(*wrapperspb.Int64Value)) *wrapperspb.Int64Value {
	if !isSpecifiedInt64Value(v) {
		return int64ValueUnspecified
	}
	out := &wrapperspb.Int64Value{Value: v.Value}
	for _, option := range options {
		option(out)
	}
	return out
}

// 12. Unspecified
func unspecifiedInt64Value() *wrapperspb.Int64Value {
	return int64ValueUnspecified
}

// Int64ValueCompanion is the ONLY public companion of this type.
var Int64ValueCompanion = companions.NewStructCompanion(
	unspecifiedInt64Value,
	stringInt64Value,
	isSpecifiedInt64Value,
	isUnspecifiedInt64Value,
	takeOrElseInt64Value,
	mergeInt64Value,
	sameInt64Value,
	semanticEqualInt64Value,
	equalInt64Value,
	coalesceInt64Value,
	copyInt64Value,
)
