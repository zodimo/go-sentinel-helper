package boolutils

import "fmt"

// booleanValue is the internal enum for BooleanValue states
type booleanValue int

const (
	booleanValueUnspecified booleanValue = iota
	booleanValueTrue
	booleanValueFalse
)

// BooleanValue is a type-safe wrapper for tri-state boolean (true/false/unspecified).
// Uses Pattern 1-D from sentinel_pattern.md for zero-allocation value semantics.
type BooleanValue struct {
	value booleanValue
}

// 1. Sentinel - BooleanValueUnspecified
var BooleanValueUnspecified = BooleanValue{value: booleanValueUnspecified}

// Constructors (the ONLY way to create valid BooleanValues)
func BooleanValueTrue() BooleanValue {
	return BooleanValue{value: booleanValueTrue}
}

func BooleanValueFalse() BooleanValue {
	return BooleanValue{value: booleanValueFalse}
}

// BooleanValueFrom creates a BooleanValue from a bool
func BooleanValueFrom(b bool) BooleanValue {
	if b {
		return BooleanValueTrue()
	}
	return BooleanValueFalse()
}

// 2. IsSpecified - predicate (method on value receiver)
func (bv BooleanValue) IsSpecified() bool {
	return bv.value != booleanValueUnspecified
}

// IsUnspecified - convenience predicate
func (bv BooleanValue) IsUnspecified() bool {
	return bv.value == booleanValueUnspecified
}

// IsTrue - check if value is explicitly true
func (bv BooleanValue) IsTrue() bool {
	return bv.value == booleanValueTrue
}

// IsFalse - check if value is explicitly false
func (bv BooleanValue) IsFalse() bool {
	return bv.value == booleanValueFalse
}

// Bool returns the bool value, defaulting to false if unspecified
func (bv BooleanValue) Bool() bool {
	return bv.value == booleanValueTrue
}

// BoolOrElse returns the bool value, or the default if unspecified
func (bv BooleanValue) BoolOrElse(def bool) bool {
	if bv.IsUnspecified() {
		return def
	}
	return bv.Bool()
}

// 3. TakeOrElse - 2-param fallback (method on value receiver)
func (bv BooleanValue) TakeOrElse(def BooleanValue) BooleanValue {
	if bv.IsSpecified() {
		return bv
	}
	return def
}

// 4. Merge - composition merge (method on value receiver for atomic types)
func (bv BooleanValue) Merge(other BooleanValue) BooleanValue {
	if other.IsSpecified() {
		return other
	}
	return bv
}

// MergeBooleanValue - package-level merge function
func MergeBooleanValue(a, b BooleanValue) BooleanValue {
	return a.Merge(b)
}

// 5. String - stringification (method on value receiver)
func (bv BooleanValue) String() string {
	switch bv.value {
	case booleanValueTrue:
		return "BooleanValue{true}"
	case booleanValueFalse:
		return "BooleanValue{false}"
	default:
		return "BooleanValue{Unspecified}"
	}
}

// StringBooleanValue - package-level string function
func StringBooleanValue(bv BooleanValue) string {
	return bv.String()
}

// 6. Coalesce - N/A for value types (no nil possible)
// Not applicable - BooleanValue is a value type

// 7. Same - identity (for value types, uses == operator)
func (bv BooleanValue) Same(other BooleanValue) bool {
	return bv.value == other.value
}

// SameBooleanValue - package-level same function
func SameBooleanValue(a, b BooleanValue) bool {
	return a.Same(b)
}

// 8. SemanticEqual - semantic equality (for value types, same as Same)
func (bv BooleanValue) SemanticEqual(other BooleanValue) bool {
	return bv.value == other.value
}

// SemanticEqualBooleanValue - package-level semantic equal function
func SemanticEqualBooleanValue(a, b BooleanValue) bool {
	return a.SemanticEqual(b)
}

// 9. Equal - equality check (combines Same and SemanticEqual)
func (bv BooleanValue) Equal(other BooleanValue) bool {
	return bv.value == other.value
}

// EqualBooleanValue - package-level equal function
func EqualBooleanValue(a, b BooleanValue) bool {
	return a.Equal(b)
}

// 10. Copy - identity for immutable value types (just returns the value)
func (bv BooleanValue) Copy() BooleanValue {
	return bv
}

// CopyBooleanValue - package-level copy function
func CopyBooleanValue(bv BooleanValue) BooleanValue {
	return bv.Copy()
}

// Format implements fmt.Formatter for custom formatting
func (bv BooleanValue) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		fmt.Fprint(f, bv.String())
	default:
		fmt.Fprintf(f, "%%!%c(BooleanValue=%s)", verb, bv.String())
	}
}
