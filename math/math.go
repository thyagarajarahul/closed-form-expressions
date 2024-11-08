package math

func Square(n int64) int64 {
	if n < 2 {
		return n
	}

	var numBy2 = n >> 1

	if n&1 != 0 {
		return ((numBy2 << 2) + 1) + (Square(numBy2) << 2)
	} else {
		return (Square(numBy2) << 2)
	}
}
