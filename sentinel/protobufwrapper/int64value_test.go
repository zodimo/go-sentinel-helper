package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedInt64Value(t *testing.T) {
	require.False(t, IsSpecifiedInt64Value(nil))
	require.False(t, IsSpecifiedInt64Value(Int64ValueUnspecified))
	require.True(t, IsSpecifiedInt64Value(&wrapperspb.Int64Value{Value: 1}))
	require.True(t, IsSpecifiedInt64Value(&wrapperspb.Int64Value{Value: 0}))
}

func TestTakeOrElseInt64Value(t *testing.T) {
	def := &wrapperspb.Int64Value{Value: 42}
	ptr := &wrapperspb.Int64Value{Value: 1}

	require.Equal(t, def, TakeOrElseInt64Value(nil, def))
	require.Equal(t, def, TakeOrElseInt64Value(Int64ValueUnspecified, def))
	require.Equal(t, ptr, TakeOrElseInt64Value(ptr, def))
}

func TestMergeInt64Value(t *testing.T) {
	a := &wrapperspb.Int64Value{Value: 1}
	b := &wrapperspb.Int64Value{Value: 2}

	result := MergeInt64Value(nil, nil)
	require.Equal(t, Int64ValueUnspecified, result)

	result = MergeInt64Value(nil, b)
	require.Equal(t, int64(2), result.Value)

	result = MergeInt64Value(a, nil)
	require.Equal(t, int64(1), result.Value)

	result = MergeInt64Value(a, b)
	require.Equal(t, int64(2), result.Value)
}

func TestStringInt64Value(t *testing.T) {
	require.Equal(t, "Int64Value{Unspecified}", StringInt64Value(nil))
	require.Equal(t, "Int64Value{Unspecified}", StringInt64Value(Int64ValueUnspecified))
	require.Equal(t, "Int64Value{42}", StringInt64Value(&wrapperspb.Int64Value{Value: 42}))
}

func TestCoalesceInt64Value(t *testing.T) {
	def := &wrapperspb.Int64Value{Value: 0}
	ptr := &wrapperspb.Int64Value{Value: 1}

	require.Equal(t, def, CoalesceInt64Value(nil, def))
	require.Equal(t, ptr, CoalesceInt64Value(ptr, def))
}

func TestSameInt64Value(t *testing.T) {
	ptr := &wrapperspb.Int64Value{}
	require.True(t, SameInt64Value(nil, nil))
	require.True(t, SameInt64Value(nil, Int64ValueUnspecified))
	require.True(t, SameInt64Value(Int64ValueUnspecified, nil))
	require.True(t, SameInt64Value(ptr, ptr))
	require.False(t, SameInt64Value(nil, ptr))
	require.False(t, SameInt64Value(ptr, nil))
	require.False(t, SameInt64Value(ptr, Int64ValueUnspecified))
}

func TestSemanticEqualInt64Value(t *testing.T) {
	tests := []struct {
		name string
		a    *wrapperspb.Int64Value
		b    *wrapperspb.Int64Value
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, Int64ValueUnspecified, true},
		{"both equal values", &wrapperspb.Int64Value{Value: 1}, &wrapperspb.Int64Value{Value: 1}, true},
		{"different values", &wrapperspb.Int64Value{Value: 1}, &wrapperspb.Int64Value{Value: 2}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, SemanticEqualInt64Value(tc.a, tc.b))
		})
	}
}

func TestEqualInt64Value(t *testing.T) {
	require.True(t, EqualInt64Value(nil, nil))
	require.True(t, EqualInt64Value(Int64ValueUnspecified, Int64ValueUnspecified))

	ptr := &wrapperspb.Int64Value{Value: 1}
	require.True(t, EqualInt64Value(ptr, ptr))

	ptr2 := &wrapperspb.Int64Value{Value: 1}
	require.True(t, EqualInt64Value(ptr, ptr2))

	ptr3 := &wrapperspb.Int64Value{Value: 2}
	require.False(t, EqualInt64Value(ptr, ptr3))
}

func TestCopyInt64Value(t *testing.T) {
	require.Equal(t, Int64ValueUnspecified, CopyInt64Value(nil))
	require.Equal(t, Int64ValueUnspecified, CopyInt64Value(Int64ValueUnspecified))

	original := &wrapperspb.Int64Value{Value: 42}
	copied := CopyInt64Value(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied)
}
