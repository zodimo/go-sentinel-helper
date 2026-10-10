package boolutils

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
)

// booleanTriValue is the internal enum for BooleanValue states
type booleanTriValue int

const (
	booleanTriValueUnspecified booleanTriValue = iota
	booleanTriValueTrue
	booleanTriValueFalse
)

// BooleanValue is a type-safe wrapper for tri-state boolean (true/false/unspecified).
// Uses Pattern 2-D from sentinel_pattern.md for zero-allocation value semantics.
type BooleanValue struct {
	value booleanTriValue
}

// 1. Sentinel
var booleanValueUnspecified = BooleanValue{value: booleanTriValueUnspecified}

// Constructors (the ONLY way to create valid BooleanValues)
func BooleanValueTrue() BooleanValue {
	return BooleanValue{value: booleanTriValueTrue}
}

func BooleanValueFalse() BooleanValue {
	return BooleanValue{value: booleanTriValueFalse}
}

// BooleanValueFrom creates a BooleanValue from a bool
func BooleanValueFrom(b bool) BooleanValue {
	if b {
		return BooleanValueTrue()
	}
	return BooleanValueFalse()
}

// --- Canonical operations (unexported; registered on the companion) ---

// 1. Unspecified
func unspecifiedBooleanValue() BooleanValue {
	return booleanValueUnspecified
}

// 2. IsSpecified
func isSpecifiedBooleanValue(bv BooleanValue) bool {
	return bv.value != booleanTriValueUnspecified
}

// 3. IsUnspecified
func isUnspecifiedBooleanValue(bv BooleanValue) bool {
	return bv.value == booleanTriValueUnspecified
}

// 4. TakeOrElse
func takeOrElseBooleanValue(bv, def BooleanValue) BooleanValue {
	if bv.value != booleanTriValueUnspecified {
		return bv
	}
	return def
}

// 5. Merge
func mergeBooleanValue(a, b BooleanValue) BooleanValue {
	if b.value != booleanTriValueUnspecified {
		return b
	}
	return a
}

// 6. String
func stringBooleanValue(bv BooleanValue) string {
	switch bv.value {
	case booleanTriValueTrue:
		return "BooleanValue{true}"
	case booleanTriValueFalse:
		return "BooleanValue{false}"
	default:
		return "BooleanValue{Unspecified}"
	}
}

// 7. Same
func sameBooleanValue(a, b BooleanValue) bool {
	return a.value == b.value
}

// 8. SemanticEqual
func semanticEqualBooleanValue(a, b BooleanValue) bool {
	return a.value == b.value
}

// 9. Equal
func equalBooleanValue(a, b BooleanValue) bool {
	return a.value == b.value
}

// 10. Copy
func copyBooleanValue(bv BooleanValue) BooleanValue {
	return bv
}

// BooleanValueCompanion is the ONLY public companion of this package.
var BooleanValueCompanion = companions.NewScalarCompanion(
	unspecifiedBooleanValue,
	stringBooleanValue,
	isSpecifiedBooleanValue,
	isUnspecifiedBooleanValue,
	takeOrElseBooleanValue,
	mergeBooleanValue,
	sameBooleanValue,
	semanticEqualBooleanValue,
	equalBooleanValue,
	copyBooleanValue,
)

// --- Domain-specific convenience methods (beyond the companion) ---

// IsSpecified reports whether the value is neither true nor false but has been set.
func (bv BooleanValue) IsSpecified() bool { return isSpecifiedBooleanValue(bv) }

// IsUnspecified reports whether the value is the tri-state "unknown".
func (bv BooleanValue) IsUnspecified() bool { return isUnspecifiedBooleanValue(bv) }

// IsTrue - check if value is explicitly true
func (bv BooleanValue) IsTrue() bool { return bv.value == booleanTriValueTrue }

// IsFalse - check if value is explicitly false
func (bv BooleanValue) IsFalse() bool { return bv.value == booleanTriValueFalse }

// Bool returns the bool value, defaulting to false if unspecified
func (bv BooleanValue) Bool() bool { return bv.value == booleanTriValueTrue }

// BoolOrElse returns the bool value, or the default if unspecified
func (bv BooleanValue) BoolOrElse(def bool) bool {
	if bv.IsUnspecified() {
		return def
	}
	return bv.Bool()
}

// String implements fmt.Stringer.
func (bv BooleanValue) String() string { return stringBooleanValue(bv) }

// Format implements fmt.Formatter for custom formatting
func (bv BooleanValue) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		fmt.Fprint(f, bv.String())
	default:
		fmt.Fprintf(f, "%%!%c(BooleanValue=%s)", verb, bv.String())
	}
}
