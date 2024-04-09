package main

import "testing"

func TestMergeAlternately1(t *testing.T) {
	word1 := "abc"
	word2 := "pqr"
	expected := "apbqcr"
	merged := mergeAlternately(word1, word2)

	if merged != expected {
		t.Fatalf(`mergeAlternately(%q, %q) = %q, want %q`, word1, word2, merged, expected)
	}
}

func TestMergeAlternately2(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"
	expected := "apbqrs"
	merged := mergeAlternately(word1, word2)

	if merged != expected {
		t.Fatalf(`mergeAlternately(%q", %q) = %q, want %q`, word1, word2, merged, expected)
	}
}

func TestMergeAlternately3(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"
	expected := "apbqcd"
	merged := mergeAlternately(word1, word2)

	if merged != expected {
		t.Fatalf(`mergeAlternately(%q, %q) = %q, want %q`, word1, word2, merged, expected)
	}
}
