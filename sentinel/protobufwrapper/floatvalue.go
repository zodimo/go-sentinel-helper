package protobufwrapper

import (
	"fmt"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// FloatValue

// 1. Sentinel
var FloatValueUnspecified = &wrapperspb.FloatValue{}

// 2. IsSpecified
func IsSpecifiedFloatValue(v *wrapperspb.FloatValue) bool {
	return v != nil && v != FloatValueUnspecified
}

// 3. TakeOrElse
func TakeOrElseFloatValue(v, def *wrapperspb.FloatValue) *wrapperspb.FloatValue {
	if v == nil || v == FloatValueUnspecified {
		return def
	}
	return v
}

// 4. Merge
func MergeFloatValue(a, b *wrapperspb.FloatValue) *wrapperspb.FloatValue {
	a = CoalesceFloatValue(a, FloatValueUnspecified)
	b = CoalesceFloatValue(b, FloatValueUnspecified)

	if a == FloatValueUnspecified {
		return b
	}
	if b == FloatValueUnspecified {
		return a
	}

	return &wrapperspb.FloatValue{Value: b.Value}
}

// 5. String
func StringFloatValue(v *wrapperspb.FloatValue) string {
	if !IsSpecifiedFloatValue(v) {
		return "FloatValue{Unspecified}"
	}
	return fmt.Sprintf("FloatValue{%g}", v.Value)
}

// 6. Coalesce
func CoalesceFloatValue(ptr, def *wrapperspb.FloatValue) *wrapperspb.FloatValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 7. Same
func SameFloatValue(a, b *wrapperspb.FloatValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == FloatValueUnspecified
	}
	if b == nil {
		return a == FloatValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func SemanticEqualFloatValue(a, b *wrapperspb.FloatValue) bool {
	a = CoalesceFloatValue(a, FloatValueUnspecified)
	b = CoalesceFloatValue(b, FloatValueUnspecified)

	return a.Value == b.Value
}

// 9. Equal
func EqualFloatValue(a, b *wrapperspb.FloatValue) bool {
	if !SameFloatValue(a, b) {
		return SemanticEqualFloatValue(a, b)
	}
	return true
}

// 10. Copy
func CopyFloatValue(v *wrapperspb.FloatValue) *wrapperspb.FloatValue {
	if !IsSpecifiedFloatValue(v) {
		return FloatValueUnspecified
	}
	return &wrapperspb.FloatValue{Value: v.Value}
}
