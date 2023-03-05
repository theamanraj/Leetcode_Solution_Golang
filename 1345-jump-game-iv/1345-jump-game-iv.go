func minJumps(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	feet := make([]bool, len(arr))
	linked := make(map[int]bool)
	hash := make(map[int][]int)
	for i, v := range arr {
		hash[v] = append(hash[v], i)
	}
	target := len(arr) - 1
	jumpTime := 0
	current := []int{0}
	feet[0] = true
	for {
		next := make([]int, 0, 256)
		jumpTime++
		for _, index := range current {
			// jump forward
			if feet[index+1] == false {
				if index+1 == target {
					return jumpTime
				}
				next = append(next, index+1)
				feet[index+1] = true
			}
			// jump back
			if index-1 >= 0 && feet[index-1] == false {
				next = append(next, index-1)
				feet[index-1] = true
			}
			if linked[arr[index]] == false {
				for _, v := range hash[arr[index]] {
					if feet[v] == false {
						if v == target {
							return jumpTime
						}
						next = append(next, v)
						feet[v] = true
					}
				}
				linked[arr[index]] = true
			}
		}
		current = next
	}
}