package problems

func FindTheDifference(s string, t string) byte {

	l := len(s)
	if len(t) < l {
		l = len(t)
	}

	for i := 0; i < l; i++ {
		if s[i] != t[i] {
			return t[i]
		}
	}

	return t[l]
}
