package protobufwrapper

import (
	"fmt"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// UInt32Value

// 1. Sentinel
var UInt32ValueUnspecified = &wrapperspb.UInt32Value{}

// 2. IsSpecified
func IsSpecifiedUInt32Value(v *wrapperspb.UInt32Value) bool {
	return v != nil && v != UInt32ValueUnspecified
}

// 3. TakeOrElse
func TakeOrElseUInt32Value(v, def *wrapperspb.UInt32Value) *wrapperspb.UInt32Value {
	if v == nil || v == UInt32ValueUnspecified {
		return def
	}
	return v
}

// 4. Merge
func MergeUInt32Value(a, b *wrapperspb.UInt32Value) *wrapperspb.UInt32Value {
	a = CoalesceUInt32Value(a, UInt32ValueUnspecified)
	b = CoalesceUInt32Value(b, UInt32ValueUnspecified)

	if a == UInt32ValueUnspecified {
		return b
	}
	if b == UInt32ValueUnspecified {
		return a
	}

	return &wrapperspb.UInt32Value{Value: b.Value}
}

// 5. String
func StringUInt32Value(v *wrapperspb.UInt32Value) string {
	if !IsSpecifiedUInt32Value(v) {
		return "UInt32Value{Unspecified}"
	}
	return fmt.Sprintf("UInt32Value{%d}", v.Value)
}

// 6. Coalesce
func CoalesceUInt32Value(ptr, def *wrapperspb.UInt32Value) *wrapperspb.UInt32Value {
	if ptr == nil {
		return def
	}
	return ptr
}

// 7. Same
func SameUInt32Value(a, b *wrapperspb.UInt32Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == UInt32ValueUnspecified
	}
	if b == nil {
		return a == UInt32ValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func SemanticEqualUInt32Value(a, b *wrapperspb.UInt32Value) bool {
	a = CoalesceUInt32Value(a, UInt32ValueUnspecified)
	b = CoalesceUInt32Value(b, UInt32ValueUnspecified)

	return a.Value == b.Value
}

// 9. Equal
func EqualUInt32Value(a, b *wrapperspb.UInt32Value) bool {
	if !SameUInt32Value(a, b) {
		return SemanticEqualUInt32Value(a, b)
	}
	return true
}

// 10. Copy
func CopyUInt32Value(v *wrapperspb.UInt32Value) *wrapperspb.UInt32Value {
	if !IsSpecifiedUInt32Value(v) {
		return UInt32ValueUnspecified
	}
	return &wrapperspb.UInt32Value{Value: v.Value}
}
