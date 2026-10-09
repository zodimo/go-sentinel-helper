package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedBytesValue(t *testing.T) {
	c := BytesValueCompanion
	require.False(t, c.IsSpecified(nil))
	require.False(t, c.IsSpecified(c.Unspecified()))
	require.True(t, c.IsSpecified(&wrapperspb.BytesValue{Value: []byte{0x01}}))
	require.True(t, c.IsSpecified(&wrapperspb.BytesValue{Value: []byte{}}))
}

func TestTakeOrElseBytesValue(t *testing.T) {
	c := BytesValueCompanion
	def := &wrapperspb.BytesValue{Value: []byte{0x00}}
	ptr := &wrapperspb.BytesValue{Value: []byte{0x01}}

	require.Equal(t, def, c.TakeOrElse(nil, def))
	require.Equal(t, def, c.TakeOrElse(c.Unspecified(), def))
	require.Equal(t, ptr, c.TakeOrElse(ptr, def))
}

func TestMergeBytesValue(t *testing.T) {
	c := BytesValueCompanion
	a := &wrapperspb.BytesValue{Value: []byte{0x01}}
	b := &wrapperspb.BytesValue{Value: []byte{0x02}}

	result := c.Merge(nil, nil)
	require.Equal(t, c.Unspecified(), result)

	result = c.Merge(nil, b)
	require.Equal(t, []byte{0x02}, result.Value)

	result = c.Merge(a, nil)
	require.Equal(t, []byte{0x01}, result.Value)

	result = c.Merge(a, b)
	require.Equal(t, []byte{0x02}, result.Value)
}

func TestStringBytesValue(t *testing.T) {
	c := BytesValueCompanion
	require.Equal(t, "BytesValue{Unspecified}", c.String(nil))
	require.Equal(t, "BytesValue{Unspecified}", c.String(c.Unspecified()))
	require.Equal(t, "BytesValue{deadbeef}", c.String(&wrapperspb.BytesValue{Value: []byte{0xde, 0xad, 0xbe, 0xef}}))
}

func TestCoalesceBytesValue(t *testing.T) {
	c := BytesValueCompanion
	def := &wrapperspb.BytesValue{Value: nil}
	ptr := &wrapperspb.BytesValue{Value: []byte{0x01}}

	require.Equal(t, def, c.Coalesce(nil, def))
	require.Equal(t, ptr, c.Coalesce(ptr, def))
}

func TestSameBytesValue(t *testing.T) {
	c := BytesValueCompanion
	ptr := &wrapperspb.BytesValue{}
	require.True(t, c.Same(nil, nil))
	require.True(t, c.Same(nil, c.Unspecified()))
	require.True(t, c.Same(c.Unspecified(), nil))
	require.True(t, c.Same(ptr, ptr))
	require.False(t, c.Same(nil, ptr))
	require.False(t, c.Same(ptr, nil))
	require.False(t, c.Same(ptr, c.Unspecified()))
}

func TestSemanticEqualBytesValue(t *testing.T) {
	c := BytesValueCompanion
	tests := []struct {
		name string
		a    *wrapperspb.BytesValue
		b    *wrapperspb.BytesValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, c.Unspecified(), true},
		{"both equal values", &wrapperspb.BytesValue{Value: []byte{0x01}}, &wrapperspb.BytesValue{Value: []byte{0x01}}, true},
		{"different values", &wrapperspb.BytesValue{Value: []byte{0x01}}, &wrapperspb.BytesValue{Value: []byte{0x02}}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, c.SemanticEqual(tc.a, tc.b))
		})
	}
}

func TestEqualBytesValue(t *testing.T) {
	c := BytesValueCompanion
	require.True(t, c.Equal(nil, nil))
	require.True(t, c.Equal(c.Unspecified(), c.Unspecified()))

	ptr := &wrapperspb.BytesValue{Value: []byte{0x01}}
	require.True(t, c.Equal(ptr, ptr))

	ptr2 := &wrapperspb.BytesValue{Value: []byte{0x01}}
	require.True(t, c.Equal(ptr, ptr2))

	ptr3 := &wrapperspb.BytesValue{Value: []byte{0x02}}
	require.False(t, c.Equal(ptr, ptr3))
}

func TestCopyBytesValue(t *testing.T) {
	c := BytesValueCompanion
	require.Equal(t, c.Unspecified(), c.Copy(nil))
	require.Equal(t, c.Unspecified(), c.Copy(c.Unspecified()))

	original := &wrapperspb.BytesValue{Value: []byte{0xde, 0xad}}
	copied := c.Copy(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied)

	// verify deep copy — modifying original should not affect copy
	original.Value[0] = 0xff
	require.NotEqual(t, original.Value, copied.Value)
}
