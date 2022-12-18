func toHex(num int) string {
	// edge case
	if num == 0 {
		return "0"
	}

	var (
		i  int
		sb = make([]byte, 8)
	)

	for i = 7; num != 0; num, i = int(uint32(num)>>4), i-1 {
		c := byte(num & 15) // 4 least significant bits
		if c < 10 {
			sb[i] = byte(c + 48) // 0123456789
		} else {
			sb[i] = byte(c + 87) // abcdef
		}
	}
	return string(sb[i+1:])
}