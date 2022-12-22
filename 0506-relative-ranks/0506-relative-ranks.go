import "strconv"
type Pair struct {
    Idx int
    Score int
}
type ScorePairs []Pair

func(s ScorePairs) Len() int {
    return len(s)
}
func(s ScorePairs) Less(i, j int) bool {
    return s[i].Score > s[j].Score
}
func(s ScorePairs) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
                          
func findRelativeRanks(score []int) []string {
    res := make([]string, len(score))
    pairs := ScorePairs{}
    for idx, s := range score {
        pairs = append(pairs, Pair{idx, s})
    }
    sort.Sort(pairs)
    for idx, p := range pairs {
        res[p.Idx] = getRank(idx+1)
    }
    return res
}

func getRank(rank int) string {
    switch rank {
        case 1: return "Gold Medal"
        case 2: return "Silver Medal"
        case 3: return "Bronze Medal"
    }
    return strconv.Itoa(rank)
}