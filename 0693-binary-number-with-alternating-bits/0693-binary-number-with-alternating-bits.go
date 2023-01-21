func hasAlternatingBits(n int) bool {
	num := uint(n)
	num2 := bits.RotateLeft(num,1)
	if num % 2  == 0 {
		num2 = num2 + 1
	}
	log := math.Log2(float64((num ^ num2)+1))
	return log == float64(int(log))
}