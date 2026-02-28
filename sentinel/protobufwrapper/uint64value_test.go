package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedUInt64Value(t *testing.T) {
	require.False(t, IsSpecifiedUInt64Value(nil))
	require.False(t, IsSpecifiedUInt64Value(UInt64ValueUnspecified))
	require.True(t, IsSpecifiedUInt64Value(&wrapperspb.UInt64Value{Value: 1}))
	require.True(t, IsSpecifiedUInt64Value(&wrapperspb.UInt64Value{Value: 0}))
}

func TestTakeOrElseUInt64Value(t *testing.T) {
	def := &wrapperspb.UInt64Value{Value: 42}
	ptr := &wrapperspb.UInt64Value{Value: 1}

	require.Equal(t, def, TakeOrElseUInt64Value(nil, def))
	require.Equal(t, def, TakeOrElseUInt64Value(UInt64ValueUnspecified, def))
	require.Equal(t, ptr, TakeOrElseUInt64Value(ptr, def))
}

func TestMergeUInt64Value(t *testing.T) {
	a := &wrapperspb.UInt64Value{Value: 1}
	b := &wrapperspb.UInt64Value{Value: 2}

	result := MergeUInt64Value(nil, nil)
	require.Equal(t, UInt64ValueUnspecified, result)

	result = MergeUInt64Value(nil, b)
	require.Equal(t, uint64(2), result.Value)

	result = MergeUInt64Value(a, nil)
	require.Equal(t, uint64(1), result.Value)

	result = MergeUInt64Value(a, b)
	require.Equal(t, uint64(2), result.Value)
}

func TestStringUInt64Value(t *testing.T) {
	require.Equal(t, "UInt64Value{Unspecified}", StringUInt64Value(nil))
	require.Equal(t, "UInt64Value{Unspecified}", StringUInt64Value(UInt64ValueUnspecified))
	require.Equal(t, "UInt64Value{42}", StringUInt64Value(&wrapperspb.UInt64Value{Value: 42}))
}

func TestCoalesceUInt64Value(t *testing.T) {
	def := &wrapperspb.UInt64Value{Value: 0}
	ptr := &wrapperspb.UInt64Value{Value: 1}

	require.Equal(t, def, CoalesceUInt64Value(nil, def))
	require.Equal(t, ptr, CoalesceUInt64Value(ptr, def))
}

func TestSameUInt64Value(t *testing.T) {
	ptr := &wrapperspb.UInt64Value{}
	require.True(t, SameUInt64Value(nil, nil))
	require.True(t, SameUInt64Value(nil, UInt64ValueUnspecified))
	require.True(t, SameUInt64Value(UInt64ValueUnspecified, nil))
	require.True(t, SameUInt64Value(ptr, ptr))
	require.False(t, SameUInt64Value(nil, ptr))
	require.False(t, SameUInt64Value(ptr, nil))
	require.False(t, SameUInt64Value(ptr, UInt64ValueUnspecified))
}

func TestSemanticEqualUInt64Value(t *testing.T) {
	tests := []struct {
		name string
		a    *wrapperspb.UInt64Value
		b    *wrapperspb.UInt64Value
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, UInt64ValueUnspecified, true},
		{"both equal values", &wrapperspb.UInt64Value{Value: 1}, &wrapperspb.UInt64Value{Value: 1}, true},
		{"different values", &wrapperspb.UInt64Value{Value: 1}, &wrapperspb.UInt64Value{Value: 2}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, SemanticEqualUInt64Value(tc.a, tc.b))
		})
	}
}

func TestEqualUInt64Value(t *testing.T) {
	require.True(t, EqualUInt64Value(nil, nil))
	require.True(t, EqualUInt64Value(UInt64ValueUnspecified, UInt64ValueUnspecified))

	ptr := &wrapperspb.UInt64Value{Value: 1}
	require.True(t, EqualUInt64Value(ptr, ptr))

	ptr2 := &wrapperspb.UInt64Value{Value: 1}
	require.True(t, EqualUInt64Value(ptr, ptr2))

	ptr3 := &wrapperspb.UInt64Value{Value: 2}
	require.False(t, EqualUInt64Value(ptr, ptr3))
}

func TestCopyUInt64Value(t *testing.T) {
	require.Equal(t, UInt64ValueUnspecified, CopyUInt64Value(nil))
	require.Equal(t, UInt64ValueUnspecified, CopyUInt64Value(UInt64ValueUnspecified))

	original := &wrapperspb.UInt64Value{Value: 42}
	copied := CopyUInt64Value(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied)
}
