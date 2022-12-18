func lemonadeChange(bills []int) bool {

	tmp := make(map[int]int)

	if bills[0] != 5 {
		return false
	}
    for _, v := range bills {
        if v == 5 {
            tmp[v]++
        }
        if v == 10 {
            tmp[5]--
            tmp[v]++
            if tmp[5] < 0 {
                return false
            }
            continue
        }
        if v == 20 {
            if tmp[5] == 0 {
                return false
            }
            if tmp[5] < 3 && tmp[10] < 1 {
                return false
            }
            if tmp[10] < 1 {
                tmp[5] -= 3
                if tmp[5] < 0 {
                    return false
                }
                continue
            }
            tmp[10]--
            tmp[5]--
            if tmp[10] < 0 || tmp[5] < 0 {
                return false
            }
        }
    }
	return true
}