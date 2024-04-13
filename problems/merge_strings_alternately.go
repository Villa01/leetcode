package problems

import "strings"

func MergeAlternately(word1 string, word2 string) string {
	merged := strings.Builder{}

	if len(word1) < 1 || len(word1) > 100 {
		panic("word1 should contain between 1 and 100 characters")
	}

	if len(word2) < 1 || len(word2) > 100 {
		panic("word2 should contain between 1 and 100 characters")
	}

	// Lenght of the biggest word
	i := len(word1)
	if len(word2) > len(word1) {
		i = len(word2)
	}

	for j := 0; j < i; j++ {
		if j >= len(word1) {
			merged.WriteString(string(word2[j:]))
			break
		}
		if j >= len(word2) {
			merged.WriteString(string(word1[j:]))
			break
		}
		merged.WriteString(string(word1[j]))
		merged.WriteString(string(word2[j]))
	}

	return merged.String()
}
