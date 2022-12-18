func bestHand(ranks []int, suits []byte) string {
    sm := make(map[byte]struct{})
    for _, suit := range suits {
        sm[suit] = struct{}{}
    }
    if len(sm) == 1 {
        return "Flush"
    }
    rm := make(map[int]int)
    for _, rank := range ranks {
        if v, ok := rm[rank]; ok {
            rm[rank] = v + 1
        } else {
            rm[rank] = 1
        }
    }
    for _, v := range rm {
        if v >= 3 {
            return "Three of a Kind"
        }
    }
    for _, v := range rm {
        if v == 2 {
            return "Pair"
        }
    }
    return "High Card"
}