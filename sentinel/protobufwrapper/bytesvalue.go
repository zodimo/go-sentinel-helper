package protobufwrapper

import (
	"bytes"
	"encoding/hex"
	"fmt"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// BytesValue

// 1. Sentinel
var BytesValueUnspecified = &wrapperspb.BytesValue{}

// 2. IsSpecified
func IsSpecifiedBytesValue(v *wrapperspb.BytesValue) bool {
	return v != nil && v != BytesValueUnspecified
}

// 3. TakeOrElse
func TakeOrElseBytesValue(v, def *wrapperspb.BytesValue) *wrapperspb.BytesValue {
	if v == nil || v == BytesValueUnspecified {
		return def
	}
	return v
}

// 4. Merge
func MergeBytesValue(a, b *wrapperspb.BytesValue) *wrapperspb.BytesValue {
	a = CoalesceBytesValue(a, BytesValueUnspecified)
	b = CoalesceBytesValue(b, BytesValueUnspecified)

	if a == BytesValueUnspecified {
		return b
	}
	if b == BytesValueUnspecified {
		return a
	}

	dst := make([]byte, len(b.Value))
	copy(dst, b.Value)
	return &wrapperspb.BytesValue{Value: dst}
}

// 5. String
func StringBytesValue(v *wrapperspb.BytesValue) string {
	if !IsSpecifiedBytesValue(v) {
		return "BytesValue{Unspecified}"
	}
	return fmt.Sprintf("BytesValue{%s}", hex.EncodeToString(v.Value))
}

// 6. Coalesce
func CoalesceBytesValue(ptr, def *wrapperspb.BytesValue) *wrapperspb.BytesValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 7. Same
func SameBytesValue(a, b *wrapperspb.BytesValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == BytesValueUnspecified
	}
	if b == nil {
		return a == BytesValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func SemanticEqualBytesValue(a, b *wrapperspb.BytesValue) bool {
	a = CoalesceBytesValue(a, BytesValueUnspecified)
	b = CoalesceBytesValue(b, BytesValueUnspecified)

	return bytes.Equal(a.Value, b.Value)
}

// 9. Equal
func EqualBytesValue(a, b *wrapperspb.BytesValue) bool {
	if !SameBytesValue(a, b) {
		return SemanticEqualBytesValue(a, b)
	}
	return true
}

// 10. Copy
func CopyBytesValue(v *wrapperspb.BytesValue) *wrapperspb.BytesValue {
	if !IsSpecifiedBytesValue(v) {
		return BytesValueUnspecified
	}
	dst := make([]byte, len(v.Value))
	copy(dst, v.Value)
	return &wrapperspb.BytesValue{Value: dst}
}
