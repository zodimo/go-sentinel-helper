package stringutils

import (
	"testing"
)

func TestStringValue(t *testing.T) {
	// 1. Sentinel
	if StringsCompanion.IsSpecified(stringValueUnspecified) {
		t.Error("StringValueUnspecified should not be specified")
	}
	if !StringsCompanion.IsUnspecified(stringValueUnspecified) {
		t.Error("StringValueUnspecified should be unspecified")
	}

	val := "hello"
	if !StringsCompanion.IsSpecified(val) {
		t.Error("'hello' should be specified")
	}
	if StringsCompanion.IsUnspecified(val) {
		t.Error("'hello' should not be unspecified")
	}

	// Empty string is valid and specified
	empty := ""
	if !StringsCompanion.IsSpecified(empty) {
		t.Error("Empty string should be specified")
	}
	if StringsCompanion.IsUnspecified(empty) {
		t.Error("Empty string should not be unspecified")
	}

	// 2. TakeOrElse
	if got := StringsCompanion.TakeOrElse(val, "world"); got != val {
		t.Errorf("TakeOrElse('hello', 'world') = %q; want %q", got, val)
	}
	if got := StringsCompanion.TakeOrElse(stringValueUnspecified, "world"); got != "world" {
		t.Errorf("TakeOrElse(Unspecified, 'world') = %q; want 'world'", got)
	}
	// TakeOrElse with empty string
	if got := StringsCompanion.TakeOrElse("", "default"); got != "" {
		t.Errorf("TakeOrElse('', 'default') = %q; want ''", got)
	}

	// 3. Merge
	if got := StringsCompanion.Merge("hello", "world"); got != "world" {
		t.Errorf("Merge('hello', 'world') = %q; want 'world'", got)
	}
	if got := StringsCompanion.Merge("hello", stringValueUnspecified); got != "hello" {
		t.Errorf("Merge('hello', Unspecified) = %q; want 'hello'", got)
	}
	if got := StringsCompanion.Merge(stringValueUnspecified, "world"); got != "world" {
		t.Errorf("Merge(Unspecified, 'world') = %q; want 'world'", got)
	}
	// Merge with empty string
	if got := StringsCompanion.Merge(stringValueUnspecified, ""); got != "" {
		t.Errorf("Merge(Unspecified, '') = %q; want ''", got)
	}
	if got := StringsCompanion.Merge("foo", ""); got != "" {
		t.Errorf("Merge('foo', '') = %q; want ''", got)
	}

	// 4. String
	if got := StringsCompanion.String(stringValueUnspecified); got != "StringValue{Unspecified}" {
		t.Errorf("StringString(Unspecified) = %q; want 'StringValue{Unspecified}'", got)
	}
	if got := StringsCompanion.String("hello"); got != "StringValue{\"hello\"}" {
		t.Errorf("StringString('hello') = %q; want 'StringValue{\"hello\"}'", got)
	}

	// 5. Identity/Equals
	if !StringsCompanion.Same("hello", "hello") {
		t.Error("Same('hello', 'hello') should be true")
	}
	if StringsCompanion.Same("hello", "world") {
		t.Error("Same('hello', 'world') should be false")
	}

	if !StringsCompanion.SemanticEqual("hello", "hello") {
		t.Error("SemanticEqual('hello', 'hello') should be true")
	}
	if !StringsCompanion.Equal("hello", "hello") {
		t.Error("Equal('hello', 'hello') should be true")
	}

	// 6. Copy
	if StringsCompanion.Copy("hello") != "hello" {
		t.Error("Copy('hello') should be 'hello'")
	}
}
