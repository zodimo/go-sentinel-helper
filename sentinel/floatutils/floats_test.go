package floatutils

import (
	"math"
	"testing"
)

func TestMerge(t *testing.T) {
	// Float32
	if got := Merge(float32(1.0), float32(2.0)); got != 2.0 {
		t.Errorf("Merge(float32) preferred value mismatch: got %v, want 2.0", got)
	}
	if got := Merge(float32(1.0), Float32Unspecified); got != 1.0 {
		t.Errorf("Merge(float32) with unspecified mismatch: got %v, want 1.0", got)
	}
	if got := Merge(Float32Unspecified, float32(2.0)); got != 2.0 {
		t.Errorf("Merge(float32) with first unspecified mismatch: got %v, want 2.0", got)
	}

	// Float64
	if got := Merge(1.0, 2.0); got != 2.0 {
		t.Errorf("Merge(float64) preferred value mismatch: got %v, want 2.0", got)
	}
	if got := Merge(1.0, Float64Unspecified); got != 1.0 {
		t.Errorf("Merge(float64) with unspecified mismatch: got %v, want 1.0", got)
	}
}

func TestString(t *testing.T) {
	// Float32
	if got := String(Float32Unspecified); got != "float32{Unspecified}" {
		t.Errorf("String(float32 unspecified) = %q, want float32{Unspecified}", got)
	}
	if got := String(float32(1.23)); got != "float32{1.23}" {
		t.Errorf("String(float32) = %q, want float32{1.23}", got)
	}

	// Float64
	if got := String(Float64Unspecified); got != "float64{Unspecified}" {
		t.Errorf("String(float64 unspecified) = %q, want float64{Unspecified}", got)
	}
	if got := String(1.23); got != "float64{1.23}" {
		t.Errorf("String(float64) = %q, want float64{1.23}", got)
	}
}

func TestSame(t *testing.T) {
	// Float32
	if !Same(float32(1.0), float32(1.0)) {
		t.Error("Same(float32 1.0, 1.0) should be true")
	}
	if Same(float32(1.0), float32(2.0)) {
		t.Error("Same(float32 1.0, 2.0) should be false")
	}
	if !Same(Float32Unspecified, Float32Unspecified) {
		t.Error("Same(Float32Unspecified, Float32Unspecified) should be true")
	}

	// Float64
	if !Same(1.0, 1.0) {
		t.Error("Same(float64 1.0, 1.0) should be true")
	}
	if !Same(Float64Unspecified, Float64Unspecified) {
		t.Error("Same(Float64Unspecified, Float64Unspecified) should be true")
	}
}

func TestSemanticEqual(t *testing.T) {
	// Float32
	if !SemanticEqual(float32(1.0), float32(1.0)) {
		t.Error("SemanticEqual(float32 1.0, 1.0) should be true")
	}
	// Test epsilon
	// 1e-6 is threshold. 1.0 + 0.5e-6 should be equal
	if !SemanticEqual(float32(1.0), float32(1.0000005)) {
		t.Error("SemanticEqual(float32) within epsilon should be true")
	}
	// 1.0 + 2e-6 should be not equal (threshold is 1e-6)
	if SemanticEqual(float32(1.0), float32(1.000002)) {
		t.Error("SemanticEqual(float32) outside epsilon should be false")
	}

	if !SemanticEqual(Float32Unspecified, Float32Unspecified) {
		t.Error("SemanticEqual(Float32Unspecified, Float32Unspecified) should be true")
	}
	if SemanticEqual(float32(1.0), Float32Unspecified) {
		t.Error("SemanticEqual(specified, unspecified) should be false")
	}

	// Float64
	if !SemanticEqual(1.0, 1.0) {
		t.Error("SemanticEqual(float64 1.0, 1.0) should be true")
	}
	// 1e-9 is threshold
	if !SemanticEqual(1.0, 1.0000000005) {
		t.Error("SemanticEqual(float64) within epsilon should be true")
	}
	if SemanticEqual(1.0, 1.000000002) {
		t.Error("SemanticEqual(float64) outside epsilon should be false")
	}
}

func TestEqual(t *testing.T) {
	// Delegates to SemanticEqual
	if !Equal(1.0, 1.0) {
		t.Error("Equal should return true for same values")
	}
}

func TestCopy(t *testing.T) {
	v := 1.23
	if got := Copy(v); got != v {
		t.Errorf("Copy(%v) = %v, want %v", v, got, v)
	}
}

func TestInfinite(t *testing.T) {
	if !IsInfinite(Float64Value(math.Inf(1))) {
		t.Error("IsInfinite(Inf) should be true")
	}
	if IsInfinite(1.0) {
		t.Error("IsInfinite(1.0) should be false")
	}
	if IsInfinite(Float64Unspecified) {
		t.Error("IsInfinite(NaN) should be false")
	}
}

// Type aliases for easier testing of "Value" concept if needed,
// though we are testing raw primitives now as per generic implementation.
type Float64Value = float64
