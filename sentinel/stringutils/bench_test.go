package stringutils

// Benchmarks for the string sentinel helpers.
//
// StringValue is an alias for string, so all operations are value-semantics and
// must be allocation-free. The only allocating operation is String(), which
// calls fmt.Sprintf.
//
// Run: go test ./sentinel/stringutils/ -bench=. -benchmem

import (
	"fmt"
	"testing"
)

func BenchmarkStringValue(b *testing.B) {
	c := StringsCompanion
	spec := StringValue("hello world")
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
	b.Run("Merge/EmptySrc", func(b *testing.B) {
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

func BenchmarkStringValueFormatting(b *testing.B) {
	c := StringsCompanion
	b.Run("CompanionString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.String("hello world")
		}
	})
	b.Run("SprintfBaseline", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("StringValue{%q}", "hello world")
		}
	})
}
