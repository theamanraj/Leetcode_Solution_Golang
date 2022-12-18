func largeGroupPositions(S string) [][]int {
    res := make([][]int, 0)
    start := 0
    length := 1
    data := []byte(S)
    curr := data[0]
    if len(data) < 3 {
        return res
    }
    for i := 1; i < len(data); i++ {
        if data[i] == curr {
            length++
        } else {
            if length >= 3 {
                tmp := make([]int, 2)
                tmp[0] = start
                tmp[1] = i - 1
                res = append(res, tmp)
            }
            curr = data[i]
            length = 1
            start = i
        }
    }
    if length >= 3 {
        tmp := make([]int, 2)
        tmp[0] = start
        tmp[1] = len(data) - 1
        res = append(res, tmp)
    }
    return res 
}