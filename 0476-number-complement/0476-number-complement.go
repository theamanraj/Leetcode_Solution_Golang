func findComplement(num int) int {
	numStrBin := strconv.FormatInt(int64(num), 2)
	p := len(numStrBin)
	x := math.Pow(2, float64(p))

	return num ^ (int(x) - 1)
}