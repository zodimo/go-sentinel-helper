package protobufwrapper

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// BytesValue

// 1. Sentinel
var bytesValueUnspecified = &wrapperspb.BytesValue{}

// 2. IsSpecified
func isSpecifiedBytesValue(v *wrapperspb.BytesValue) bool {
	return v != nil && v != bytesValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedBytesValue(v *wrapperspb.BytesValue) bool {
	return !isSpecifiedBytesValue(v)
}

// 4. TakeOrElse
func takeOrElseBytesValue(v, def *wrapperspb.BytesValue) *wrapperspb.BytesValue {
	if !isSpecifiedBytesValue(v) {
		return def
	}
	return v
}

// 5. Merge
func mergeBytesValue(a, b *wrapperspb.BytesValue) *wrapperspb.BytesValue {
	a = coalesceBytesValue(a, bytesValueUnspecified)
	b = coalesceBytesValue(b, bytesValueUnspecified)

	if a == bytesValueUnspecified {
		return b
	}
	if b == bytesValueUnspecified {
		return a
	}

	dst := make([]byte, len(b.Value))
	copy(dst, b.Value)
	return &wrapperspb.BytesValue{Value: dst}
}

// 6. String
func stringBytesValue(v *wrapperspb.BytesValue) string {
	if !isSpecifiedBytesValue(v) {
		return "BytesValue{Unspecified}"
	}
	return fmt.Sprintf("BytesValue{%s}", hex.EncodeToString(v.Value))
}

// 7. Same
func sameBytesValue(a, b *wrapperspb.BytesValue) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == bytesValueUnspecified
	}
	if b == nil {
		return a == bytesValueUnspecified
	}
	return a == b
}

// 8. SemanticEqual
func semanticEqualBytesValue(a, b *wrapperspb.BytesValue) bool {
	a = coalesceBytesValue(a, bytesValueUnspecified)
	b = coalesceBytesValue(b, bytesValueUnspecified)
	return bytes.Equal(a.Value, b.Value)
}

// 9. Equal
func equalBytesValue(a, b *wrapperspb.BytesValue) bool {
	if sameBytesValue(a, b) {
		return true
	}
	return semanticEqualBytesValue(a, b)
}

// 10. Coalesce
func coalesceBytesValue(ptr, def *wrapperspb.BytesValue) *wrapperspb.BytesValue {
	if ptr == nil {
		return def
	}
	return ptr
}

// 11. Copy
func copyBytesValue(v *wrapperspb.BytesValue, options ...func(*wrapperspb.BytesValue)) *wrapperspb.BytesValue {
	if !isSpecifiedBytesValue(v) {
		return bytesValueUnspecified
	}
	dst := make([]byte, len(v.Value))
	copy(dst, v.Value)
	out := &wrapperspb.BytesValue{Value: dst}
	for _, option := range options {
		option(out)
	}
	return out
}

// 12. Unspecified
func unspecifiedBytesValue() *wrapperspb.BytesValue {
	return bytesValueUnspecified
}

// BytesValueCompanion is the ONLY public companion of this type.
var BytesValueCompanion = companions.NewStructCompanion(
	unspecifiedBytesValue,
	stringBytesValue,
	isSpecifiedBytesValue,
	isUnspecifiedBytesValue,
	takeOrElseBytesValue,
	mergeBytesValue,
	sameBytesValue,
	semanticEqualBytesValue,
	equalBytesValue,
	coalesceBytesValue,
	copyBytesValue,
)
