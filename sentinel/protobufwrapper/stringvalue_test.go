package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedStringValue(t *testing.T) {
	require.False(t, IsSpecifiedStringValue(nil))
	require.False(t, IsSpecifiedStringValue(StringValueUnspecified))
	require.True(t, IsSpecifiedStringValue(&wrapperspb.StringValue{Value: "hello"}))
	require.True(t, IsSpecifiedStringValue(&wrapperspb.StringValue{Value: ""}))
}

func TestTakeOrElseStringValue(t *testing.T) {
	def := &wrapperspb.StringValue{Value: "default"}
	ptr := &wrapperspb.StringValue{Value: "value"}

	require.Equal(t, def, TakeOrElseStringValue(nil, def))
	require.Equal(t, def, TakeOrElseStringValue(StringValueUnspecified, def))
	require.Equal(t, ptr, TakeOrElseStringValue(ptr, def))
}

func TestMergeStringValue(t *testing.T) {
	a := &wrapperspb.StringValue{Value: "a"}
	b := &wrapperspb.StringValue{Value: "b"}

	result := MergeStringValue(nil, nil)
	require.Equal(t, StringValueUnspecified, result)

	result = MergeStringValue(nil, b)
	require.Equal(t, "b", result.Value)

	result = MergeStringValue(a, nil)
	require.Equal(t, "a", result.Value)

	result = MergeStringValue(a, b)
	require.Equal(t, "b", result.Value)
}

func TestStringStringValue(t *testing.T) {
	require.Equal(t, "StringValue{Unspecified}", StringStringValue(nil))
	require.Equal(t, "StringValue{Unspecified}", StringStringValue(StringValueUnspecified))
	require.Equal(t, `StringValue{"hello"}`, StringStringValue(&wrapperspb.StringValue{Value: "hello"}))
}

func TestCoalesceStringValue(t *testing.T) {
	def := &wrapperspb.StringValue{Value: ""}
	ptr := &wrapperspb.StringValue{Value: "a"}

	require.Equal(t, def, CoalesceStringValue(nil, def))
	require.Equal(t, ptr, CoalesceStringValue(ptr, def))
}

func TestSameStringValue(t *testing.T) {
	ptr := &wrapperspb.StringValue{}
	require.True(t, SameStringValue(nil, nil))
	require.True(t, SameStringValue(nil, StringValueUnspecified))
	require.True(t, SameStringValue(StringValueUnspecified, nil))
	require.True(t, SameStringValue(ptr, ptr))
	require.False(t, SameStringValue(nil, ptr))
	require.False(t, SameStringValue(ptr, nil))
	require.False(t, SameStringValue(ptr, StringValueUnspecified))
}

func TestSemanticEqualStringValue(t *testing.T) {
	tests := []struct {
		name string
		a    *wrapperspb.StringValue
		b    *wrapperspb.StringValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, StringValueUnspecified, true},
		{"both equal values", &wrapperspb.StringValue{Value: "a"}, &wrapperspb.StringValue{Value: "a"}, true},
		{"different values", &wrapperspb.StringValue{Value: "a"}, &wrapperspb.StringValue{Value: "b"}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, SemanticEqualStringValue(tc.a, tc.b))
		})
	}
}

func TestEqualStringValue(t *testing.T) {
	require.True(t, EqualStringValue(nil, nil))
	require.True(t, EqualStringValue(StringValueUnspecified, StringValueUnspecified))

	ptr := &wrapperspb.StringValue{Value: "a"}
	require.True(t, EqualStringValue(ptr, ptr))

	ptr2 := &wrapperspb.StringValue{Value: "a"}
	require.True(t, EqualStringValue(ptr, ptr2))

	ptr3 := &wrapperspb.StringValue{Value: "b"}
	require.False(t, EqualStringValue(ptr, ptr3))
}

func TestCopyStringValue(t *testing.T) {
	require.Equal(t, StringValueUnspecified, CopyStringValue(nil))
	require.Equal(t, StringValueUnspecified, CopyStringValue(StringValueUnspecified))

	original := &wrapperspb.StringValue{Value: "hello"}
	copied := CopyStringValue(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied)
}
