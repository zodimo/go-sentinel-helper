package floatutils

import (
	"math"
	"testing"
)

func TestMerge(t *testing.T) {
	c32 := Float32Companion
	c64 := Float64Companion

	// Float32
	if got := c32.Merge(float32(1.0), float32(2.0)); got != 2.0 {
		t.Errorf("Merge(float32) preferred value mismatch: got %v, want 2.0", got)
	}
	if got := c32.Merge(float32(1.0), float32Unspecified); got != 1.0 {
		t.Errorf("Merge(float32) with unspecified mismatch: got %v, want 1.0", got)
	}
	if got := c32.Merge(float32Unspecified, float32(2.0)); got != 2.0 {
		t.Errorf("Merge(float32) with first unspecified mismatch: got %v, want 2.0", got)
	}

	// Float64
	if got := c64.Merge(1.0, 2.0); got != 2.0 {
		t.Errorf("Merge(float64) preferred value mismatch: got %v, want 2.0", got)
	}
	if got := c64.Merge(1.0, float64Unspecified); got != 1.0 {
		t.Errorf("Merge(float64) with unspecified mismatch: got %v, want 1.0", got)
	}
}

func TestString(t *testing.T) {
	c32 := Float32Companion
	c64 := Float64Companion

	// Float32
	if got := c32.String(float32Unspecified); got != "float32{Unspecified}" {
		t.Errorf("String(float32 unspecified) = %q, want float32{Unspecified}", got)
	}
	if got := c32.String(float32(1.23)); got != "float32{1.23}" {
		t.Errorf("String(float32) = %q, want float32{1.23}", got)
	}

	// Float64
	if got := c64.String(float64Unspecified); got != "float64{Unspecified}" {
		t.Errorf("String(float64 unspecified) = %q, want float64{Unspecified}", got)
	}
	if got := c64.String(1.23); got != "float64{1.23}" {
		t.Errorf("String(float64) = %q, want float64{1.23}", got)
	}
}

func TestSame(t *testing.T) {
	c32 := Float32Companion
	c64 := Float64Companion

	// Float32
	if !c32.Same(float32(1.0), float32(1.0)) {
		t.Error("Same(float32 1.0, 1.0) should be true")
	}
	if c32.Same(float32(1.0), float32(2.0)) {
		t.Error("Same(float32 1.0, 2.0) should be false")
	}
	if !c32.Same(float32Unspecified, float32Unspecified) {
		t.Error("Same(Float32Unspecified, Float32Unspecified) should be true")
	}

	// Float64
	if !c64.Same(1.0, 1.0) {
		t.Error("Same(float64 1.0, 1.0) should be true")
	}
	if !c64.Same(float64Unspecified, float64Unspecified) {
		t.Error("Same(Float64Unspecified, Float64Unspecified) should be true")
	}
}

func TestSemanticEqual(t *testing.T) {
	c32 := Float32Companion
	c64 := Float64Companion

	// Float32
	if !c32.SemanticEqual(float32(1.0), float32(1.0)) {
		t.Error("SemanticEqual(float32 1.0, 1.0) should be true")
	}
	// 1e-6 is threshold. 1.0 + 0.5e-6 should be equal
	if !c32.SemanticEqual(float32(1.0), float32(1.0000005)) {
		t.Error("SemanticEqual(float32) within epsilon should be true")
	}
	// 1.0 + 2e-6 should be not equal (threshold is 1e-6)
	if c32.SemanticEqual(float32(1.0), float32(1.000002)) {
		t.Error("SemanticEqual(float32) outside epsilon should be false")
	}

	if !c32.SemanticEqual(float32Unspecified, float32Unspecified) {
		t.Error("SemanticEqual(Float32Unspecified, Float32Unspecified) should be true")
	}
	if c32.SemanticEqual(float32(1.0), float32Unspecified) {
		t.Error("SemanticEqual(specified, unspecified) should be false")
	}

	// Float64
	if !c64.SemanticEqual(1.0, 1.0) {
		t.Error("SemanticEqual(float64 1.0, 1.0) should be true")
	}
	// 1e-9 is threshold
	if !c64.SemanticEqual(1.0, 1.0000000005) {
		t.Error("SemanticEqual(float64) within epsilon should be true")
	}
	if c64.SemanticEqual(1.0, 1.000000002) {
		t.Error("SemanticEqual(float64) outside epsilon should be false")
	}
}

func TestEqual(t *testing.T) {
	// Delegates to SemanticEqual
	if !Float64Companion.Equal(1.0, 1.0) {
		t.Error("Equal should return true for same values")
	}
}

func TestCopy(t *testing.T) {
	v := 1.23
	if got := Float64Companion.Copy(v); got != v {
		t.Errorf("Copy(%v) = %v, want %v", v, got, v)
	}
}

func TestInfinite(t *testing.T) {
	if !IsInfinite(math.Inf(1)) {
		t.Error("IsInfinite(Inf) should be true")
	}
	if IsInfinite(1.0) {
		t.Error("IsInfinite(1.0) should be false")
	}
	if IsInfinite(float64Unspecified) {
		t.Error("IsInfinite(NaN) should be false")
	}
}
