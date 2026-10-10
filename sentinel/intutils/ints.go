package intutils

import (
	"fmt"
	"math"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
)

type IntValue = int

// COMPANION_PACKAGE_CONTRACT
// This package exports exactly ONE public value: IntValueCompanion.
// 1. Sentinel  - IntValueUnspecified (math.MinInt; a reserved impossible value)
// 2. IsSpecified
// 3. IsUnspecified
// 4. TakeOrElse
// 5. Merge
// 6. String
// 7. Same
// 8. SemanticEqual
// 9. Equal
// 10. Copy (scalar: identity)
// END_CONTRACT

// 1. Sentinel - intValueUnspecified
// IntValue is the type for sentinel int pattern.
// The sentinel math.MinInt is used, so math.MinInt cannot be a valid value.
const intValueUnspecified IntValue = math.MinInt

// 2. IsSpecified - predicate
func isSpecifiedIntValue(i IntValue) bool {
	return i != intValueUnspecified
}

// 3. IsUnspecified - inverse predicate
func isUnspecifiedIntValue(i IntValue) bool {
	return i == intValueUnspecified
}

// 4. TakeOrElse - 2-param fallback
func takeOrElseIntValue(a, b IntValue) IntValue {
	if a != intValueUnspecified {
		return a
	}
	return b
}

// 5. Merge - composition merge (prefers incoming specified values)
func mergeIntValue(a, b IntValue) IntValue {
	if b != intValueUnspecified {
		return b
	}
	return a
}

// 6. String - stringification
func stringIntValue(i IntValue) string {
	if i == intValueUnspecified {
		return "IntValue{Unspecified}"
	}
	return fmt.Sprintf("IntValue{%d}", i)
}

// 7. Same - identity
func sameIntValue(a, b IntValue) bool {
	return a == b
}

// 8. SemanticEqual - semantic equality (same as Same for ints)
func semanticEqualIntValue(a, b IntValue) bool {
	return a == b
}

// 9. Equal - equality check
func equalIntValue(a, b IntValue) bool {
	return a == b
}

// 10. Copy - identity for immutable value types
func copyIntValue(i IntValue) IntValue {
	return i
}

// IntValueCompanion is the ONLY public symbol of this package.
var IntValueCompanion = companions.NewScalarCompanion(
	func() IntValue { return intValueUnspecified },
	stringIntValue,
	isSpecifiedIntValue,
	isUnspecifiedIntValue,
	takeOrElseIntValue,
	mergeIntValue,
	sameIntValue,
	semanticEqualIntValue,
	equalIntValue,
	copyIntValue,
)
