package protobufwrapper

import (
	"fmt"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// BoolValue

// 1. Sentinel
var BoolValueUnspecified = &wrapperspb.BoolValue{}

// 2. IsSpecified
func IsSpecifiedBoolValue(v *wrapperspb.BoolValue) bool {
	return v != nil && v != BoolValueUnspecified
}

// 3. TakeOrElse
func TakeOrElseBoolValue(v, def *wrapperspb.BoolValue) *wrapperspb.BoolValue {
	if v == nil || v == BoolValueUnspecified {
		return def
	}
	return v
}

// 4. Merge
func MergeBoolValue(a, b *wrapperspb.BoolValue) *wrapperspb.BoolValue {
	a = CoalesceBoolValue(a, BoolValueUnspecified)
	b = CoalesceBoolValue(b, BoolValueUnspecified)

	if a == BoolValueUnspecified {
		return b
	}
	if b == BoolValueUnspecified {
		return a
	}

	return &wrapperspb.BoolValue{Value: b.Value}
}

// 5. String
func StringBoolValue(v *wrapperspb.BoolValue) string {
	if !IsSpecifiedBoolValue(v) {
		return "BoolValue{Unspecified}"
	}
	return fmt.Sprintf("BoolValue{%t}", v.Value)
}

// 6. Coalesce
func CoalesceBoolValue(ptr, def *wrapperspb.BoolValue) *wrapperspb.BoolValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 7. Same
func SameBoolValue(a, b *wrapperspb.BoolValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == BoolValueUnspecified
	}
	if b == nil {
		return a == BoolValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func SemanticEqualBoolValue(a, b *wrapperspb.BoolValue) bool {
	a = CoalesceBoolValue(a, BoolValueUnspecified)
	b = CoalesceBoolValue(b, BoolValueUnspecified)

	return a.Value == b.Value
}

// 9. Equal
func EqualBoolValue(a, b *wrapperspb.BoolValue) bool {
	if !SameBoolValue(a, b) {
		return SemanticEqualBoolValue(a, b)
	}
	return true
}

// 10. Copy
func CopyBoolValue(v *wrapperspb.BoolValue) *wrapperspb.BoolValue {
	if !IsSpecifiedBoolValue(v) {
		return BoolValueUnspecified
	}
	return &wrapperspb.BoolValue{Value: v.Value}
}
