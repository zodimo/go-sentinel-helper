package protobufwrapper

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// FloatValue

// 1. Sentinel
var floatValueUnspecified = &wrapperspb.FloatValue{}

// 2. IsSpecified
func isSpecifiedFloatValue(v *wrapperspb.FloatValue) bool {
	return v != nil && v != floatValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedFloatValue(v *wrapperspb.FloatValue) bool {
	return !isSpecifiedFloatValue(v)
}

// 4. TakeOrElse
func takeOrElseFloatValue(v, def *wrapperspb.FloatValue) *wrapperspb.FloatValue {
	if !isSpecifiedFloatValue(v) {
		return def
	}
	return v
}

// 5. Merge
func mergeFloatValue(a, b *wrapperspb.FloatValue) *wrapperspb.FloatValue {
	a = coalesceFloatValue(a, floatValueUnspecified)
	b = coalesceFloatValue(b, floatValueUnspecified)

	if a == floatValueUnspecified {
		return b
	}
	if b == floatValueUnspecified {
		return a
	}

	return &wrapperspb.FloatValue{Value: b.Value}
}

// 6. String
func stringFloatValue(v *wrapperspb.FloatValue) string {
	if !isSpecifiedFloatValue(v) {
		return "FloatValue{Unspecified}"
	}
	return fmt.Sprintf("FloatValue{%g}", v.Value)
}

// 7. Same
func sameFloatValue(a, b *wrapperspb.FloatValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == floatValueUnspecified
	}
	if b == nil {
		return a == floatValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func semanticEqualFloatValue(a, b *wrapperspb.FloatValue) bool {
	a = coalesceFloatValue(a, floatValueUnspecified)
	b = coalesceFloatValue(b, floatValueUnspecified)
	return a.Value == b.Value
}

// 9. Equal
func equalFloatValue(a, b *wrapperspb.FloatValue) bool {
	if sameFloatValue(a, b) {
		return true
	}
	return semanticEqualFloatValue(a, b)
}

// 10. Coalesce
func coalesceFloatValue(ptr, def *wrapperspb.FloatValue) *wrapperspb.FloatValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy
func copyFloatValue(v *wrapperspb.FloatValue, options ...func(*wrapperspb.FloatValue)) *wrapperspb.FloatValue {
	if !isSpecifiedFloatValue(v) {
		return floatValueUnspecified
	}
	out := &wrapperspb.FloatValue{Value: v.Value}
	for _, option := range options {
		option(out)
	}
	return out
}

// 12. Unspecified
func unspecifiedFloatValue() *wrapperspb.FloatValue {
	return floatValueUnspecified
}

// FloatValueCompanion is the ONLY public companion of this type.
var FloatValueCompanion = companions.NewStructCompanion(
	unspecifiedFloatValue,
	stringFloatValue,
	isSpecifiedFloatValue,
	isUnspecifiedFloatValue,
	takeOrElseFloatValue,
	mergeFloatValue,
	sameFloatValue,
	semanticEqualFloatValue,
	equalFloatValue,
	coalesceFloatValue,
	copyFloatValue,
)
