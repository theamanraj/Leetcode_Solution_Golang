func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    var cntBoats int
    lPtr := 0
    rPtr := len(people)-1
    LOOP:
    for {
        l := people[lPtr]
        r := people[rPtr]
        switch {
            case lPtr > rPtr:
                break LOOP
            case lPtr == rPtr: 
                cntBoats++
                break LOOP
            case l + r > limit:
                cntBoats++
                rPtr--
            case l+r <= limit: 
                cntBoats++
                lPtr++
                rPtr--
        }
    }
    return cntBoats
}