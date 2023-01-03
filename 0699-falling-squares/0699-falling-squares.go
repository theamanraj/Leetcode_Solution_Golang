func fallingSquares(positions [][]int) []int {
	curMaxHeight := make([]int, len(positions))
	curSquares := [][]int{{positions[0][0], positions[0][0]+positions[0][1], positions[0][1]}}
	curMaxHeight[0] = positions[0][1]
	for i := 1; i < len(positions); i++ {
		curHeight := binarySearch(&curSquares, positions[i])
		if curHeight > curMaxHeight[i-1] {
			curMaxHeight[i] = curHeight
		} else {
			curMaxHeight[i] = curMaxHeight[i-1]
		}
	}
	return curMaxHeight
}

func binarySearch(pCurSquares *[][]int, position []int) int {
	l, r := 0, len(*pCurSquares)-1
	// find first X while (*pCurSquares)[X][1] greater than position[0]
	for l <= r {
		m := (l+r)>>1
		if (*pCurSquares)[m][1] <= position[0] {
			l = m+1
		} else {
			r = m-1
		}
	}
	// (*pCurSquares)[r+1][1], first greater than position[0]
	// CASE 1: r+1 is invalid, append to *pCurSquares
	if r+1 >= len(*pCurSquares) {
		if (*pCurSquares)[r][1] == position[0] && (*pCurSquares)[r][2] == position[1] {
			(*pCurSquares)[r][1] = position[0] + position[1]
			return position[1]
		}
		*pCurSquares = append(*pCurSquares, []int{position[0], position[0]+position[1], position[1]})
		return position[1]
	}
	// CASE 2: (*pCurSquares)[r+1][0] ≥ position[0]+position[1], new square fall on ground
	if (*pCurSquares)[r+1][0] >= position[0]+position[1] {
		*pCurSquares = append(*pCurSquares, nil)
		copy((*pCurSquares)[r+2:], (*pCurSquares)[r+1:])
		(*pCurSquares)[r+1] = []int{position[0], position[0]+position[1], position[1]}
		return position[1]
	}
	// CASE 3: square fall on square(s)
	curHeight := position[1]
	for index := r+1; index < len(*pCurSquares); index++ {
		if tempHeight := (*pCurSquares)[index][2] + position[1]; tempHeight > curHeight {
			curHeight = tempHeight
		}
		// check if reached last
		if index == len(*pCurSquares)-1 {
			more := []int{}
			if position[0] + position[1] <= (*pCurSquares)[index][1] {
				more = []int{position[0] + position[1], (*pCurSquares)[index][1], (*pCurSquares)[index][2]}
			}
			if (*pCurSquares)[r+1][0] >= position[0] {
				(*pCurSquares)[r+1][0] = position[0]
				(*pCurSquares)[r+1][1] = position[0] + position[1]
				(*pCurSquares)[r+1][2] = curHeight
				*pCurSquares = (*pCurSquares)[:r+2]
			} else {
				(*pCurSquares)[r+1][1] = position[0]
				*pCurSquares = (*pCurSquares)[:r+2]
				*pCurSquares = append(*pCurSquares, []int{position[0], position[0]+position[1], curHeight})
			}
			if len(more) != 0 {
				*pCurSquares = append(*pCurSquares, more)
			}
			break
		}
		// check if new square will fall next index is possible
		if position[0] + position[1] <= (*pCurSquares)[index+1][0] {
			// impossible fall on later index
			newElements := make([][]int, 0)
			if (*pCurSquares)[r+1][0] < position[0] {
				newElements = append(newElements, []int{(*pCurSquares)[r+1][0], position[0], (*pCurSquares)[r+1][2]})
			}
			newElements = append(newElements, []int{position[0], position[0]+position[1], curHeight})
			if position[0] + position[1] < (*pCurSquares)[index][1] {
				newElements = append(newElements, []int{position[0] + position[1], (*pCurSquares)[index][1], (*pCurSquares)[index][2]})
			}
			// lengthChange may positive, zero and negative
			lengthChange := r - index + len(newElements)
			for i := 0; i < lengthChange; i++ {
				*pCurSquares = append(*pCurSquares, nil)
			}
			copy((*pCurSquares)[index+1+lengthChange:], (*pCurSquares)[index+1:])
			copy((*pCurSquares)[r+1:], newElements)
			if lengthChange < 0 {
				*pCurSquares = (*pCurSquares)[:len(*pCurSquares)+lengthChange]
			}
			break
		}
	}
	return curHeight
}