package companions

import "testing"

// scalarValue models a simple string-backed sentinel type.
type scalarValue string

const scalarUnspecified scalarValue = "\x00unspecified"

func newScalarCompanion() ScalarCompanion[scalarValue] {
	return NewScalarCompanion(
		func() scalarValue { return scalarUnspecified },
		func(s scalarValue) string { return string(s) },
		func(s scalarValue) bool { return s != scalarUnspecified },
		func(s scalarValue) bool { return s == scalarUnspecified },
		func(src, def scalarValue) scalarValue {
			if src != scalarUnspecified {
				return src
			}
			return def
		},
		func(a, b scalarValue) scalarValue {
			if b != scalarUnspecified {
				return b
			}
			return a
		},
		func(a, b scalarValue) bool { return a == b },
		func(a, b scalarValue) bool { return a == b },
		func(a, b scalarValue) bool { return a == b },
		func(s scalarValue) scalarValue { return s },
	)
}

func TestScalarCompanion(t *testing.T) {
	c := newScalarCompanion()

	if c.Unspecified() != scalarUnspecified {
		t.Error("Unspecified() should return the sentinel")
	}
	if !c.IsUnspecified(c.Unspecified()) {
		t.Error("sentinel should be unspecified")
	}
	if c.IsSpecified(c.Unspecified()) {
		t.Error("sentinel should not be specified")
	}
	if !c.IsSpecified("hello") {
		t.Error("hello should be specified")
	}
	if got := c.TakeOrElse(c.Unspecified(), "def"); got != "def" {
		t.Errorf("TakeOrElse(Unspecified, def) = %q; want def", got)
	}
	if got := c.Merge(scalarUnspecified, "b"); got != "b" {
		t.Errorf("Merge(Unspecified, b) = %q; want b", got)
	}
	if !c.Same("a", "a") || !c.SemanticEqual("a", "a") || !c.Equal("a", "a") {
		t.Error("equality helpers should agree on identical values")
	}
	if got := c.Copy("a"); got != "a" {
		t.Errorf("Copy(a) = %q; want a", got)
	}
}

// structValue models a pointer-companion type.
type structValue struct {
	v int
}

func newStructCompanion() StructCompanion[*structValue] {
	unspecified := &structValue{v: -1}
	return NewStructCompanion(
		func() *structValue { return unspecified },
		func(s *structValue) string { return "structValue" },
		func(s *structValue) bool { return s != nil && s != unspecified },
		func(s *structValue) bool { return s == nil || s == unspecified },
		func(src, def *structValue) *structValue {
			if src == nil || src == unspecified {
				return def
			}
			return src
		},
		func(a, b *structValue) *structValue {
			if b == nil || b == unspecified {
				return a
			}
			return b
		},
		func(a, b *structValue) bool { return a == b },
		func(a, b *structValue) bool {
			if a == nil || a == unspecified || b == nil || b == unspecified {
				return false
			}
			return a.v == b.v
		},
		func(a, b *structValue) bool { return a == b },
		func(ptr, def *structValue) *structValue {
			if ptr == nil {
				return def
			}
			return ptr
		},
		func(s *structValue, options ...func(*structValue)) *structValue {
			if s == nil || s == unspecified {
				return unspecified
			}
			out := &structValue{v: s.v}
			for _, opt := range options {
				opt(out)
			}
			return out
		},
	)
}

func TestStructCompanion(t *testing.T) {
	c := newStructCompanion()

	if c.Unspecified() == nil {
		t.Fatal("Unspecified() should return the singleton")
	}
	if !c.IsUnspecified(c.Unspecified()) {
		t.Error("singleton should be unspecified")
	}
	if c.IsSpecified(nil) {
		t.Error("nil should not be specified")
	}
	if !c.IsSpecified(&structValue{v: 1}) {
		t.Error("custom value should be specified")
	}
	if got := c.Coalesce(nil, c.Unspecified()); got == nil {
		t.Error("Coalesce(nil, def) should return def")
	}

	v := &structValue{v: 7}
	// Regression: Copy must not recurse infinitely.
	copied := c.Copy(v, func(sv *structValue) { sv.v = 9 })
	if copied.v != 9 {
		t.Errorf("Copy with option = %d; want 9", copied.v)
	}
	if v.v != 7 {
		t.Errorf("Copy must not mutate source; source = %d; want 7", v.v)
	}
}
