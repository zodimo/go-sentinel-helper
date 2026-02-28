package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedDoubleValue(t *testing.T) {
	require.False(t, IsSpecifiedDoubleValue(nil))
	require.False(t, IsSpecifiedDoubleValue(DoubleValueUnspecified))
	require.True(t, IsSpecifiedDoubleValue(&wrapperspb.DoubleValue{Value: 1.5}))
	require.True(t, IsSpecifiedDoubleValue(&wrapperspb.DoubleValue{Value: 0}))
}

func TestTakeOrElseDoubleValue(t *testing.T) {
	def := &wrapperspb.DoubleValue{Value: 3.14}
	ptr := &wrapperspb.DoubleValue{Value: 2.71}

	require.Equal(t, def, TakeOrElseDoubleValue(nil, def))
	require.Equal(t, def, TakeOrElseDoubleValue(DoubleValueUnspecified, def))
	require.Equal(t, ptr, TakeOrElseDoubleValue(ptr, def))
}

func TestMergeDoubleValue(t *testing.T) {
	a := &wrapperspb.DoubleValue{Value: 1.0}
	b := &wrapperspb.DoubleValue{Value: 2.0}

	result := MergeDoubleValue(nil, nil)
	require.Equal(t, DoubleValueUnspecified, result)

	result = MergeDoubleValue(nil, b)
	require.Equal(t, 2.0, result.Value)

	result = MergeDoubleValue(a, nil)
	require.Equal(t, 1.0, result.Value)

	result = MergeDoubleValue(a, b)
	require.Equal(t, 2.0, result.Value)
}

func TestStringDoubleValue(t *testing.T) {
	require.Equal(t, "DoubleValue{Unspecified}", StringDoubleValue(nil))
	require.Equal(t, "DoubleValue{Unspecified}", StringDoubleValue(DoubleValueUnspecified))
	require.Equal(t, "DoubleValue{3.14}", StringDoubleValue(&wrapperspb.DoubleValue{Value: 3.14}))
}

func TestCoalesceDoubleValue(t *testing.T) {
	def := &wrapperspb.DoubleValue{Value: 0}
	ptr := &wrapperspb.DoubleValue{Value: 1.5}

	require.Equal(t, def, CoalesceDoubleValue(nil, def))
	require.Equal(t, ptr, CoalesceDoubleValue(ptr, def))
}

func TestSameDoubleValue(t *testing.T) {
	ptr := &wrapperspb.DoubleValue{}
	require.True(t, SameDoubleValue(nil, nil))
	require.True(t, SameDoubleValue(nil, DoubleValueUnspecified))
	require.True(t, SameDoubleValue(DoubleValueUnspecified, nil))
	require.True(t, SameDoubleValue(ptr, ptr))
	require.False(t, SameDoubleValue(nil, ptr))
	require.False(t, SameDoubleValue(ptr, nil))
	require.False(t, SameDoubleValue(ptr, DoubleValueUnspecified))
}

func TestSemanticEqualDoubleValue(t *testing.T) {
	tests := []struct {
		name string
		a    *wrapperspb.DoubleValue
		b    *wrapperspb.DoubleValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, DoubleValueUnspecified, true},
		{"both equal values", &wrapperspb.DoubleValue{Value: 1.5}, &wrapperspb.DoubleValue{Value: 1.5}, true},
		{"different values", &wrapperspb.DoubleValue{Value: 1.5}, &wrapperspb.DoubleValue{Value: 2.5}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, SemanticEqualDoubleValue(tc.a, tc.b))
		})
	}
}

func TestEqualDoubleValue(t *testing.T) {
	require.True(t, EqualDoubleValue(nil, nil))
	require.True(t, EqualDoubleValue(DoubleValueUnspecified, DoubleValueUnspecified))

	ptr := &wrapperspb.DoubleValue{Value: 1.5}
	require.True(t, EqualDoubleValue(ptr, ptr))

	ptr2 := &wrapperspb.DoubleValue{Value: 1.5}
	require.True(t, EqualDoubleValue(ptr, ptr2))

	ptr3 := &wrapperspb.DoubleValue{Value: 2.5}
	require.False(t, EqualDoubleValue(ptr, ptr3))
}

func TestCopyDoubleValue(t *testing.T) {
	require.Equal(t, DoubleValueUnspecified, CopyDoubleValue(nil))
	require.Equal(t, DoubleValueUnspecified, CopyDoubleValue(DoubleValueUnspecified))

	original := &wrapperspb.DoubleValue{Value: 3.14}
	copied := CopyDoubleValue(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied)
}
