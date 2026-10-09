package boolutils

import (
	"fmt"
	"testing"
)

var c = BooleanValueCompanion

func TestBooleanValue_Constructors(t *testing.T) {
	tests := []struct {
		name     string
		got      BooleanValue
		wantBool bool
		wantSpec bool
	}{
		{"Unspecified", BooleanValueUnspecified, false, false},
		{"True", BooleanValueTrue(), true, true},
		{"False", BooleanValueFalse(), false, true},
		{"From(true)", BooleanValueFrom(true), true, true},
		{"From(false)", BooleanValueFrom(false), false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsSpecified() != tt.wantSpec {
				t.Errorf("IsSpecified() = %v, want %v", tt.got.IsSpecified(), tt.wantSpec)
			}
			if tt.got.IsUnspecified() == tt.wantSpec {
				t.Errorf("IsUnspecified() = %v, want %v", tt.got.IsUnspecified(), !tt.wantSpec)
			}
			// Bool() defaults to false for Unspecified
			if tt.got.Bool() != tt.wantBool {
				t.Errorf("Bool() = %v, want %v", tt.got.Bool(), tt.wantBool)
			}
			// Companion predicates agree with the methods.
			if c.IsSpecified(tt.got) != tt.wantSpec {
				t.Errorf("companion IsSpecified() = %v, want %v", c.IsSpecified(tt.got), tt.wantSpec)
			}
		})
	}
}

func TestBooleanValue_Predicates(t *testing.T) {
	vTrue := BooleanValueTrue()
	vFalse := BooleanValueFalse()
	vUnspec := BooleanValueUnspecified

	if !vTrue.IsTrue() {
		t.Error("IsTrue() failed for True value")
	}
	if vTrue.IsFalse() {
		t.Error("IsFalse() failed for True value")
	}

	if !vFalse.IsFalse() {
		t.Error("IsFalse() failed for False value")
	}
	if vFalse.IsTrue() {
		t.Error("IsTrue() failed for False value")
	}

	if vUnspec.IsTrue() {
		t.Error("IsTrue() failed for Unspecified value")
	}
	if vUnspec.IsFalse() {
		t.Error("IsFalse() failed for Unspecified value")
	}
}

func TestBooleanValue_BoolOrElse(t *testing.T) {
	tests := []struct {
		val  BooleanValue
		def  bool
		want bool
	}{
		{BooleanValueTrue(), false, true},
		{BooleanValueTrue(), true, true},
		{BooleanValueFalse(), false, false},
		{BooleanValueFalse(), true, false},
		{BooleanValueUnspecified, true, true},
		{BooleanValueUnspecified, false, false},
	}

	for _, tt := range tests {
		if got := tt.val.BoolOrElse(tt.def); got != tt.want {
			t.Errorf("BoolOrElse(%v) = %v, want %v", tt.def, got, tt.want)
		}
	}
}

func TestBooleanValue_TakeOrElse(t *testing.T) {
	vTrue := BooleanValueTrue()
	vFalse := BooleanValueFalse()
	vMaybe := BooleanValueUnspecified

	// If specified, should return itself
	if got := c.TakeOrElse(vTrue, vFalse); !c.Equal(got, vTrue) {
		t.Errorf("TakeOrElse failed: got %v, want %v", got, vTrue)
	}
	if got := c.TakeOrElse(vFalse, vTrue); !c.Equal(got, vFalse) {
		t.Errorf("TakeOrElse failed: got %v, want %v", got, vFalse)
	}

	// If unspecified, should return default
	if got := c.TakeOrElse(vMaybe, vTrue); !c.Equal(got, vTrue) {
		t.Errorf("TakeOrElse failed: got %v, want %v", got, vTrue)
	}
}

func TestBooleanValue_Merge(t *testing.T) {
	vTrue := BooleanValueTrue()
	vFalse := BooleanValueFalse()
	vUnspec := BooleanValueUnspecified

	// Merge(current, incoming) -> incoming if specified, else current

	// Case 1: Incoming is specified (overrides current)
	if got := c.Merge(vUnspec, vTrue); !c.Equal(got, vTrue) {
		t.Errorf("Merge(Unspec, True) = %v, want True", got)
	}
	if got := c.Merge(vFalse, vTrue); !c.Equal(got, vTrue) {
		t.Errorf("Merge(False, True) = %v, want True", got)
	}

	// Case 2: Incoming is unspecified (keeps current)
	if got := c.Merge(vTrue, vUnspec); !c.Equal(got, vTrue) {
		t.Errorf("Merge(True, Unspec) = %v, want True", got)
	}
	if got := c.Merge(vUnspec, vUnspec); !c.Equal(got, vUnspec) {
		t.Errorf("Merge(Unspec, Unspec) = %v, want Unspec", got)
	}
}

func TestBooleanValue_String(t *testing.T) {
	tests := []struct {
		val  BooleanValue
		want string
	}{
		{BooleanValueTrue(), "BooleanValue{true}"},
		{BooleanValueFalse(), "BooleanValue{false}"},
		{BooleanValueUnspecified, "BooleanValue{Unspecified}"},
	}

	for _, tt := range tests {
		if got := tt.val.String(); got != tt.want {
			t.Errorf("String() = %q, want %q", got, tt.want)
		}
		// Companion String agrees with the method.
		if got := c.String(tt.val); got != tt.want {
			t.Errorf("companion String() = %q, want %q", got, tt.want)
		}
	}
}

func TestBooleanValue_Equality(t *testing.T) {
	t1 := BooleanValueTrue()
	t2 := BooleanValueTrue()
	f1 := BooleanValueFalse()
	u1 := BooleanValueUnspecified
	u2 := BooleanValueUnspecified

	if !c.Equal(t1, t2) {
		t.Error("True should equal True")
	}
	if !c.Equal(u1, u2) {
		t.Error("Unspecified should equal Unspecified")
	}
	if c.Equal(t1, f1) {
		t.Error("True should not equal False")
	}
	if c.Equal(t1, u1) {
		t.Error("True should not equal Unspecified")
	}

	if !c.Same(t1, t2) {
		t.Error("Same failed")
	}
	if !c.Equal(t1, t2) {
		t.Error("Equal failed")
	}
	if !c.SemanticEqual(t1, t2) {
		t.Error("SemanticEqual failed")
	}
}

func TestBooleanValue_Copy(t *testing.T) {
	v := BooleanValueTrue()
	dup := c.Copy(v)
	if !c.Equal(v, dup) {
		t.Error("Copy failed identity check")
	}
}

func TestBooleanValue_Format(t *testing.T) {
	v := BooleanValueTrue()
	got := fmt.Sprintf("%v", v)
	want := "BooleanValue{true}"
	if got != want {
		t.Errorf("Format %%v = %q, want %q", got, want)
	}

	got = fmt.Sprintf("%s", v)
	if got != want {
		t.Errorf("Format %%s = %q, want %q", got, want)
	}
}
