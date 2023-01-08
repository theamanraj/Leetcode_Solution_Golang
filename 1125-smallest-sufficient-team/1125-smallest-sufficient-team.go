func smallestSufficientTeam(req_skills []string, people [][]string) []int {
    d := make(map[string]int)
    for i, skill := range req_skills {
        d[skill] = i
    }
    
    type state struct {
        pre int
        last int
        count int
    }
    
    n := len(req_skills)
    states := make([]*state, 1 << uint(n))
    states[0] = &state{}
    for ith, skills := range people {
        s := 0
        for _, skill := range skills {
            s |= 1 << uint(d[skill])
        }
        
        for i := (1 << uint(n)) - 1; i >= 0; i-- {
            next := i | s
            if i == next {
                continue
            }

            if states[i] != nil {
                if states[next] == nil || ((*states[next]).count > (*states[i]).count + 1) {
                    if states[next] == nil {
                        states[next] = &state{}
                    }
                    (*states[next]).count = (*states[i]).count + 1
                    (*states[next]).last = ith
                    (*states[next]).pre = i
                }
            }
        }
    }
    
    var ans []int
    s := (1 << uint(n)) - 1
    for s != 0 {
        ans = append(ans, (*states[s]).last)
        s = (*states[s]).pre
    }
    return ans
}