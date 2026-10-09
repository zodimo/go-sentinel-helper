package floatutils

import (
	"fmt"
	"math"

	"github.com/zodimo/go-sentinel-helper/sentinel/companions"
)

const (
	Float64EqualityThreshold float64 = 1e-9
	Float32EqualityThreshold float32 = 1e-6
)

var Float64Unspecified = math.NaN()
var Float32Unspecified = float32(math.NaN())

var Float32Infinite = float32(math.Inf(1))
var FloatInfinite = math.Inf(1)

// Float is the constraint shared by the float sentinel helpers. It is exposed
// so downstream packages can build their own companions over either width.
type Float interface {
	~float32 | ~float64
}

// Unsafe/private generic helpers. The generic functions stay unexported because
// a single companion cannot be generic over both float32 and float64; instead
// the package exposes two concrete companions (Float32Companion,
// Float64Companion) built from these helpers.

func takeOrElse[T Float](v, defaultValue T) T {
	if !isSpecified(v) {
		return defaultValue
	}
	return v
}

func isSpecified[T Float](f T) bool {
	return !math.IsNaN(float64(f))
}

func isUnspecified[T Float](f T) bool {
	return math.IsNaN(float64(f))
}

// IsInfinite reports whether f is a specified infinity. It is an extra beyond
// the ten canonical companion operations, so it lives alongside Float32Equals/
// Float64Equals as a package-level helper rather than on the companion.
func IsInfinite[T Float](f T) bool {
	if !isSpecified(f) {
		return false
	}
	return math.IsInf(float64(f), 0)
}

// Float32Equals compares two float32 values with absolute epsilon tolerance.
func Float32Equals(a, b, epsilon float32) bool {
	return math.Abs(float64(a-b)) <= float64(epsilon)
}

// Float64Equals compares two float64 values with absolute epsilon tolerance.
func Float64Equals(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func merge[T Float](a, b T) T {
	if isSpecified(b) {
		return b
	}
	return a
}

func stringOf[T Float](f T) string {
	if isUnspecified(f) {
		return fmt.Sprintf("%T{Unspecified}", f)
	}
	return fmt.Sprintf("%T{%v}", f, f)
}

func same[T Float](a, b T) bool {
	if isUnspecified(a) && isUnspecified(b) {
		return true
	}
	return a == b
}

func semanticEqual[T Float](a, b T) bool {
	if isUnspecified(a) && isUnspecified(b) {
		return true
	}
	if isUnspecified(a) || isUnspecified(b) {
		return false
	}

	vA := any(a)
	vB := any(b)
	switch fA := vA.(type) {
	case float32:
		return Float32Equals(fA, vB.(float32), Float32EqualityThreshold)
	case float64:
		return Float64Equals(fA, vB.(float64), Float64EqualityThreshold)
	default:
		return a == b
	}
}

func equal[T Float](a, b T) bool {
	return semanticEqual(a, b)
}

func copyValue[T Float](f T) T {
	return f
}

// Float32Companion is the scalar companion for float32 sentinels.
var Float32Companion = companions.NewScalarCompanion(
	func() float32 { return Float32Unspecified },
	func(f float32) string { return stringOf(f) },
	func(f float32) bool { return isSpecified(f) },
	func(f float32) bool { return isUnspecified(f) },
	func(v, def float32) float32 { return takeOrElse(v, def) },
	func(a, b float32) float32 { return merge(a, b) },
	func(a, b float32) bool { return same(a, b) },
	func(a, b float32) bool { return semanticEqual(a, b) },
	func(a, b float32) bool { return equal(a, b) },
	func(f float32) float32 { return copyValue(f) },
)

// Float64Companion is the scalar companion for float64 sentinels.
var Float64Companion = companions.NewScalarCompanion(
	func() float64 { return Float64Unspecified },
	func(f float64) string { return stringOf(f) },
	func(f float64) bool { return isSpecified(f) },
	func(f float64) bool { return isUnspecified(f) },
	func(v, def float64) float64 { return takeOrElse(v, def) },
	func(a, b float64) float64 { return merge(a, b) },
	func(a, b float64) bool { return same(a, b) },
	func(a, b float64) bool { return semanticEqual(a, b) },
	func(a, b float64) bool { return equal(a, b) },
	func(f float64) float64 { return copyValue(f) },
)
