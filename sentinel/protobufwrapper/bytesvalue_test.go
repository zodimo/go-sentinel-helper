package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedBytesValue(t *testing.T) {
	require.False(t, IsSpecifiedBytesValue(nil))
	require.False(t, IsSpecifiedBytesValue(BytesValueUnspecified))
	require.True(t, IsSpecifiedBytesValue(&wrapperspb.BytesValue{Value: []byte{0x01}}))
	require.True(t, IsSpecifiedBytesValue(&wrapperspb.BytesValue{Value: []byte{}}))
}

func TestTakeOrElseBytesValue(t *testing.T) {
	def := &wrapperspb.BytesValue{Value: []byte{0x00}}
	ptr := &wrapperspb.BytesValue{Value: []byte{0x01}}

	require.Equal(t, def, TakeOrElseBytesValue(nil, def))
	require.Equal(t, def, TakeOrElseBytesValue(BytesValueUnspecified, def))
	require.Equal(t, ptr, TakeOrElseBytesValue(ptr, def))
}

func TestMergeBytesValue(t *testing.T) {
	a := &wrapperspb.BytesValue{Value: []byte{0x01}}
	b := &wrapperspb.BytesValue{Value: []byte{0x02}}

	result := MergeBytesValue(nil, nil)
	require.Equal(t, BytesValueUnspecified, result)

	result = MergeBytesValue(nil, b)
	require.Equal(t, []byte{0x02}, result.Value)

	result = MergeBytesValue(a, nil)
	require.Equal(t, []byte{0x01}, result.Value)

	result = MergeBytesValue(a, b)
	require.Equal(t, []byte{0x02}, result.Value)
}

func TestStringBytesValue(t *testing.T) {
	require.Equal(t, "BytesValue{Unspecified}", StringBytesValue(nil))
	require.Equal(t, "BytesValue{Unspecified}", StringBytesValue(BytesValueUnspecified))
	require.Equal(t, "BytesValue{deadbeef}", StringBytesValue(&wrapperspb.BytesValue{Value: []byte{0xde, 0xad, 0xbe, 0xef}}))
}

func TestCoalesceBytesValue(t *testing.T) {
	def := &wrapperspb.BytesValue{Value: nil}
	ptr := &wrapperspb.BytesValue{Value: []byte{0x01}}

	require.Equal(t, def, CoalesceBytesValue(nil, def))
	require.Equal(t, ptr, CoalesceBytesValue(ptr, def))
}

func TestSameBytesValue(t *testing.T) {
	ptr := &wrapperspb.BytesValue{}
	require.True(t, SameBytesValue(nil, nil))
	require.True(t, SameBytesValue(nil, BytesValueUnspecified))
	require.True(t, SameBytesValue(BytesValueUnspecified, nil))
	require.True(t, SameBytesValue(ptr, ptr))
	require.False(t, SameBytesValue(nil, ptr))
	require.False(t, SameBytesValue(ptr, nil))
	require.False(t, SameBytesValue(ptr, BytesValueUnspecified))
}

func TestSemanticEqualBytesValue(t *testing.T) {
	tests := []struct {
		name string
		a    *wrapperspb.BytesValue
		b    *wrapperspb.BytesValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, BytesValueUnspecified, true},
		{"both equal values", &wrapperspb.BytesValue{Value: []byte{0x01}}, &wrapperspb.BytesValue{Value: []byte{0x01}}, true},
		{"different values", &wrapperspb.BytesValue{Value: []byte{0x01}}, &wrapperspb.BytesValue{Value: []byte{0x02}}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, SemanticEqualBytesValue(tc.a, tc.b))
		})
	}
}

func TestEqualBytesValue(t *testing.T) {
	require.True(t, EqualBytesValue(nil, nil))
	require.True(t, EqualBytesValue(BytesValueUnspecified, BytesValueUnspecified))

	ptr := &wrapperspb.BytesValue{Value: []byte{0x01}}
	require.True(t, EqualBytesValue(ptr, ptr))

	ptr2 := &wrapperspb.BytesValue{Value: []byte{0x01}}
	require.True(t, EqualBytesValue(ptr, ptr2))

	ptr3 := &wrapperspb.BytesValue{Value: []byte{0x02}}
	require.False(t, EqualBytesValue(ptr, ptr3))
}

func TestCopyBytesValue(t *testing.T) {
	require.Equal(t, BytesValueUnspecified, CopyBytesValue(nil))
	require.Equal(t, BytesValueUnspecified, CopyBytesValue(BytesValueUnspecified))

	original := &wrapperspb.BytesValue{Value: []byte{0xde, 0xad}}
	copied := CopyBytesValue(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied)

	// verify deep copy — modifying original should not affect copy
	original.Value[0] = 0xff
	require.NotEqual(t, original.Value, copied.Value)
}
