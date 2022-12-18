func hasGroupsSizeX(deck []int) bool {
	counter := make(map[int]int)
	for _, card := range deck {
		counter[card]++
	}
	gcd := 0
	for _, count := range counter {
		gcd = GCD(gcd, count)
	}
	return gcd > 1
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}