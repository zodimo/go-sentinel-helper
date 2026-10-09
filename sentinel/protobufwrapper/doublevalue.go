package protobufwrapper

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// DoubleValue

// 1. Sentinel
var doubleValueUnspecified = &wrapperspb.DoubleValue{}

// 2. IsSpecified
func isSpecifiedDoubleValue(v *wrapperspb.DoubleValue) bool {
	return v != nil && v != doubleValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedDoubleValue(v *wrapperspb.DoubleValue) bool {
	return !isSpecifiedDoubleValue(v)
}

// 4. TakeOrElse
func takeOrElseDoubleValue(v, def *wrapperspb.DoubleValue) *wrapperspb.DoubleValue {
	if !isSpecifiedDoubleValue(v) {
		return def
	}
	return v
}

// 5. Merge
func mergeDoubleValue(a, b *wrapperspb.DoubleValue) *wrapperspb.DoubleValue {
	a = coalesceDoubleValue(a, doubleValueUnspecified)
	b = coalesceDoubleValue(b, doubleValueUnspecified)

	if a == doubleValueUnspecified {
		return b
	}
	if b == doubleValueUnspecified {
		return a
	}

	return &wrapperspb.DoubleValue{Value: b.Value}
}

// 6. String
func stringDoubleValue(v *wrapperspb.DoubleValue) string {
	if !isSpecifiedDoubleValue(v) {
		return "DoubleValue{Unspecified}"
	}
	return fmt.Sprintf("DoubleValue{%g}", v.Value)
}

// 7. Same
func sameDoubleValue(a, b *wrapperspb.DoubleValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == doubleValueUnspecified
	}
	if b == nil {
		return a == doubleValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func semanticEqualDoubleValue(a, b *wrapperspb.DoubleValue) bool {
	a = coalesceDoubleValue(a, doubleValueUnspecified)
	b = coalesceDoubleValue(b, doubleValueUnspecified)
	return a.Value == b.Value
}

// 9. Equal
func equalDoubleValue(a, b *wrapperspb.DoubleValue) bool {
	if sameDoubleValue(a, b) {
		return true
	}
	return semanticEqualDoubleValue(a, b)
}

// 10. Coalesce
func coalesceDoubleValue(ptr, def *wrapperspb.DoubleValue) *wrapperspb.DoubleValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy
func copyDoubleValue(v *wrapperspb.DoubleValue, options ...func(*wrapperspb.DoubleValue)) *wrapperspb.DoubleValue {
	if !isSpecifiedDoubleValue(v) {
		return doubleValueUnspecified
	}
	out := &wrapperspb.DoubleValue{Value: v.Value}
	for _, option := range options {
		option(out)
	}
	return out
}

// 12. Unspecified
func unspecifiedDoubleValue() *wrapperspb.DoubleValue {
	return doubleValueUnspecified
}

// DoubleValueCompanion is the ONLY public companion of this type.
var DoubleValueCompanion = companions.NewStructCompanion(
	unspecifiedDoubleValue,
	stringDoubleValue,
	isSpecifiedDoubleValue,
	isUnspecifiedDoubleValue,
	takeOrElseDoubleValue,
	mergeDoubleValue,
	sameDoubleValue,
	semanticEqualDoubleValue,
	equalDoubleValue,
	coalesceDoubleValue,
	copyDoubleValue,
)
