package intutils

// Benchmarks for the int sentinel helpers.
//
// IntValue is an alias for int, so every operation is value-semantics and must
// be allocation-free. String() is the only allocating path (fmt.Sprintf).
//
// Run: go test ./sentinel/intutils/ -bench=. -benchmem

import (
	"fmt"
	"testing"
)

func BenchmarkIntValue(b *testing.B) {
	c := IntValueCompanion
	spec := IntValue(42)
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

func BenchmarkIntValueFormatting(b *testing.B) {
	c := IntValueCompanion
	b.Run("CompanionString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.String(42)
		}
	})
	b.Run("SprintfBaseline", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("IntValue{%d}", 42)
		}
	})
}
