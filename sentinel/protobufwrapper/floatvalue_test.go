package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedFloatValue(t *testing.T) {
	require.False(t, IsSpecifiedFloatValue(nil))
	require.False(t, IsSpecifiedFloatValue(FloatValueUnspecified))
	require.True(t, IsSpecifiedFloatValue(&wrapperspb.FloatValue{Value: 1.5}))
	require.True(t, IsSpecifiedFloatValue(&wrapperspb.FloatValue{Value: 0}))
}

func TestTakeOrElseFloatValue(t *testing.T) {
	def := &wrapperspb.FloatValue{Value: 3.14}
	ptr := &wrapperspb.FloatValue{Value: 2.71}

	require.Equal(t, def, TakeOrElseFloatValue(nil, def))
	require.Equal(t, def, TakeOrElseFloatValue(FloatValueUnspecified, def))
	require.Equal(t, ptr, TakeOrElseFloatValue(ptr, def))
}

func TestMergeFloatValue(t *testing.T) {
	a := &wrapperspb.FloatValue{Value: 1.0}
	b := &wrapperspb.FloatValue{Value: 2.0}

	result := MergeFloatValue(nil, nil)
	require.Equal(t, FloatValueUnspecified, result)

	result = MergeFloatValue(nil, b)
	require.Equal(t, float32(2.0), result.Value)

	result = MergeFloatValue(a, nil)
	require.Equal(t, float32(1.0), result.Value)

	result = MergeFloatValue(a, b)
	require.Equal(t, float32(2.0), result.Value)
}

func TestStringFloatValue(t *testing.T) {
	require.Equal(t, "FloatValue{Unspecified}", StringFloatValue(nil))
	require.Equal(t, "FloatValue{Unspecified}", StringFloatValue(FloatValueUnspecified))
	require.Equal(t, "FloatValue{3.14}", StringFloatValue(&wrapperspb.FloatValue{Value: 3.14}))
}

func TestCoalesceFloatValue(t *testing.T) {
	def := &wrapperspb.FloatValue{Value: 0}
	ptr := &wrapperspb.FloatValue{Value: 1.5}

	require.Equal(t, def, CoalesceFloatValue(nil, def))
	require.Equal(t, ptr, CoalesceFloatValue(ptr, def))
}

func TestSameFloatValue(t *testing.T) {
	ptr := &wrapperspb.FloatValue{}
	require.True(t, SameFloatValue(nil, nil))
	require.True(t, SameFloatValue(nil, FloatValueUnspecified))
	require.True(t, SameFloatValue(FloatValueUnspecified, nil))
	require.True(t, SameFloatValue(ptr, ptr))
	require.False(t, SameFloatValue(nil, ptr))
	require.False(t, SameFloatValue(ptr, nil))
	require.False(t, SameFloatValue(ptr, FloatValueUnspecified))
}

func TestSemanticEqualFloatValue(t *testing.T) {
	tests := []struct {
		name string
		a    *wrapperspb.FloatValue
		b    *wrapperspb.FloatValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, FloatValueUnspecified, true},
		{"both equal values", &wrapperspb.FloatValue{Value: 1.5}, &wrapperspb.FloatValue{Value: 1.5}, true},
		{"different values", &wrapperspb.FloatValue{Value: 1.5}, &wrapperspb.FloatValue{Value: 2.5}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, SemanticEqualFloatValue(tc.a, tc.b))
		})
	}
}

func TestEqualFloatValue(t *testing.T) {
	require.True(t, EqualFloatValue(nil, nil))
	require.True(t, EqualFloatValue(FloatValueUnspecified, FloatValueUnspecified))

	ptr := &wrapperspb.FloatValue{Value: 1.5}
	require.True(t, EqualFloatValue(ptr, ptr))

	ptr2 := &wrapperspb.FloatValue{Value: 1.5}
	require.True(t, EqualFloatValue(ptr, ptr2))

	ptr3 := &wrapperspb.FloatValue{Value: 2.5}
	require.False(t, EqualFloatValue(ptr, ptr3))
}

func TestCopyFloatValue(t *testing.T) {
	require.Equal(t, FloatValueUnspecified, CopyFloatValue(nil))
	require.Equal(t, FloatValueUnspecified, CopyFloatValue(FloatValueUnspecified))

	original := &wrapperspb.FloatValue{Value: 3.14}
	copied := CopyFloatValue(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied)
}
