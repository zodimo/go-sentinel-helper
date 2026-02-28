package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedInt32Value(t *testing.T) {
	require.False(t, IsSpecifiedInt32Value(nil))
	require.False(t, IsSpecifiedInt32Value(Int32ValueUnspecified))
	require.True(t, IsSpecifiedInt32Value(&wrapperspb.Int32Value{Value: 1}))
	require.True(t, IsSpecifiedInt32Value(&wrapperspb.Int32Value{Value: 0}))
}

func TestTakeOrElseInt32Value(t *testing.T) {
	def := &wrapperspb.Int32Value{Value: 42}
	ptr := &wrapperspb.Int32Value{Value: 1}

	require.Equal(t, def, TakeOrElseInt32Value(nil, def))
	require.Equal(t, def, TakeOrElseInt32Value(Int32ValueUnspecified, def))
	require.Equal(t, ptr, TakeOrElseInt32Value(ptr, def))
}

func TestMergeInt32Value(t *testing.T) {
	a := &wrapperspb.Int32Value{Value: 1}
	b := &wrapperspb.Int32Value{Value: 2}

	result := MergeInt32Value(nil, nil)
	require.Equal(t, Int32ValueUnspecified, result)

	result = MergeInt32Value(nil, b)
	require.Equal(t, int32(2), result.Value)

	result = MergeInt32Value(a, nil)
	require.Equal(t, int32(1), result.Value)

	result = MergeInt32Value(a, b)
	require.Equal(t, int32(2), result.Value)
}

func TestStringInt32Value(t *testing.T) {
	require.Equal(t, "Int32Value{Unspecified}", StringInt32Value(nil))
	require.Equal(t, "Int32Value{Unspecified}", StringInt32Value(Int32ValueUnspecified))
	require.Equal(t, "Int32Value{42}", StringInt32Value(&wrapperspb.Int32Value{Value: 42}))
}

func TestCoalesceInt32Value(t *testing.T) {
	def := &wrapperspb.Int32Value{Value: 0}
	ptr := &wrapperspb.Int32Value{Value: 1}

	require.Equal(t, def, CoalesceInt32Value(nil, def))
	require.Equal(t, ptr, CoalesceInt32Value(ptr, def))
}

func TestSameInt32Value(t *testing.T) {
	ptr := &wrapperspb.Int32Value{}
	require.True(t, SameInt32Value(nil, nil))
	require.True(t, SameInt32Value(nil, Int32ValueUnspecified))
	require.True(t, SameInt32Value(Int32ValueUnspecified, nil))
	require.True(t, SameInt32Value(ptr, ptr))
	require.False(t, SameInt32Value(nil, ptr))
	require.False(t, SameInt32Value(ptr, nil))
	require.False(t, SameInt32Value(ptr, Int32ValueUnspecified))
}

func TestSemanticEqualInt32Value(t *testing.T) {
	tests := []struct {
		name string
		a    *wrapperspb.Int32Value
		b    *wrapperspb.Int32Value
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, Int32ValueUnspecified, true},
		{"both equal values", &wrapperspb.Int32Value{Value: 1}, &wrapperspb.Int32Value{Value: 1}, true},
		{"different values", &wrapperspb.Int32Value{Value: 1}, &wrapperspb.Int32Value{Value: 2}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, SemanticEqualInt32Value(tc.a, tc.b))
		})
	}
}

func TestEqualInt32Value(t *testing.T) {
	require.True(t, EqualInt32Value(nil, nil))
	require.True(t, EqualInt32Value(Int32ValueUnspecified, Int32ValueUnspecified))

	ptr := &wrapperspb.Int32Value{Value: 1}
	require.True(t, EqualInt32Value(ptr, ptr))

	ptr2 := &wrapperspb.Int32Value{Value: 1}
	require.True(t, EqualInt32Value(ptr, ptr2))

	ptr3 := &wrapperspb.Int32Value{Value: 2}
	require.False(t, EqualInt32Value(ptr, ptr3))
}

func TestCopyInt32Value(t *testing.T) {
	require.Equal(t, Int32ValueUnspecified, CopyInt32Value(nil))
	require.Equal(t, Int32ValueUnspecified, CopyInt32Value(Int32ValueUnspecified))

	original := &wrapperspb.Int32Value{Value: 42}
	copied := CopyInt32Value(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied)
}
