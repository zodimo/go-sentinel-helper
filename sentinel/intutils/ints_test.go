package intutils

import (
	"testing"
)

func TestIntValue(t *testing.T) {
	c := IntValueCompanion

	if c.IsSpecified(IntValueUnspecified) {
		t.Error("IntValueUnspecified should not be specified")
	}
	if !c.IsUnspecified(IntValueUnspecified) {
		t.Error("IntValueUnspecified should be unspecified")
	}

	val := 42
	if !c.IsSpecified(val) {
		t.Error("42 should be specified")
	}
	if c.IsUnspecified(val) {
		t.Error("42 should not be unspecified")
	}

	// TakeOrElse
	if got := c.TakeOrElse(val, 100); got != val {
		t.Errorf("TakeOrElse(42, 100) = %d; want 42", got)
	}
	if got := c.TakeOrElse(IntValueUnspecified, 100); got != 100 {
		t.Errorf("TakeOrElse(Unspecified, 100) = %d; want 100", got)
	}

	// Merge
	if got := c.Merge(10, 20); got != 20 {
		t.Errorf("Merge(10, 20) = %d; want 20", got)
	}
	if got := c.Merge(10, IntValueUnspecified); got != 10 {
		t.Errorf("Merge(10, Unspecified) = %d; want 10", got)
	}
	if got := c.Merge(IntValueUnspecified, 20); got != 20 {
		t.Errorf("Merge(Unspecified, 20) = %d; want 20", got)
	}

	// String
	if got := c.String(IntValueUnspecified); got != "IntValue{Unspecified}" {
		t.Errorf("String(Unspecified) = %q; want 'IntValue{Unspecified}'", got)
	}
	if got := c.String(42); got != "IntValue{42}" {
		t.Errorf("String(42) = %q; want 'IntValue{42}'", got)
	}

	// Identity/Equals
	if !c.Same(42, 42) {
		t.Error("Same(42, 42) should be true")
	}
	if c.Same(42, 43) {
		t.Error("Same(42, 43) should be false")
	}

	if !c.SemanticEqual(42, 42) {
		t.Error("SemanticEqual(42, 42) should be true")
	}
	if !c.Equal(42, 42) {
		t.Error("Equal(42, 42) should be true")
	}

	if c.Copy(42) != 42 {
		t.Error("Copy(42) should be 42")
	}
}
