package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedDoubleValue(t *testing.T) {
	c := DoubleValueCompanion
	require.False(t, c.IsSpecified(nil))
	require.False(t, c.IsSpecified(c.Unspecified()))
	require.True(t, c.IsSpecified(&wrapperspb.DoubleValue{Value: 1.5}))
	require.True(t, c.IsSpecified(&wrapperspb.DoubleValue{Value: 0}))
}

func TestTakeOrElseDoubleValue(t *testing.T) {
	c := DoubleValueCompanion
	def := &wrapperspb.DoubleValue{Value: 1.5}
	ptr := &wrapperspb.DoubleValue{Value: 2.5}

	require.Equal(t, def, c.TakeOrElse(nil, def))
	require.Equal(t, def, c.TakeOrElse(c.Unspecified(), def))
	require.Equal(t, ptr, c.TakeOrElse(ptr, def))
}

func TestMergeDoubleValue(t *testing.T) {
	c := DoubleValueCompanion
	a := &wrapperspb.DoubleValue{Value: 1.5}
	b := &wrapperspb.DoubleValue{Value: 2.5}

	result := c.Merge(nil, nil)
	require.Equal(t, c.Unspecified(), result)

	result = c.Merge(nil, b)
	require.Equal(t, float64(2.5), result.Value)

	result = c.Merge(a, nil)
	require.Equal(t, float64(1.5), result.Value)

	result = c.Merge(a, b)
	require.Equal(t, float64(2.5), result.Value)
}

func TestStringDoubleValue(t *testing.T) {
	c := DoubleValueCompanion
	require.Equal(t, "DoubleValue{Unspecified}", c.String(nil))
	require.Equal(t, "DoubleValue{Unspecified}", c.String(c.Unspecified()))
	require.Equal(t, "DoubleValue{1.5}", c.String(&wrapperspb.DoubleValue{Value: 1.5}))
}

func TestCoalesceDoubleValue(t *testing.T) {
	c := DoubleValueCompanion
	def := &wrapperspb.DoubleValue{Value: 0}
	ptr := &wrapperspb.DoubleValue{Value: 1.5}

	require.Equal(t, def, c.Coalesce(nil, def))
	require.Equal(t, ptr, c.Coalesce(ptr, def))
}

func TestSameDoubleValue(t *testing.T) {
	c := DoubleValueCompanion
	ptr := &wrapperspb.DoubleValue{}
	require.True(t, c.Same(nil, nil))
	require.True(t, c.Same(nil, c.Unspecified()))
	require.True(t, c.Same(c.Unspecified(), nil))
	require.True(t, c.Same(ptr, ptr))
	require.False(t, c.Same(nil, ptr))
	require.False(t, c.Same(ptr, nil))
	require.False(t, c.Same(ptr, c.Unspecified()))
}

func TestSemanticEqualDoubleValue(t *testing.T) {
	c := DoubleValueCompanion
	tests := []struct {
		name string
		a    *wrapperspb.DoubleValue
		b    *wrapperspb.DoubleValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, c.Unspecified(), true},
		{"both equal values", &wrapperspb.DoubleValue{Value: 1.5}, &wrapperspb.DoubleValue{Value: 1.5}, true},
		{"different values", &wrapperspb.DoubleValue{Value: 1.5}, &wrapperspb.DoubleValue{Value: 2.5}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, c.SemanticEqual(tc.a, tc.b))
		})
	}
}

func TestEqualDoubleValue(t *testing.T) {
	c := DoubleValueCompanion
	require.True(t, c.Equal(nil, nil))
	require.True(t, c.Equal(c.Unspecified(), c.Unspecified()))

	ptr := &wrapperspb.DoubleValue{Value: 1.5}
	require.True(t, c.Equal(ptr, ptr))

	ptr2 := &wrapperspb.DoubleValue{Value: 1.5}
	require.True(t, c.Equal(ptr, ptr2))

	ptr3 := &wrapperspb.DoubleValue{Value: 2.5}
	require.False(t, c.Equal(ptr, ptr3))
}

func TestCopyDoubleValue(t *testing.T) {
	c := DoubleValueCompanion
	require.Equal(t, c.Unspecified(), c.Copy(nil))
	require.Equal(t, c.Unspecified(), c.Copy(c.Unspecified()))

	original := &wrapperspb.DoubleValue{Value: 1.5}
	copied := c.Copy(original, func(v *wrapperspb.DoubleValue) { v.Value = 9.5 })
	require.Equal(t, float64(9.5), copied.Value)
	require.Equal(t, float64(1.5), original.Value)
	require.True(t, original != copied)
}
