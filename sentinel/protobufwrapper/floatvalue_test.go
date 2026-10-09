package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedFloatValue(t *testing.T) {
	c := FloatValueCompanion
	require.False(t, c.IsSpecified(nil))
	require.False(t, c.IsSpecified(c.Unspecified()))
	require.True(t, c.IsSpecified(&wrapperspb.FloatValue{Value: 1.5}))
	require.True(t, c.IsSpecified(&wrapperspb.FloatValue{Value: 0}))
}

func TestTakeOrElseFloatValue(t *testing.T) {
	c := FloatValueCompanion
	def := &wrapperspb.FloatValue{Value: 1.5}
	ptr := &wrapperspb.FloatValue{Value: 2.5}

	require.Equal(t, def, c.TakeOrElse(nil, def))
	require.Equal(t, def, c.TakeOrElse(c.Unspecified(), def))
	require.Equal(t, ptr, c.TakeOrElse(ptr, def))
}

func TestMergeFloatValue(t *testing.T) {
	c := FloatValueCompanion
	a := &wrapperspb.FloatValue{Value: 1.5}
	b := &wrapperspb.FloatValue{Value: 2.5}

	result := c.Merge(nil, nil)
	require.Equal(t, c.Unspecified(), result)

	result = c.Merge(nil, b)
	require.Equal(t, float32(2.5), result.Value)

	result = c.Merge(a, nil)
	require.Equal(t, float32(1.5), result.Value)

	result = c.Merge(a, b)
	require.Equal(t, float32(2.5), result.Value)
}

func TestStringFloatValue(t *testing.T) {
	c := FloatValueCompanion
	require.Equal(t, "FloatValue{Unspecified}", c.String(nil))
	require.Equal(t, "FloatValue{Unspecified}", c.String(c.Unspecified()))
	require.Equal(t, "FloatValue{1.5}", c.String(&wrapperspb.FloatValue{Value: 1.5}))
}

func TestCoalesceFloatValue(t *testing.T) {
	c := FloatValueCompanion
	def := &wrapperspb.FloatValue{Value: 0}
	ptr := &wrapperspb.FloatValue{Value: 1.5}

	require.Equal(t, def, c.Coalesce(nil, def))
	require.Equal(t, ptr, c.Coalesce(ptr, def))
}

func TestSameFloatValue(t *testing.T) {
	c := FloatValueCompanion
	ptr := &wrapperspb.FloatValue{}
	require.True(t, c.Same(nil, nil))
	require.True(t, c.Same(nil, c.Unspecified()))
	require.True(t, c.Same(c.Unspecified(), nil))
	require.True(t, c.Same(ptr, ptr))
	require.False(t, c.Same(nil, ptr))
	require.False(t, c.Same(ptr, nil))
	require.False(t, c.Same(ptr, c.Unspecified()))
}

func TestSemanticEqualFloatValue(t *testing.T) {
	c := FloatValueCompanion
	tests := []struct {
		name string
		a    *wrapperspb.FloatValue
		b    *wrapperspb.FloatValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, c.Unspecified(), true},
		{"both equal values", &wrapperspb.FloatValue{Value: 1.5}, &wrapperspb.FloatValue{Value: 1.5}, true},
		{"different values", &wrapperspb.FloatValue{Value: 1.5}, &wrapperspb.FloatValue{Value: 2.5}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, c.SemanticEqual(tc.a, tc.b))
		})
	}
}

func TestEqualFloatValue(t *testing.T) {
	c := FloatValueCompanion
	require.True(t, c.Equal(nil, nil))
	require.True(t, c.Equal(c.Unspecified(), c.Unspecified()))

	ptr := &wrapperspb.FloatValue{Value: 1.5}
	require.True(t, c.Equal(ptr, ptr))

	ptr2 := &wrapperspb.FloatValue{Value: 1.5}
	require.True(t, c.Equal(ptr, ptr2))

	ptr3 := &wrapperspb.FloatValue{Value: 2.5}
	require.False(t, c.Equal(ptr, ptr3))
}

func TestCopyFloatValue(t *testing.T) {
	c := FloatValueCompanion
	require.Equal(t, c.Unspecified(), c.Copy(nil))
	require.Equal(t, c.Unspecified(), c.Copy(c.Unspecified()))

	original := &wrapperspb.FloatValue{Value: 1.5}
	copied := c.Copy(original, func(v *wrapperspb.FloatValue) { v.Value = 9.5 })
	require.Equal(t, float32(9.5), copied.Value)
	require.Equal(t, float32(1.5), original.Value)
	require.True(t, original != copied)
}
