func canThreePartsEqualSum(arr []int) bool {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    if sum % 3 != 0 {
        return false
    }
    onePart := sum / 3
    tmp, count := 0, 0
    for _, v := range arr {
        tmp += v
        if tmp == onePart {
            tmp = 0
            count ++
            if count == 3 {
                return true
            }
        }
    }
    return false
}