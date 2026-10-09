package protobufwrapper

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// StringValue

// 1. Sentinel
var stringValueUnspecified = &wrapperspb.StringValue{}

// 2. IsSpecified
func isSpecifiedStringValue(v *wrapperspb.StringValue) bool {
	return v != nil && v != stringValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedStringValue(v *wrapperspb.StringValue) bool {
	return !isSpecifiedStringValue(v)
}

// 4. TakeOrElse
func takeOrElseStringValue(v, def *wrapperspb.StringValue) *wrapperspb.StringValue {
	if !isSpecifiedStringValue(v) {
		return def
	}
	return v
}

// 5. Merge
func mergeStringValue(a, b *wrapperspb.StringValue) *wrapperspb.StringValue {
	a = coalesceStringValue(a, stringValueUnspecified)
	b = coalesceStringValue(b, stringValueUnspecified)

	if a == stringValueUnspecified {
		return b
	}
	if b == stringValueUnspecified {
		return a
	}

	return &wrapperspb.StringValue{Value: b.Value}
}

// 6. String
func stringStringValue(v *wrapperspb.StringValue) string {
	if !isSpecifiedStringValue(v) {
		return "StringValue{Unspecified}"
	}
	return fmt.Sprintf("StringValue{%q}", v.Value)
}

// 7. Same
func sameStringValue(a, b *wrapperspb.StringValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == stringValueUnspecified
	}
	if b == nil {
		return a == stringValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func semanticEqualStringValue(a, b *wrapperspb.StringValue) bool {
	a = coalesceStringValue(a, stringValueUnspecified)
	b = coalesceStringValue(b, stringValueUnspecified)
	return a.Value == b.Value
}

// 9. Equal
func equalStringValue(a, b *wrapperspb.StringValue) bool {
	if sameStringValue(a, b) {
		return true
	}
	return semanticEqualStringValue(a, b)
}

// 10. Coalesce
func coalesceStringValue(ptr, def *wrapperspb.StringValue) *wrapperspb.StringValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy
func copyStringValue(v *wrapperspb.StringValue, options ...func(*wrapperspb.StringValue)) *wrapperspb.StringValue {
	if !isSpecifiedStringValue(v) {
		return stringValueUnspecified
	}
	out := &wrapperspb.StringValue{Value: v.Value}
	for _, option := range options {
		option(out)
	}
	return out
}

// 12. Unspecified
func unspecifiedStringValue() *wrapperspb.StringValue {
	return stringValueUnspecified
}

// StringValueCompanion is the ONLY public companion of this type.
var StringValueCompanion = companions.NewStructCompanion(
	unspecifiedStringValue,
	stringStringValue,
	isSpecifiedStringValue,
	isUnspecifiedStringValue,
	takeOrElseStringValue,
	mergeStringValue,
	sameStringValue,
	semanticEqualStringValue,
	equalStringValue,
	coalesceStringValue,
	copyStringValue,
)
