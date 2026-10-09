package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedBoolValue(t *testing.T) {
	c := BoolValueCompanion
	require.False(t, c.IsSpecified(nil))
	require.False(t, c.IsSpecified(c.Unspecified()))
	require.True(t, c.IsSpecified(&wrapperspb.BoolValue{Value: true}))
	require.True(t, c.IsSpecified(&wrapperspb.BoolValue{Value: false}))
}

func TestTakeOrElseBoolValue(t *testing.T) {
	c := BoolValueCompanion
	def := &wrapperspb.BoolValue{Value: true}
	ptr := &wrapperspb.BoolValue{Value: false}

	require.Equal(t, def, c.TakeOrElse(nil, def))
	require.Equal(t, def, c.TakeOrElse(c.Unspecified(), def))
	require.Equal(t, ptr, c.TakeOrElse(ptr, def))
}

func TestMergeBoolValue(t *testing.T) {
	c := BoolValueCompanion
	a := &wrapperspb.BoolValue{Value: true}
	b := &wrapperspb.BoolValue{Value: false}

	// both unspecified
	result := c.Merge(nil, nil)
	require.Equal(t, c.Unspecified(), result)

	// a unspecified, b specified
	result = c.Merge(nil, b)
	require.Equal(t, false, result.Value)

	// a specified, b unspecified
	result = c.Merge(a, nil)
	require.Equal(t, true, result.Value)

	// both specified — b wins
	result = c.Merge(a, b)
	require.Equal(t, false, result.Value)
}

func TestStringBoolValue(t *testing.T) {
	c := BoolValueCompanion
	require.Equal(t, "BoolValue{Unspecified}", c.String(nil))
	require.Equal(t, "BoolValue{Unspecified}", c.String(c.Unspecified()))
	require.Equal(t, "BoolValue{true}", c.String(&wrapperspb.BoolValue{Value: true}))
	require.Equal(t, "BoolValue{false}", c.String(&wrapperspb.BoolValue{Value: false}))
}

func TestCoalesceBoolValue(t *testing.T) {
	c := BoolValueCompanion
	def := &wrapperspb.BoolValue{Value: false}
	ptr := &wrapperspb.BoolValue{Value: true}

	require.Equal(t, def, c.Coalesce(nil, def))
	require.Equal(t, ptr, c.Coalesce(ptr, def))
}

func TestSameBoolValue(t *testing.T) {
	c := BoolValueCompanion
	ptr := &wrapperspb.BoolValue{}
	require.True(t, c.Same(nil, nil))
	require.True(t, c.Same(nil, c.Unspecified()))
	require.True(t, c.Same(c.Unspecified(), nil))
	require.True(t, c.Same(ptr, ptr))
	require.False(t, c.Same(nil, ptr))
	require.False(t, c.Same(ptr, nil))
	require.False(t, c.Same(ptr, c.Unspecified()))
}

func TestSemanticEqualBoolValue(t *testing.T) {
	c := BoolValueCompanion
	tests := []struct {
		name string
		a    *wrapperspb.BoolValue
		b    *wrapperspb.BoolValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, c.Unspecified(), true},
		{"both equal values", &wrapperspb.BoolValue{Value: true}, &wrapperspb.BoolValue{Value: true}, true},
		{"different values", &wrapperspb.BoolValue{Value: true}, &wrapperspb.BoolValue{Value: false}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, c.SemanticEqual(tc.a, tc.b))
		})
	}
}

func TestEqualBoolValue(t *testing.T) {
	c := BoolValueCompanion
	require.True(t, c.Equal(nil, nil))
	require.True(t, c.Equal(c.Unspecified(), c.Unspecified()))

	ptr := &wrapperspb.BoolValue{Value: true}
	require.True(t, c.Equal(ptr, ptr))

	ptr2 := &wrapperspb.BoolValue{Value: true}
	require.True(t, c.Equal(ptr, ptr2))

	ptr3 := &wrapperspb.BoolValue{Value: false}
	require.False(t, c.Equal(ptr, ptr3))
}

func TestCopyBoolValue(t *testing.T) {
	c := BoolValueCompanion
	require.Equal(t, c.Unspecified(), c.Copy(nil))
	require.Equal(t, c.Unspecified(), c.Copy(c.Unspecified()))

	original := &wrapperspb.BoolValue{Value: true}
	copied := c.Copy(original, func(v *wrapperspb.BoolValue) { v.Value = false })
	require.Equal(t, false, copied.Value)
	require.Equal(t, true, original.Value)
	require.True(t, original != copied) // different pointer
}
