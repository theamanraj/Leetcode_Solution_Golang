func checkSame(day1, day2 []int) bool {
	for index, cell := range day1 {
		if cell != day2[index] {
			return false
		}
	}
	return true
}

func prisonAfterNDays(cells []int, N int) []int {
	repeatCycle := 0
	pastCellsArray := make([][]int, 0)
	for i := 0; i < N; i++ {
		cells = changeState(cells)
		repeatCycle = checkRepeatCycle(cells, pastCellsArray)
		if repeatCycle > 0 {
			
			break
		}
		pastCellsArray = append(pastCellsArray, cells)
	}
	if repeatCycle > 0 {
		
		mod := N % repeatCycle
		if mod != 0 {
			cells = pastCellsArray[mod-1]
		} else {
			cells = pastCellsArray[repeatCycle-1]
		}

	}
	return cells
}

func checkRepeatCycle(todayCell []int, stateArray [][]int) int {
	cycle := 0
	for index, oldCells := range stateArray {
		if checkSame(todayCell, oldCells) {
			cycle = len(stateArray) - index
			break
		}
	}
	return cycle
}

func changeState(cells []int) []int {
	newCells := make([]int, len(cells))
	for index := range cells {
		if index == 0 || index == len(cells)-1 {
			newCells[index] = 0
		} else {
			if cells[index-1] == cells[index+1] {
				newCells[index] = 1
			} else if cells[index-1] != cells[index+1] {
				newCells[index] = 0
			}
		}
	}
	return newCells
}