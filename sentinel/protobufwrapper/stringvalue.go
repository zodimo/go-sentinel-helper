package protobufwrapper

import (
	"fmt"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// StringValue

// 1. Sentinel
var StringValueUnspecified = &wrapperspb.StringValue{}

// 2. IsSpecified
func IsSpecifiedStringValue(v *wrapperspb.StringValue) bool {
	return v != nil && v != StringValueUnspecified
}

// 3. TakeOrElse
func TakeOrElseStringValue(v, def *wrapperspb.StringValue) *wrapperspb.StringValue {
	if v == nil || v == StringValueUnspecified {
		return def
	}
	return v
}

// 4. Merge
func MergeStringValue(a, b *wrapperspb.StringValue) *wrapperspb.StringValue {
	a = CoalesceStringValue(a, StringValueUnspecified)
	b = CoalesceStringValue(b, StringValueUnspecified)

	if a == StringValueUnspecified {
		return b
	}
	if b == StringValueUnspecified {
		return a
	}

	return &wrapperspb.StringValue{Value: b.Value}
}

// 5. String
func StringStringValue(v *wrapperspb.StringValue) string {
	if !IsSpecifiedStringValue(v) {
		return "StringValue{Unspecified}"
	}
	return fmt.Sprintf("StringValue{%q}", v.Value)
}

// 6. Coalesce
func CoalesceStringValue(ptr, def *wrapperspb.StringValue) *wrapperspb.StringValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 7. Same
func SameStringValue(a, b *wrapperspb.StringValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == StringValueUnspecified
	}
	if b == nil {
		return a == StringValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func SemanticEqualStringValue(a, b *wrapperspb.StringValue) bool {
	a = CoalesceStringValue(a, StringValueUnspecified)
	b = CoalesceStringValue(b, StringValueUnspecified)

	return a.Value == b.Value
}

// 9. Equal
func EqualStringValue(a, b *wrapperspb.StringValue) bool {
	if !SameStringValue(a, b) {
		return SemanticEqualStringValue(a, b)
	}
	return true
}

// 10. Copy
func CopyStringValue(v *wrapperspb.StringValue) *wrapperspb.StringValue {
	if !IsSpecifiedStringValue(v) {
		return StringValueUnspecified
	}
	return &wrapperspb.StringValue{Value: v.Value}
}
