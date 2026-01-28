package floatutils

import (
	"fmt"
	"math"
)

const (
	Float64EqualityThreshold float64 = 1e-9
	Float32EqualityThreshold float32 = 1e-6
)

var Float64Unspecified = math.NaN()
var Float32Unspecified = float32(math.NaN())

var Float32Infinite = float32(math.Inf(1))
var FloatInfinite = math.Inf(1)

type Float interface {
	~float32 | ~float64
}

func TakeOrElse[T Float](v T, defaultValue T) T {
	if !IsSpecified(v) {
		return defaultValue
	}
	return v
}

func IsSpecified[T Float](f T) bool {
	return !math.IsNaN(float64(f))
}

func IsUnspecified[T Float](f T) bool {
	return math.IsNaN(float64(f))
}

func IsInfinite[T Float](f T) bool {
	if !IsSpecified(f) {
		return false
	}
	return math.IsInf(float64(f), 0)
}

// floatEquals compares two float32 values with absolute epsilon tolerance.
func Float32Equals(a, b, epsilon float32) bool {
	return math.Abs(float64(a-b)) <= float64(epsilon)
}

// floatEquals compares two float32 values with absolute epsilon tolerance.
func Float64Equals(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

// 4. Merge - composition merge (package-level function)
// Prefers incoming specified values over current values
func Merge[T Float](a, b T) T {
	if IsSpecified(b) {
		return b
	}
	return a
}

// 5. String - stringification (package-level function)
func String[T Float](f T) string {
	if IsUnspecified(f) {
		return fmt.Sprintf("%T{Unspecified}", f)
	}
	return fmt.Sprintf("%T{%v}", f, f)
}

// 7. Same - identity (package-level function)
func Same[T Float](a, b T) bool {
	// For floats, direct comparison might be tricky with NaN,
	// but IsUnspecified checks should handle the sentinel cases.
	// If both are NaNs, they are technically not Equal in Go, but "Same" in our sentinel context?
	// The docs say: "Identity (2 ns)". For primitive types, it's usually `==`.
	// But `NaN == NaN` is false.
	// However, `SameT` often implies strict equality.
	// Let's follow the simple `==` for now, but if both are Unspecified, they are effectively "Same" sentinel.
	if IsUnspecified(a) && IsUnspecified(b) {
		return true
	}
	return a == b
}

// 8. SemanticEqual - semantic equality (package-level function)
func SemanticEqual[T Float](a, b T) bool {
	if IsUnspecified(a) && IsUnspecified(b) {
		return true
	}
	if IsUnspecified(a) || IsUnspecified(b) {
		return false
	}

	// We need to switch on type to use correct threshold
	// This requires type assertion on the generic type
	vA := any(a)
	vB := any(b)

	switch fA := vA.(type) {
	case float32:
		return Float32Equals(fA, vB.(float32), Float32EqualityThreshold)
	case float64:
		return Float64Equals(fA, vB.(float64), Float64EqualityThreshold)
	default:
		// Should not happen given generic constraint, but fallback to equality
		return a == b
	}
}

// 9. Equal - equality check (package-level function)
// For floats, Equal is usually the same as SemanticEqual (approximate equality)
// or strict equality. The docs patterns usually suggest Equal delegates to SemanticEqual or Same.
func Equal[T Float](a, b T) bool {
	return SemanticEqual(a, b)
}

// 10. Copy - identity for immutable value types (package-level function)
func Copy[T Float](f T) T {
	return f
}
