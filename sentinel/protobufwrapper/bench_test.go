package protobufwrapper

// Benchmarks for the protobuf wrapper companions.
//
// These types are pointer-based (and BytesValue/stringValue hold heap data), so
// the operations that look "free" at the call site are the ones that allocate.
// Every benchmark reports allocs/op; the ones expected to be allocation-free
// assert it via b.ReportAllocs() + a follow-up b.ReportMetric, so a regression
// shows up as a failing/regressed number rather than a silent cost.
//
// Run: go test ./sentinel/protobufwrapper/ -bench=. -benchmem

import (
	"fmt"
	"testing"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// --- Int32Value -------------------------------------------------------------

func BenchmarkInt32Value(b *testing.B) {
	c := Int32ValueCompanion
	spec := &wrapperspb.Int32Value{Value: 7}
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
	b.Run("Same/Pointer", func(b *testing.B) {
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
	b.Run("String/Specified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.String(spec)
		}
	})
	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Copy(spec)
		}
	})
	b.Run("Coalesce/Nil", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Coalesce(nil, spec)
		}
	})
}

// --- Int64Value -------------------------------------------------------------

func BenchmarkInt64Value(b *testing.B) {
	c := Int64ValueCompanion
	spec := &wrapperspb.Int64Value{Value: 7}
	unspec := c.Unspecified()

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
	b.Run("Equal/SamePointer", func(b *testing.B) {
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
}

// --- UInt32Value ------------------------------------------------------------

func BenchmarkUInt32Value(b *testing.B) {
	c := UInt32ValueCompanion
	spec := &wrapperspb.UInt32Value{Value: 7}

	b.Run("Merge/BothSpecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Merge(spec, spec)
		}
	})
	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Copy(spec)
		}
	})
}

// --- UInt64Value ------------------------------------------------------------

func BenchmarkUInt64Value(b *testing.B) {
	c := UInt64ValueCompanion
	spec := &wrapperspb.UInt64Value{Value: 7}

	b.Run("Merge/BothSpecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Merge(spec, spec)
		}
	})
	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Copy(spec)
		}
	})
}

// --- FloatValue -------------------------------------------------------------

func BenchmarkFloatValue(b *testing.B) {
	c := FloatValueCompanion
	spec := &wrapperspb.FloatValue{Value: 1.5}

	b.Run("Merge/BothSpecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Merge(spec, spec)
		}
	})
	b.Run("SemanticEqual", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.SemanticEqual(spec, spec)
		}
	})
	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Copy(spec)
		}
	})
}

// --- DoubleValue ------------------------------------------------------------

func BenchmarkDoubleValue(b *testing.B) {
	c := DoubleValueCompanion
	spec := &wrapperspb.DoubleValue{Value: 1.5}

	b.Run("Merge/BothSpecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Merge(spec, spec)
		}
	})
	b.Run("SemanticEqual", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.SemanticEqual(spec, spec)
		}
	})
	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Copy(spec)
		}
	})
}

// --- BoolValue --------------------------------------------------------------

func BenchmarkBoolValue(b *testing.B) {
	c := BoolValueCompanion
	spec := &wrapperspb.BoolValue{Value: true}

	b.Run("Merge/BothSpecified", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Merge(spec, spec)
		}
	})
	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Copy(spec)
		}
	})
}

// --- StringValue ------------------------------------------------------------

func BenchmarkStringValue(b *testing.B) {
	c := StringValueCompanion
	spec := &wrapperspb.StringValue{Value: "hello world"}
	unspec := c.Unspecified()

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
	b.Run("SemanticEqual", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.SemanticEqual(spec, spec)
		}
	})
	b.Run("TakeOrElse/Miss", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.TakeOrElse(unspec, spec)
		}
	})
	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.Copy(spec)
		}
	})
}

// --- BytesValue -------------------------------------------------------------
//
// BytesValue is the only wrapper whose Merge/Copy touch the payload slice, so it
// is the one place where an O(n) copy is unavoidable. Benchmarks sweep payload
// sizes to expose the linear term.

func BenchmarkBytesValue(b *testing.B) {
	c := BytesValueCompanion
	small := &wrapperspb.BytesValue{Value: make([]byte, 16)}
	medium := &wrapperspb.BytesValue{Value: make([]byte, 1024)}
	large := &wrapperspb.BytesValue{Value: make([]byte, 1<<20)}

	for _, tc := range []struct {
		name string
		v    *wrapperspb.BytesValue
	}{
		{"16B", small}, {"1KiB", medium}, {"1MiB", large},
	} {
		b.Run("Merge/"+tc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(tc.v.Value)))
			for i := 0; i < b.N; i++ {
				_ = c.Merge(tc.v, tc.v)
			}
		})
		b.Run("Copy/"+tc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(tc.v.Value)))
			for i := 0; i < b.N; i++ {
				_ = c.Copy(tc.v)
			}
		})
		b.Run("SemanticEqual/"+tc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(tc.v.Value)))
			for i := 0; i < b.N; i++ {
				_ = c.SemanticEqual(tc.v, tc.v)
			}
		})
	}
}

// --- Companion indirection --------------------------------------------------
//
// The companion stores each operation in a func field, so every call is an
// indirect call through a struct. This benchmark measures that dispatch cost
// against calling the concrete helper directly.

func BenchmarkCompanionDispatch(b *testing.B) {
	spec := &wrapperspb.Int32Value{Value: 7}
	b.Run("ViaCompanion", func(b *testing.B) {
		c := Int32ValueCompanion
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.IsSpecified(spec)
		}
	})
	b.Run("DirectHelper", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = isSpecifiedInt32Value(spec)
		}
	})
}

// --- String formatting sanity ----------------------------------------------
//
// String() calls fmt.Sprintf, which is the single most expensive operation in
// the package. Kept separate so the cost of formatting is obvious in the table.

func BenchmarkStringFormatting(b *testing.B) {
	b.Run("Int32Value", func(b *testing.B) {
		c := Int32ValueCompanion
		spec := &wrapperspb.Int32Value{Value: 7}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = c.String(spec)
		}
	})
	b.Run("SprintfBaseline", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("Int32Value{%d}", 7)
		}
	})
}
