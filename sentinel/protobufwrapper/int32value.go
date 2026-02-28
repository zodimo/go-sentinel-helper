package protobufwrapper

import (
	"fmt"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// Int32Value

// 1. Sentinel
var Int32ValueUnspecified = &wrapperspb.Int32Value{}

// 2. IsSpecified
func IsSpecifiedInt32Value(v *wrapperspb.Int32Value) bool {
	return v != nil && v != Int32ValueUnspecified
}

// 3. TakeOrElse
func TakeOrElseInt32Value(v, def *wrapperspb.Int32Value) *wrapperspb.Int32Value {
	if v == nil || v == Int32ValueUnspecified {
		return def
	}
	return v
}

// 4. Merge
func MergeInt32Value(a, b *wrapperspb.Int32Value) *wrapperspb.Int32Value {
	a = CoalesceInt32Value(a, Int32ValueUnspecified)
	b = CoalesceInt32Value(b, Int32ValueUnspecified)

	if a == Int32ValueUnspecified {
		return b
	}
	if b == Int32ValueUnspecified {
		return a
	}

	return &wrapperspb.Int32Value{Value: b.Value}
}

// 5. String
func StringInt32Value(v *wrapperspb.Int32Value) string {
	if !IsSpecifiedInt32Value(v) {
		return "Int32Value{Unspecified}"
	}
	return fmt.Sprintf("Int32Value{%d}", v.Value)
}

// 6. Coalesce
func CoalesceInt32Value(ptr, def *wrapperspb.Int32Value) *wrapperspb.Int32Value {
	if ptr == nil {
		return def
	}
	return ptr
}

// 7. Same
func SameInt32Value(a, b *wrapperspb.Int32Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == Int32ValueUnspecified
	}
	if b == nil {
		return a == Int32ValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func SemanticEqualInt32Value(a, b *wrapperspb.Int32Value) bool {
	a = CoalesceInt32Value(a, Int32ValueUnspecified)
	b = CoalesceInt32Value(b, Int32ValueUnspecified)

	return a.Value == b.Value
}

// 9. Equal
func EqualInt32Value(a, b *wrapperspb.Int32Value) bool {
	if !SameInt32Value(a, b) {
		return SemanticEqualInt32Value(a, b)
	}
	return true
}

// 10. Copy
func CopyInt32Value(v *wrapperspb.Int32Value) *wrapperspb.Int32Value {
	if !IsSpecifiedInt32Value(v) {
		return Int32ValueUnspecified
	}
	return &wrapperspb.Int32Value{Value: v.Value}
}
