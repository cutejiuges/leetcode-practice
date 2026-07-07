package LC3754

func sumAndMultiply(n int) int64 {
	base, sum, mult := 1, 0, 0
	for n != 0 {
		remain := n % 10
		if remain != 0 {
			sum += remain
			mult += remain * base
			base *= 10
		}
		n /= 10
	}
	return int64(sum * mult)
}
