package stringutils

import (
	"testing"
)

func TestStringValue(t *testing.T) {
	// 1. Sentinel
	if IsSpecifiedString(StringValueUnspecified) {
		t.Error("StringValueUnspecified should not be specified")
	}
	if !IsUnspecifiedString(StringValueUnspecified) {
		t.Error("StringValueUnspecified should be unspecified")
	}

	val := "hello"
	if !IsSpecifiedString(val) {
		t.Error("'hello' should be specified")
	}
	if IsUnspecifiedString(val) {
		t.Error("'hello' should not be unspecified")
	}

	// Empty string is valid and specified
	empty := ""
	if !IsSpecifiedString(empty) {
		t.Error("Empty string should be specified")
	}
	if IsUnspecifiedString(empty) {
		t.Error("Empty string should not be unspecified")
	}

	// 2. TakeOrElse
	if got := TakeOrElseString(val, "world"); got != val {
		t.Errorf("TakeOrElse('hello', 'world') = %q; want %q", got, val)
	}
	if got := TakeOrElseString(StringValueUnspecified, "world"); got != "world" {
		t.Errorf("TakeOrElse(Unspecified, 'world') = %q; want 'world'", got)
	}
	// TakeOrElse with empty string
	if got := TakeOrElseString("", "default"); got != "" {
		t.Errorf("TakeOrElse('', 'default') = %q; want ''", got)
	}

	// 3. Merge
	if got := MergeString("hello", "world"); got != "world" {
		t.Errorf("Merge('hello', 'world') = %q; want 'world'", got)
	}
	if got := MergeString("hello", StringValueUnspecified); got != "hello" {
		t.Errorf("Merge('hello', Unspecified) = %q; want 'hello'", got)
	}
	if got := MergeString(StringValueUnspecified, "world"); got != "world" {
		t.Errorf("Merge(Unspecified, 'world') = %q; want 'world'", got)
	}
	// Merge with empty string
	if got := MergeString(StringValueUnspecified, ""); got != "" {
		t.Errorf("Merge(Unspecified, '') = %q; want ''", got)
	}
	if got := MergeString("foo", ""); got != "" {
		t.Errorf("Merge('foo', '') = %q; want ''", got)
	}

	// 4. String
	if got := StringString(StringValueUnspecified); got != "StringValue{Unspecified}" {
		t.Errorf("StringString(Unspecified) = %q; want 'StringValue{Unspecified}'", got)
	}
	if got := StringString("hello"); got != "StringValue{\"hello\"}" {
		t.Errorf("StringString('hello') = %q; want 'StringValue{\"hello\"}'", got)
	}

	// 5. Identity/Equals
	if !SameString("hello", "hello") {
		t.Error("Same('hello', 'hello') should be true")
	}
	if SameString("hello", "world") {
		t.Error("Same('hello', 'world') should be false")
	}

	if !SemanticEqualString("hello", "hello") {
		t.Error("SemanticEqual('hello', 'hello') should be true")
	}
	if !EqualString("hello", "hello") {
		t.Error("Equal('hello', 'hello') should be true")
	}

	// 6. Copy
	if CopyString("hello") != "hello" {
		t.Error("Copy('hello') should be 'hello'")
	}
}
