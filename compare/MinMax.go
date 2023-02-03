package compare

func MinMax(n, min, max int) int {
	if min <= n && n <= max {
		return n
	} else if n > max {
		return max
	} else {
		return min
	}
}
