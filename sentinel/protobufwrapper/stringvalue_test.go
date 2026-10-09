package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedStringValue(t *testing.T) {
	c := StringValueCompanion
	require.False(t, c.IsSpecified(nil))
	require.False(t, c.IsSpecified(c.Unspecified()))
	require.True(t, c.IsSpecified(&wrapperspb.StringValue{Value: "hello"}))
	require.True(t, c.IsSpecified(&wrapperspb.StringValue{Value: ""}))
}

func TestTakeOrElseStringValue(t *testing.T) {
	c := StringValueCompanion
	def := &wrapperspb.StringValue{Value: "default"}
	ptr := &wrapperspb.StringValue{Value: "value"}

	require.Equal(t, def, c.TakeOrElse(nil, def))
	require.Equal(t, def, c.TakeOrElse(c.Unspecified(), def))
	require.Equal(t, ptr, c.TakeOrElse(ptr, def))
}

func TestMergeStringValue(t *testing.T) {
	c := StringValueCompanion
	a := &wrapperspb.StringValue{Value: "a"}
	b := &wrapperspb.StringValue{Value: "b"}

	result := c.Merge(nil, nil)
	require.Equal(t, c.Unspecified(), result)

	result = c.Merge(nil, b)
	require.Equal(t, "b", result.Value)

	result = c.Merge(a, nil)
	require.Equal(t, "a", result.Value)

	result = c.Merge(a, b)
	require.Equal(t, "b", result.Value)
}

func TestStringStringValue(t *testing.T) {
	c := StringValueCompanion
	require.Equal(t, "StringValue{Unspecified}", c.String(nil))
	require.Equal(t, "StringValue{Unspecified}", c.String(c.Unspecified()))
	require.Equal(t, `StringValue{"hello"}`, c.String(&wrapperspb.StringValue{Value: "hello"}))
}

func TestCoalesceStringValue(t *testing.T) {
	c := StringValueCompanion
	def := &wrapperspb.StringValue{Value: ""}
	ptr := &wrapperspb.StringValue{Value: "a"}

	require.Equal(t, def, c.Coalesce(nil, def))
	require.Equal(t, ptr, c.Coalesce(ptr, def))
}

func TestSameStringValue(t *testing.T) {
	c := StringValueCompanion
	ptr := &wrapperspb.StringValue{}
	require.True(t, c.Same(nil, nil))
	require.True(t, c.Same(nil, c.Unspecified()))
	require.True(t, c.Same(c.Unspecified(), nil))
	require.True(t, c.Same(ptr, ptr))
	require.False(t, c.Same(nil, ptr))
	require.False(t, c.Same(ptr, nil))
	require.False(t, c.Same(ptr, c.Unspecified()))
}

func TestSemanticEqualStringValue(t *testing.T) {
	c := StringValueCompanion
	tests := []struct {
		name string
		a    *wrapperspb.StringValue
		b    *wrapperspb.StringValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, c.Unspecified(), true},
		{"both equal values", &wrapperspb.StringValue{Value: "a"}, &wrapperspb.StringValue{Value: "a"}, true},
		{"different values", &wrapperspb.StringValue{Value: "a"}, &wrapperspb.StringValue{Value: "b"}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, c.SemanticEqual(tc.a, tc.b))
		})
	}
}

func TestEqualStringValue(t *testing.T) {
	c := StringValueCompanion
	require.True(t, c.Equal(nil, nil))
	require.True(t, c.Equal(c.Unspecified(), c.Unspecified()))

	ptr := &wrapperspb.StringValue{Value: "a"}
	require.True(t, c.Equal(ptr, ptr))

	ptr2 := &wrapperspb.StringValue{Value: "a"}
	require.True(t, c.Equal(ptr, ptr2))

	ptr3 := &wrapperspb.StringValue{Value: "b"}
	require.False(t, c.Equal(ptr, ptr3))
}

func TestCopyStringValue(t *testing.T) {
	c := StringValueCompanion
	require.Equal(t, c.Unspecified(), c.Copy(nil))
	require.Equal(t, c.Unspecified(), c.Copy(c.Unspecified()))

	original := &wrapperspb.StringValue{Value: "hello"}
	copied := c.Copy(original, func(v *wrapperspb.StringValue) { v.Value = "world" })
	require.Equal(t, "world", copied.Value)
	require.Equal(t, "hello", original.Value)
	require.True(t, original != copied)
}
