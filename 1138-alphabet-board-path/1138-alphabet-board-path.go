
var board = []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"}

func alphabetBoardPath(target string) string {
	sequence := ""
	alphabet := strings.Join(board, "")
	coord := [2]int{0, 0}
	for i := range target {
		letter := string(target[i])
		idxInAlphabet := strings.Index(alphabet, letter)
		coord = calculateMovement(coord, idxInAlphabet, letter, &sequence)
	}
	return sequence
}

func calculateMovement(coord [2]int, index int, letter string, sequence *string) [2]int {
	if coord == [2]int{5, 0} && letter == "z" {
		*sequence += "!"
		return coord
	}
	xCoord := int(math.Floor(float64(index) / 5))
	yCoord := strings.Index(board[int(xCoord)], letter)
	xTravel := xCoord - coord[0]
	if letter == "z" {
		xTravel--
	}
	if xTravel < 0 {
		*sequence += strings.Repeat("U", -xTravel)
	} else if xTravel > 0 {
		*sequence += strings.Repeat("D", xTravel)
	}
	yTravel := yCoord - coord[1]
	if yTravel < 0 {
		*sequence += strings.Repeat("L", -yTravel)
	} else if yTravel > 0 {
		*sequence += strings.Repeat("R", yTravel)
	}
	if letter == "z" {
		*sequence += "D"
	}
	*sequence += "!"
	return [2]int{xCoord, yCoord}
}