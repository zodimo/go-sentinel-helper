package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedUInt32Value(t *testing.T) {
	require.False(t, IsSpecifiedUInt32Value(nil))
	require.False(t, IsSpecifiedUInt32Value(UInt32ValueUnspecified))
	require.True(t, IsSpecifiedUInt32Value(&wrapperspb.UInt32Value{Value: 1}))
	require.True(t, IsSpecifiedUInt32Value(&wrapperspb.UInt32Value{Value: 0}))
}

func TestTakeOrElseUInt32Value(t *testing.T) {
	def := &wrapperspb.UInt32Value{Value: 42}
	ptr := &wrapperspb.UInt32Value{Value: 1}

	require.Equal(t, def, TakeOrElseUInt32Value(nil, def))
	require.Equal(t, def, TakeOrElseUInt32Value(UInt32ValueUnspecified, def))
	require.Equal(t, ptr, TakeOrElseUInt32Value(ptr, def))
}

func TestMergeUInt32Value(t *testing.T) {
	a := &wrapperspb.UInt32Value{Value: 1}
	b := &wrapperspb.UInt32Value{Value: 2}

	result := MergeUInt32Value(nil, nil)
	require.Equal(t, UInt32ValueUnspecified, result)

	result = MergeUInt32Value(nil, b)
	require.Equal(t, uint32(2), result.Value)

	result = MergeUInt32Value(a, nil)
	require.Equal(t, uint32(1), result.Value)

	result = MergeUInt32Value(a, b)
	require.Equal(t, uint32(2), result.Value)
}

func TestStringUInt32Value(t *testing.T) {
	require.Equal(t, "UInt32Value{Unspecified}", StringUInt32Value(nil))
	require.Equal(t, "UInt32Value{Unspecified}", StringUInt32Value(UInt32ValueUnspecified))
	require.Equal(t, "UInt32Value{42}", StringUInt32Value(&wrapperspb.UInt32Value{Value: 42}))
}

func TestCoalesceUInt32Value(t *testing.T) {
	def := &wrapperspb.UInt32Value{Value: 0}
	ptr := &wrapperspb.UInt32Value{Value: 1}

	require.Equal(t, def, CoalesceUInt32Value(nil, def))
	require.Equal(t, ptr, CoalesceUInt32Value(ptr, def))
}

func TestSameUInt32Value(t *testing.T) {
	ptr := &wrapperspb.UInt32Value{}
	require.True(t, SameUInt32Value(nil, nil))
	require.True(t, SameUInt32Value(nil, UInt32ValueUnspecified))
	require.True(t, SameUInt32Value(UInt32ValueUnspecified, nil))
	require.True(t, SameUInt32Value(ptr, ptr))
	require.False(t, SameUInt32Value(nil, ptr))
	require.False(t, SameUInt32Value(ptr, nil))
	require.False(t, SameUInt32Value(ptr, UInt32ValueUnspecified))
}

func TestSemanticEqualUInt32Value(t *testing.T) {
	tests := []struct {
		name string
		a    *wrapperspb.UInt32Value
		b    *wrapperspb.UInt32Value
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, UInt32ValueUnspecified, true},
		{"both equal values", &wrapperspb.UInt32Value{Value: 1}, &wrapperspb.UInt32Value{Value: 1}, true},
		{"different values", &wrapperspb.UInt32Value{Value: 1}, &wrapperspb.UInt32Value{Value: 2}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, SemanticEqualUInt32Value(tc.a, tc.b))
		})
	}
}

func TestEqualUInt32Value(t *testing.T) {
	require.True(t, EqualUInt32Value(nil, nil))
	require.True(t, EqualUInt32Value(UInt32ValueUnspecified, UInt32ValueUnspecified))

	ptr := &wrapperspb.UInt32Value{Value: 1}
	require.True(t, EqualUInt32Value(ptr, ptr))

	ptr2 := &wrapperspb.UInt32Value{Value: 1}
	require.True(t, EqualUInt32Value(ptr, ptr2))

	ptr3 := &wrapperspb.UInt32Value{Value: 2}
	require.False(t, EqualUInt32Value(ptr, ptr3))
}

func TestCopyUInt32Value(t *testing.T) {
	require.Equal(t, UInt32ValueUnspecified, CopyUInt32Value(nil))
	require.Equal(t, UInt32ValueUnspecified, CopyUInt32Value(UInt32ValueUnspecified))

	original := &wrapperspb.UInt32Value{Value: 42}
	copied := CopyUInt32Value(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied)
}
