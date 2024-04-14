package problems

import (
	"strings"
)

func StrStr(haystack string, needle string) int {

	return strings.Index(haystack, needle)
}

func StrStrManual(haystack string, needle string) int {

	for i := 0; i <= len(haystack)-len(needle); i++ {

		substring := haystack[i : i+len(needle)]

		if strings.Compare(substring, needle) == 0 {
			return i
		}
	}

	return -1
}
