package floatutils

// Benchmarks for the float sentinel helpers.
//
// Float32Companion/Float64Companion are value-type scalar companions: the
// operations themselves must not allocate. The interesting axes here are:
//   - NaN sentinel checks (math.IsNaN on the hot path),
//   - the generic helpers vs. the concrete companion (dispatch through func
//     fields vs. direct inlining),
//   - String(), which goes through fmt.Sprintf.
//
// Run: go test ./sentinel/floatutils/ -bench=. -benchmem

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkFloat64Companion(b *testing.B) {
	c := Float64Companion
	spec := 1.5
	unspec := c.Unspecified()

	b.Run("IsSpecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.IsSpecified(spec)
		}
	})
	b.Run("IsSpecified/Unspecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.IsSpecified(unspec)
		}
	})
	b.Run("TakeOrElse/Miss", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.TakeOrElse(unspec, spec)
		}
	})
	b.Run("Merge/BothSpecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Merge(spec, spec)
		}
	})
	b.Run("Merge/UnspecifiedSrc", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Merge(spec, unspec)
		}
	})
	b.Run("Same", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Same(spec, spec)
		}
	})
	b.Run("SemanticEqual/Equal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.SemanticEqual(spec, spec)
		}
	})
	b.Run("SemanticEqual/Unspecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.SemanticEqual(unspec, unspec)
		}
	})
	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Copy(spec)
		}
	})
	b.Run("String/Specified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.String(spec)
		}
	})
	b.Run("String/Unspecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.String(unspec)
		}
	})
}

func BenchmarkFloat32Companion(b *testing.B) {
	c := Float32Companion
	spec := float32(1.5)

	b.Run("IsSpecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.IsSpecified(spec)
		}
	})
	b.Run("SemanticEqual/Equal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.SemanticEqual(spec, spec)
		}
	})
	b.Run("Merge/BothSpecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Merge(spec, spec)
		}
	})
	b.Run("String/Specified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.String(spec)
		}
	})
}

// BenchmarkFloatGenericVsDirect compares the companion dispatch path against the
// generic package-level helpers, quantifying the cost of the func-field
// indirection that the companion pattern introduces.
func BenchmarkFloatGenericVsDirect(b *testing.B) {
	c := Float64Companion
	spec := 1.5

	b.Run("ViaCompanion", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.IsSpecified(spec)
		}
	})
	b.Run("GenericHelper", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = isSpecified(spec)
		}
	})
	b.Run("Inlined", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = !math.IsNaN(spec)
		}
	})
}

func BenchmarkFloatEquals(b *testing.B) {
	b.Run("Float64Equals", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Float64Equals(1.5, 1.5, Float64EqualityThreshold)
		}
	})
	b.Run("Float32Equals", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Float32Equals(1.5, 1.5, Float32EqualityThreshold)
		}
	})
}

func BenchmarkIsInfinite(b *testing.B) {
	b.Run("Infinite", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = IsInfinite(FloatInfinite)
		}
	})
	b.Run("Finite", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = IsInfinite(1.5)
		}
	})
}

func BenchmarkFloatFormatting(b *testing.B) {
	c := Float64Companion
	b.Run("CompanionString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.String(1.5)
		}
	})
	b.Run("SprintfBaseline", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%T{%v}", 1.5, 1.5)
		}
	})
}
