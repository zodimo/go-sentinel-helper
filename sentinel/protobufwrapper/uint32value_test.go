package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedUInt32Value(t *testing.T) {
	c := UInt32ValueCompanion
	require.False(t, c.IsSpecified(nil))
	require.False(t, c.IsSpecified(c.Unspecified()))
	require.True(t, c.IsSpecified(&wrapperspb.UInt32Value{Value: 1}))
	require.True(t, c.IsSpecified(&wrapperspb.UInt32Value{Value: 0}))
}

func TestTakeOrElseUInt32Value(t *testing.T) {
	c := UInt32ValueCompanion
	def := &wrapperspb.UInt32Value{Value: 42}
	ptr := &wrapperspb.UInt32Value{Value: 7}

	require.Equal(t, def, c.TakeOrElse(nil, def))
	require.Equal(t, def, c.TakeOrElse(c.Unspecified(), def))
	require.Equal(t, ptr, c.TakeOrElse(ptr, def))
}

func TestMergeUInt32Value(t *testing.T) {
	c := UInt32ValueCompanion
	a := &wrapperspb.UInt32Value{Value: 1}
	b := &wrapperspb.UInt32Value{Value: 2}

	result := c.Merge(nil, nil)
	require.Equal(t, c.Unspecified(), result)

	result = c.Merge(nil, b)
	require.Equal(t, uint32(2), result.Value)

	result = c.Merge(a, nil)
	require.Equal(t, uint32(1), result.Value)

	result = c.Merge(a, b)
	require.Equal(t, uint32(2), result.Value)
}

func TestStringUInt32Value(t *testing.T) {
	c := UInt32ValueCompanion
	require.Equal(t, "UInt32Value{Unspecified}", c.String(nil))
	require.Equal(t, "UInt32Value{Unspecified}", c.String(c.Unspecified()))
	require.Equal(t, "UInt32Value{42}", c.String(&wrapperspb.UInt32Value{Value: 42}))
}

func TestCoalesceUInt32Value(t *testing.T) {
	c := UInt32ValueCompanion
	def := &wrapperspb.UInt32Value{Value: 0}
	ptr := &wrapperspb.UInt32Value{Value: 1}

	require.Equal(t, def, c.Coalesce(nil, def))
	require.Equal(t, ptr, c.Coalesce(ptr, def))
}

func TestSameUInt32Value(t *testing.T) {
	c := UInt32ValueCompanion
	ptr := &wrapperspb.UInt32Value{}
	require.True(t, c.Same(nil, nil))
	require.True(t, c.Same(nil, c.Unspecified()))
	require.True(t, c.Same(c.Unspecified(), nil))
	require.True(t, c.Same(ptr, ptr))
	require.False(t, c.Same(nil, ptr))
	require.False(t, c.Same(ptr, nil))
	require.False(t, c.Same(ptr, c.Unspecified()))
}

func TestSemanticEqualUInt32Value(t *testing.T) {
	c := UInt32ValueCompanion
	tests := []struct {
		name string
		a    *wrapperspb.UInt32Value
		b    *wrapperspb.UInt32Value
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, c.Unspecified(), true},
		{"both equal values", &wrapperspb.UInt32Value{Value: 1}, &wrapperspb.UInt32Value{Value: 1}, true},
		{"different values", &wrapperspb.UInt32Value{Value: 1}, &wrapperspb.UInt32Value{Value: 2}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, c.SemanticEqual(tc.a, tc.b))
		})
	}
}

func TestEqualUInt32Value(t *testing.T) {
	c := UInt32ValueCompanion
	require.True(t, c.Equal(nil, nil))
	require.True(t, c.Equal(c.Unspecified(), c.Unspecified()))

	ptr := &wrapperspb.UInt32Value{Value: 1}
	require.True(t, c.Equal(ptr, ptr))

	ptr2 := &wrapperspb.UInt32Value{Value: 1}
	require.True(t, c.Equal(ptr, ptr2))

	ptr3 := &wrapperspb.UInt32Value{Value: 2}
	require.False(t, c.Equal(ptr, ptr3))
}

func TestCopyUInt32Value(t *testing.T) {
	c := UInt32ValueCompanion
	require.Equal(t, c.Unspecified(), c.Copy(nil))
	require.Equal(t, c.Unspecified(), c.Copy(c.Unspecified()))

	original := &wrapperspb.UInt32Value{Value: 1}
	copied := c.Copy(original, func(v *wrapperspb.UInt32Value) { v.Value = 9 })
	require.Equal(t, uint32(9), copied.Value)
	require.Equal(t, uint32(1), original.Value)
	require.True(t, original != copied)
}
