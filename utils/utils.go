package utils

func SlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
