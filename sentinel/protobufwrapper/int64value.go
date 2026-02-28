package protobufwrapper

import (
	"fmt"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// Int64Value

// 1. Sentinel
var Int64ValueUnspecified = &wrapperspb.Int64Value{}

// 2. IsSpecified
func IsSpecifiedInt64Value(v *wrapperspb.Int64Value) bool {
	return v != nil && v != Int64ValueUnspecified
}

// 3. TakeOrElse
func TakeOrElseInt64Value(v, def *wrapperspb.Int64Value) *wrapperspb.Int64Value {
	if v == nil || v == Int64ValueUnspecified {
		return def
	}
	return v
}

// 4. Merge
func MergeInt64Value(a, b *wrapperspb.Int64Value) *wrapperspb.Int64Value {
	a = CoalesceInt64Value(a, Int64ValueUnspecified)
	b = CoalesceInt64Value(b, Int64ValueUnspecified)

	if a == Int64ValueUnspecified {
		return b
	}
	if b == Int64ValueUnspecified {
		return a
	}

	return &wrapperspb.Int64Value{Value: b.Value}
}

// 5. String
func StringInt64Value(v *wrapperspb.Int64Value) string {
	if !IsSpecifiedInt64Value(v) {
		return "Int64Value{Unspecified}"
	}
	return fmt.Sprintf("Int64Value{%d}", v.Value)
}

// 6. Coalesce
func CoalesceInt64Value(ptr, def *wrapperspb.Int64Value) *wrapperspb.Int64Value {
	if ptr == nil {
		return def
	}
	return ptr
}

// 7. Same
func SameInt64Value(a, b *wrapperspb.Int64Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == Int64ValueUnspecified
	}
	if b == nil {
		return a == Int64ValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func SemanticEqualInt64Value(a, b *wrapperspb.Int64Value) bool {
	a = CoalesceInt64Value(a, Int64ValueUnspecified)
	b = CoalesceInt64Value(b, Int64ValueUnspecified)

	return a.Value == b.Value
}

// 9. Equal
func EqualInt64Value(a, b *wrapperspb.Int64Value) bool {
	if !SameInt64Value(a, b) {
		return SemanticEqualInt64Value(a, b)
	}
	return true
}

// 10. Copy
func CopyInt64Value(v *wrapperspb.Int64Value) *wrapperspb.Int64Value {
	if !IsSpecifiedInt64Value(v) {
		return Int64ValueUnspecified
	}
	return &wrapperspb.Int64Value{Value: v.Value}
}
