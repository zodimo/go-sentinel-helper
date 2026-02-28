package protobufwrapper

import (
	"fmt"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// UInt64Value

// 1. Sentinel
var UInt64ValueUnspecified = &wrapperspb.UInt64Value{}

// 2. IsSpecified
func IsSpecifiedUInt64Value(v *wrapperspb.UInt64Value) bool {
	return v != nil && v != UInt64ValueUnspecified
}

// 3. TakeOrElse
func TakeOrElseUInt64Value(v, def *wrapperspb.UInt64Value) *wrapperspb.UInt64Value {
	if v == nil || v == UInt64ValueUnspecified {
		return def
	}
	return v
}

// 4. Merge
func MergeUInt64Value(a, b *wrapperspb.UInt64Value) *wrapperspb.UInt64Value {
	a = CoalesceUInt64Value(a, UInt64ValueUnspecified)
	b = CoalesceUInt64Value(b, UInt64ValueUnspecified)

	if a == UInt64ValueUnspecified {
		return b
	}
	if b == UInt64ValueUnspecified {
		return a
	}

	return &wrapperspb.UInt64Value{Value: b.Value}
}

// 5. String
func StringUInt64Value(v *wrapperspb.UInt64Value) string {
	if !IsSpecifiedUInt64Value(v) {
		return "UInt64Value{Unspecified}"
	}
	return fmt.Sprintf("UInt64Value{%d}", v.Value)
}

// 6. Coalesce
func CoalesceUInt64Value(ptr, def *wrapperspb.UInt64Value) *wrapperspb.UInt64Value {
	if ptr == nil {
		return def
	}
	return ptr
}

// 7. Same
func SameUInt64Value(a, b *wrapperspb.UInt64Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == UInt64ValueUnspecified
	}
	if b == nil {
		return a == UInt64ValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func SemanticEqualUInt64Value(a, b *wrapperspb.UInt64Value) bool {
	a = CoalesceUInt64Value(a, UInt64ValueUnspecified)
	b = CoalesceUInt64Value(b, UInt64ValueUnspecified)

	return a.Value == b.Value
}

// 9. Equal
func EqualUInt64Value(a, b *wrapperspb.UInt64Value) bool {
	if !SameUInt64Value(a, b) {
		return SemanticEqualUInt64Value(a, b)
	}
	return true
}

// 10. Copy
func CopyUInt64Value(v *wrapperspb.UInt64Value) *wrapperspb.UInt64Value {
	if !IsSpecifiedUInt64Value(v) {
		return UInt64ValueUnspecified
	}
	return &wrapperspb.UInt64Value{Value: v.Value}
}
