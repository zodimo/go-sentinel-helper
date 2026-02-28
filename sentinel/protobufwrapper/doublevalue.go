package protobufwrapper

import (
	"fmt"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// DoubleValue

// 1. Sentinel
var DoubleValueUnspecified = &wrapperspb.DoubleValue{}

// 2. IsSpecified
func IsSpecifiedDoubleValue(v *wrapperspb.DoubleValue) bool {
	return v != nil && v != DoubleValueUnspecified
}

// 3. TakeOrElse
func TakeOrElseDoubleValue(v, def *wrapperspb.DoubleValue) *wrapperspb.DoubleValue {
	if v == nil || v == DoubleValueUnspecified {
		return def
	}
	return v
}

// 4. Merge
func MergeDoubleValue(a, b *wrapperspb.DoubleValue) *wrapperspb.DoubleValue {
	a = CoalesceDoubleValue(a, DoubleValueUnspecified)
	b = CoalesceDoubleValue(b, DoubleValueUnspecified)

	if a == DoubleValueUnspecified {
		return b
	}
	if b == DoubleValueUnspecified {
		return a
	}

	return &wrapperspb.DoubleValue{Value: b.Value}
}

// 5. String
func StringDoubleValue(v *wrapperspb.DoubleValue) string {
	if !IsSpecifiedDoubleValue(v) {
		return "DoubleValue{Unspecified}"
	}
	return fmt.Sprintf("DoubleValue{%g}", v.Value)
}

// 6. Coalesce
func CoalesceDoubleValue(ptr, def *wrapperspb.DoubleValue) *wrapperspb.DoubleValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 7. Same
func SameDoubleValue(a, b *wrapperspb.DoubleValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == DoubleValueUnspecified
	}
	if b == nil {
		return a == DoubleValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func SemanticEqualDoubleValue(a, b *wrapperspb.DoubleValue) bool {
	a = CoalesceDoubleValue(a, DoubleValueUnspecified)
	b = CoalesceDoubleValue(b, DoubleValueUnspecified)

	return a.Value == b.Value
}

// 9. Equal
func EqualDoubleValue(a, b *wrapperspb.DoubleValue) bool {
	if !SameDoubleValue(a, b) {
		return SemanticEqualDoubleValue(a, b)
	}
	return true
}

// 10. Copy
func CopyDoubleValue(v *wrapperspb.DoubleValue) *wrapperspb.DoubleValue {
	if !IsSpecifiedDoubleValue(v) {
		return DoubleValueUnspecified
	}
	return &wrapperspb.DoubleValue{Value: v.Value}
}
