package stringutils

import (
	"fmt"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
)

type StringValue = string

// 1. Sentinel - StringValueUnspecified
// StringValue is the type for sentinel string pattern.
// The sentinel "\x00unspecified" is used when empty string is meaningful.
const StringValueUnspecified StringValue = "\x00unspecified"

// 2. IsSpecified - predicate (package-level function)
func isSpecifiedString(s StringValue) bool {
	return s != StringValueUnspecified
}

// isUnspecifiedString - convenience predicate
func isUnspecifiedString(s StringValue) bool {
	return s == StringValueUnspecified
}

// 3. TakeOrElse - 2-param fallback (package-level function)
func takeOrElseString(a, b StringValue) StringValue {
	if a != StringValueUnspecified {
		return a
	}
	return b
}

// 4. Merge - composition merge (package-level function)
// Prefers incoming specified values over current values
func mergeString(a, b StringValue) StringValue {
	if b != StringValueUnspecified {
		return b
	}
	return a
}

// 5. String - stringification (package-level function)
func stringString(s StringValue) string {
	if s == StringValueUnspecified {
		return "StringValue{Unspecified}"
	}
	return fmt.Sprintf("StringValue{%q}", s)
}

// 6. Coalesce - N/A for string type (string is a value type, no nil possible)
// Not applicable

// 7. Same - identity (package-level function)
func sameString(a, b StringValue) bool {
	return a == b
}

// 8. SemanticEqual - semantic equality (package-level function)
// For strings, this is the same as Same
func semanticEqualString(a, b StringValue) bool {
	return a == b
}

// 9. Equal - equality check (package-level function)
func equalString(a, b StringValue) bool {
	return a == b
}

// 10. Copy - identity for immutable value types (package-level function)
func copyString(s StringValue) StringValue {
	return s
}

var StringsCompanion = companions.NewScalarCompanion(
	func() StringValue { return StringValueUnspecified },
	stringString,
	isSpecifiedString,
	isUnspecifiedString,
	takeOrElseString,
	mergeString,
	sameString,
	semanticEqualString,
	equalString,
	copyString,
)
