package protobufwrapper

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// BoolValue

// 1. Sentinel
var boolValueUnspecified = &wrapperspb.BoolValue{}

// 2. IsSpecified
func isSpecifiedBoolValue(v *wrapperspb.BoolValue) bool {
	return v != nil && v != boolValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedBoolValue(v *wrapperspb.BoolValue) bool {
	return !isSpecifiedBoolValue(v)
}

// 4. TakeOrElse
func takeOrElseBoolValue(v, def *wrapperspb.BoolValue) *wrapperspb.BoolValue {
	if !isSpecifiedBoolValue(v) {
		return def
	}
	return v
}

// 5. Merge
func mergeBoolValue(a, b *wrapperspb.BoolValue) *wrapperspb.BoolValue {
	a = coalesceBoolValue(a, boolValueUnspecified)
	b = coalesceBoolValue(b, boolValueUnspecified)

	if a == boolValueUnspecified {
		return b
	}
	if b == boolValueUnspecified {
		return a
	}

	return &wrapperspb.BoolValue{Value: b.Value}
}

// 6. String
func stringBoolValue(v *wrapperspb.BoolValue) string {
	if !isSpecifiedBoolValue(v) {
		return "BoolValue{Unspecified}"
	}
	return fmt.Sprintf("BoolValue{%t}", v.Value)
}

// 7. Same
func sameBoolValue(a, b *wrapperspb.BoolValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == boolValueUnspecified
	}
	if b == nil {
		return a == boolValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func semanticEqualBoolValue(a, b *wrapperspb.BoolValue) bool {
	a = coalesceBoolValue(a, boolValueUnspecified)
	b = coalesceBoolValue(b, boolValueUnspecified)
	return a.Value == b.Value
}

// 9. Equal
func equalBoolValue(a, b *wrapperspb.BoolValue) bool {
	if sameBoolValue(a, b) {
		return true
	}
	return semanticEqualBoolValue(a, b)
}

// 10. Coalesce
func coalesceBoolValue(ptr, def *wrapperspb.BoolValue) *wrapperspb.BoolValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy
func copyBoolValue(v *wrapperspb.BoolValue, options ...func(*wrapperspb.BoolValue)) *wrapperspb.BoolValue {
	if !isSpecifiedBoolValue(v) {
		return boolValueUnspecified
	}
	out := &wrapperspb.BoolValue{Value: v.Value}
	for _, option := range options {
		option(out)
	}
	return out
}

// 12. Unspecified
func unspecifiedBoolValue() *wrapperspb.BoolValue {
	return boolValueUnspecified
}

// BoolValueCompanion is the ONLY public companion of this type.
var BoolValueCompanion = companions.NewStructCompanion(
	unspecifiedBoolValue,
	stringBoolValue,
	isSpecifiedBoolValue,
	isUnspecifiedBoolValue,
	takeOrElseBoolValue,
	mergeBoolValue,
	sameBoolValue,
	semanticEqualBoolValue,
	equalBoolValue,
	coalesceBoolValue,
	copyBoolValue,
)
