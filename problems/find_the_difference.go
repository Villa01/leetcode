package problems

import (
	"fmt"
	"strings"
)

func FindTheDifference(s string, t string) byte {

	for _, char := range s {
		t = strings.Replace(t, string(char), "", 1)
	}

	if len(t) == 0 {
		panic(fmt.Sprintf("there is no difference between %q and %q", s, t))
	}
	return byte(t[0])
}

func FindTheDifferenceASCII(s string, t string) byte {
	res := 0
	for _, char := range t {
		res += int(char)
	}

	for _, char := range s {
		res -= int(char)
	}
	return byte(res)
}
