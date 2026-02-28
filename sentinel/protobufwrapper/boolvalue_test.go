package protobufwrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIsSpecifiedBoolValue(t *testing.T) {
	require.False(t, IsSpecifiedBoolValue(nil))
	require.False(t, IsSpecifiedBoolValue(BoolValueUnspecified))
	require.True(t, IsSpecifiedBoolValue(&wrapperspb.BoolValue{Value: true}))
	require.True(t, IsSpecifiedBoolValue(&wrapperspb.BoolValue{Value: false}))
}

func TestTakeOrElseBoolValue(t *testing.T) {
	def := &wrapperspb.BoolValue{Value: true}
	ptr := &wrapperspb.BoolValue{Value: false}

	require.Equal(t, def, TakeOrElseBoolValue(nil, def))
	require.Equal(t, def, TakeOrElseBoolValue(BoolValueUnspecified, def))
	require.Equal(t, ptr, TakeOrElseBoolValue(ptr, def))
}

func TestMergeBoolValue(t *testing.T) {
	a := &wrapperspb.BoolValue{Value: true}
	b := &wrapperspb.BoolValue{Value: false}

	// both unspecified
	result := MergeBoolValue(nil, nil)
	require.Equal(t, BoolValueUnspecified, result)

	// a unspecified, b specified
	result = MergeBoolValue(nil, b)
	require.Equal(t, false, result.Value)

	// a specified, b unspecified
	result = MergeBoolValue(a, nil)
	require.Equal(t, true, result.Value)

	// both specified — b wins
	result = MergeBoolValue(a, b)
	require.Equal(t, false, result.Value)
}

func TestStringBoolValue(t *testing.T) {
	require.Equal(t, "BoolValue{Unspecified}", StringBoolValue(nil))
	require.Equal(t, "BoolValue{Unspecified}", StringBoolValue(BoolValueUnspecified))
	require.Equal(t, "BoolValue{true}", StringBoolValue(&wrapperspb.BoolValue{Value: true}))
	require.Equal(t, "BoolValue{false}", StringBoolValue(&wrapperspb.BoolValue{Value: false}))
}

func TestCoalesceBoolValue(t *testing.T) {
	def := &wrapperspb.BoolValue{Value: false}
	ptr := &wrapperspb.BoolValue{Value: true}

	require.Equal(t, def, CoalesceBoolValue(nil, def))
	require.Equal(t, ptr, CoalesceBoolValue(ptr, def))
}

func TestSameBoolValue(t *testing.T) {
	ptr := &wrapperspb.BoolValue{}
	require.True(t, SameBoolValue(nil, nil))
	require.True(t, SameBoolValue(nil, BoolValueUnspecified))
	require.True(t, SameBoolValue(BoolValueUnspecified, nil))
	require.True(t, SameBoolValue(ptr, ptr))
	require.False(t, SameBoolValue(nil, ptr))
	require.False(t, SameBoolValue(ptr, nil))
	require.False(t, SameBoolValue(ptr, BoolValueUnspecified))
}

func TestSemanticEqualBoolValue(t *testing.T) {
	tests := []struct {
		name string
		a    *wrapperspb.BoolValue
		b    *wrapperspb.BoolValue
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one Unspecified", nil, BoolValueUnspecified, true},
		{"both equal values", &wrapperspb.BoolValue{Value: true}, &wrapperspb.BoolValue{Value: true}, true},
		{"different values", &wrapperspb.BoolValue{Value: true}, &wrapperspb.BoolValue{Value: false}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, SemanticEqualBoolValue(tc.a, tc.b))
		})
	}
}

func TestEqualBoolValue(t *testing.T) {
	require.True(t, EqualBoolValue(nil, nil))
	require.True(t, EqualBoolValue(BoolValueUnspecified, BoolValueUnspecified))

	ptr := &wrapperspb.BoolValue{Value: true}
	require.True(t, EqualBoolValue(ptr, ptr))

	ptr2 := &wrapperspb.BoolValue{Value: true}
	require.True(t, EqualBoolValue(ptr, ptr2))

	ptr3 := &wrapperspb.BoolValue{Value: false}
	require.False(t, EqualBoolValue(ptr, ptr3))
}

func TestCopyBoolValue(t *testing.T) {
	require.Equal(t, BoolValueUnspecified, CopyBoolValue(nil))
	require.Equal(t, BoolValueUnspecified, CopyBoolValue(BoolValueUnspecified))

	original := &wrapperspb.BoolValue{Value: true}
	copied := CopyBoolValue(original)
	require.Equal(t, original.Value, copied.Value)
	require.True(t, original != copied) // different pointer
}
