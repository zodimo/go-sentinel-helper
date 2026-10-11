package boolutils

// Benchmarks for the tri-state BooleanValue.
//
// BooleanValue is a struct wrapping an int enum (Pattern 2-D), chosen precisely
// for zero-allocation value semantics. These benchmarks are the evidence that
// the pattern holds: every companion operation and every convenience method
// should report 0 allocs/op. String()/Format() go through fmt and may allocate.
//
// Run: go test ./sentinel/boolutils/ -bench=. -benchmem

import (
	"fmt"
	"testing"
)

func BenchmarkBooleanValue(b *testing.B) {
	c := BooleanValueCompanion
	spec := BooleanValueTrue()
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
	b.Run("IsUnspecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.IsUnspecified(unspec)
		}
	})
	b.Run("TakeOrElse/Hit", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.TakeOrElse(spec, spec)
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
	b.Run("Same", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Same(spec, spec)
		}
	})
	b.Run("SemanticEqual", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.SemanticEqual(spec, spec)
		}
	})
	b.Run("Equal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Equal(spec, spec)
		}
	})
	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Copy(spec)
		}
	})
	b.Run("String/True", func(b *testing.B) {
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

// BenchmarkBooleanValueConstructors covers the public construction paths.
func BenchmarkBooleanValueConstructors(b *testing.B) {
	b.Run("BooleanValueTrue", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = BooleanValueTrue()
		}
	})
	b.Run("BooleanValueFrom/True", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = BooleanValueFrom(true)
		}
	})
	b.Run("BooleanValueFrom/False", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = BooleanValueFrom(false)
		}
	})
}

// BenchmarkBooleanValueAccessors covers the domain-specific convenience methods
// that live beyond the companion (IsTrue/IsFalse/Bool/BoolOrElse).
func BenchmarkBooleanValueAccessors(b *testing.B) {
	spec := BooleanValueTrue()
	unspec := BooleanValueCompanion.Unspecified()

	b.Run("IsTrue", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = spec.IsTrue()
		}
	})
	b.Run("IsFalse", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = spec.IsFalse()
		}
	})
	b.Run("Bool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = spec.Bool()
		}
	})
	b.Run("BoolOrElse/Specified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = spec.BoolOrElse(false)
		}
	})
	b.Run("BoolOrElse/Unspecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = unspec.BoolOrElse(true)
		}
	})
}

// BenchmarkBooleanValueFormatting covers String()/Format() via fmt.
func BenchmarkBooleanValueFormatting(b *testing.B) {
	spec := BooleanValueTrue()
	b.Run("String", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = spec.String()
		}
	})
	b.Run("SprintfV", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%v", spec)
		}
	})
}
