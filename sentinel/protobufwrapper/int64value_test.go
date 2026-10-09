package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedInt64Value(t *testing.T) {
	c := Int64ValueCompanion
	require.False(t, c.IsSpecified(nil))
	require.False(t, c.IsSpecified(c.Unspecified()))
	require.True(t, c.IsSpecified(&wrapperspb.Int64Value{Value: 1}))
	require.True(t, c.IsSpecified(&wrapperspb.Int64Value{Value: 0}))
}

func TestTakeOrElseInt64Value(t *testing.T) {
	c := Int64ValueCompanion
	def := &wrapperspb.Int64Value{Value: 42}
	ptr := &wrapperspb.Int64Value{Value: 7}

	require.Equal(t, def, c.TakeOrElse(nil, def))
	require.Equal(t, def, c.TakeOrElse(c.Unspecified(), def))
	require.Equal(t, ptr, c.TakeOrElse(ptr, def))
}

func TestMergeInt64Value(t *testing.T) {
	c := Int64ValueCompanion
	a := &wrapperspb.Int64Value{Value: 1}
	b := &wrapperspb.Int64Value{Value: 2}

	result := c.Merge(nil, nil)
	require.Equal(t, c.Unspecified(), result)

	result = c.Merge(nil, b)
	require.Equal(t, int64(2), result.Value)

	result = c.Merge(a, nil)
	require.Equal(t, int64(1), result.Value)

	result = c.Merge(a, b)
	require.Equal(t, int64(2), result.Value)
}

func TestStringInt64Value(t *testing.T) {
	c := Int64ValueCompanion
	require.Equal(t, "Int64Value{Unspecified}", c.String(nil))
	require.Equal(t, "Int64Value{Unspecified}", c.String(c.Unspecified()))
	require.Equal(t, "Int64Value{42}", c.String(&wrapperspb.Int64Value{Value: 42}))
}

func TestCoalesceInt64Value(t *testing.T) {
	c := Int64ValueCompanion
	def := &wrapperspb.Int64Value{Value: 0}
	ptr := &wrapperspb.Int64Value{Value: 1}

	require.Equal(t, def, c.Coalesce(nil, def))
	require.Equal(t, ptr, c.Coalesce(ptr, def))
}

func TestSameInt64Value(t *testing.T) {
	c := Int64ValueCompanion
	ptr := &wrapperspb.Int64Value{}
	require.True(t, c.Same(nil, nil))
	require.True(t, c.Same(nil, c.Unspecified()))
	require.True(t, c.Same(c.Unspecified(), nil))
	require.True(t, c.Same(ptr, ptr))
	require.False(t, c.Same(nil, ptr))
	require.False(t, c.Same(ptr, nil))
	require.False(t, c.Same(ptr, c.Unspecified()))
}

func TestSemanticEqualInt64Value(t *testing.T) {
	c := Int64ValueCompanion
	tests := []struct {
		name string
		a    *wrapperspb.Int64Value
		b    *wrapperspb.Int64Value
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, c.Unspecified(), true},
		{"both equal values", &wrapperspb.Int64Value{Value: 1}, &wrapperspb.Int64Value{Value: 1}, true},
		{"different values", &wrapperspb.Int64Value{Value: 1}, &wrapperspb.Int64Value{Value: 2}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, c.SemanticEqual(tc.a, tc.b))
		})
	}
}

func TestEqualInt64Value(t *testing.T) {
	c := Int64ValueCompanion
	require.True(t, c.Equal(nil, nil))
	require.True(t, c.Equal(c.Unspecified(), c.Unspecified()))

	ptr := &wrapperspb.Int64Value{Value: 1}
	require.True(t, c.Equal(ptr, ptr))

	ptr2 := &wrapperspb.Int64Value{Value: 1}
	require.True(t, c.Equal(ptr, ptr2))

	ptr3 := &wrapperspb.Int64Value{Value: 2}
	require.False(t, c.Equal(ptr, ptr3))
}

func TestCopyInt64Value(t *testing.T) {
	c := Int64ValueCompanion
	require.Equal(t, c.Unspecified(), c.Copy(nil))
	require.Equal(t, c.Unspecified(), c.Copy(c.Unspecified()))

	original := &wrapperspb.Int64Value{Value: 1}
	copied := c.Copy(original, func(v *wrapperspb.Int64Value) { v.Value = 9 })
	require.Equal(t, int64(9), copied.Value)
	require.Equal(t, int64(1), original.Value)
	require.True(t, original != copied)
}
