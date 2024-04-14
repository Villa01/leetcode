package tests

import (
	"leetcode/problems"
	"testing"
)

func TestStrStr(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	expected := 0
	output := problems.StrStr(haystack, needle)

	if output != expected {
		t.Fatalf(`StrStr(%q", %q) = %q, want %q`, haystack, needle, output, expected)
	}
}

func TestStrStr2(t *testing.T) {
	haystack := "leetcode"
	needle := "leeto"
	expected := -1
	output := problems.StrStr(haystack, needle)

	if output != expected {
		t.Fatalf(`StrStr(%q", %q) = %q, want %q`, haystack, needle, output, expected)
	}
}
