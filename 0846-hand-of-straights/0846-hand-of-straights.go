func isNStraightHand(hand []int, W int) bool {
    if len(hand)%W != 0 {
        return false
    }
    sort.Ints(hand)
    hash := make(map[int]int)
    for _, num := range hand {
        hash[num]+= 1
    }
    
    for _, num := range hand {
        if hash[num] == 0 {
            continue
        }
        hash[num]-=1
        if !dfs(hash, W-1, num+1)  {
            //fmt.Println(hash, num)
            return false
        }
    }
    return true
}

func dfs(hash map[int]int, window, num int) bool {
    if window == 0 {
        return true
    }
    if hash[num] == 0 {
        return false
    }
    hash[num]-=1
    return dfs(hash, window-1, num+1)
}