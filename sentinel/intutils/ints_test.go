package intutils

import (
	"testing"
)

func TestIntValue(t *testing.T) {
	if IsSpecifiedIntValue(IntValueUnspecified) {
		t.Error("IntValueUnspecified should not be specified")
	}
	if !IsUnspecifiedIntValue(IntValueUnspecified) {
		t.Error("IntValueUnspecified should be unspecified")
	}

	val := 42
	if !IsSpecifiedIntValue(val) {
		t.Error("42 should be specified")
	}
	if IsUnspecifiedIntValue(val) {
		t.Error("42 should not be unspecified")
	}

	// TakeOrElse
	if got := TakeOrElseIntValue(val, 100); got != val {
		t.Errorf("TakeOrElse(42, 100) = %d; want 42", got)
	}
	if got := TakeOrElseIntValue(IntValueUnspecified, 100); got != 100 {
		t.Errorf("TakeOrElse(Unspecified, 100) = %d; want 100", got)
	}

	// Merge
	if got := MergeIntValue(10, 20); got != 20 {
		t.Errorf("Merge(10, 20) = %d; want 20", got)
	}
	if got := MergeIntValue(10, IntValueUnspecified); got != 10 {
		t.Errorf("Merge(10, Unspecified) = %d; want 10", got)
	}
	if got := MergeIntValue(IntValueUnspecified, 20); got != 20 {
		t.Errorf("Merge(Unspecified, 20) = %d; want 20", got)
	}

	// String
	if got := StringIntValue(IntValueUnspecified); got != "IntValue{Unspecified}" {
		t.Errorf("StringIntValue(Unspecified) = %q; want 'IntValue{Unspecified}'", got)
	}
	if got := StringIntValue(42); got != "IntValue{42}" {
		t.Errorf("StringIntValue(42) = %q; want 'IntValue{42}'", got)
	}

	// Identity/Equals
	if !SameIntValue(42, 42) {
		t.Error("Same(42, 42) should be true")
	}
	if SameIntValue(42, 43) {
		t.Error("Same(42, 43) should be false")
	}

	if !SemanticEqualIntValue(42, 42) {
		t.Error("SemanticEqual(42, 42) should be true")
	}
	if !EqualIntValue(42, 42) {
		t.Error("Equal(42, 42) should be true")
	}

	if CopyIntValue(42) != 42 {
		t.Error("Copy(42) should be 42")
	}
}
