func maxEnvelopes(envelopes [][]int) int {
    sort.Slice(envelopes, func(a, b int)bool{
        if envelopes[a][0] == envelopes[b][0] {
            //if width of envelope are the same, sort by length in descedending order
            //to avoid duplicates in result, for example if we have below 3 envelopes
            //[1,2], [1,3], [1,4]
            //if sort length in ascending order, then the result is 3 which is incorrect 
            return envelopes[a][1] > envelopes[b][1]
        }
        return envelopes[a][0] < envelopes[b][0]
    })
    list := make([]int, 0)
    
    for _, val := range envelopes {
        if len(list) == 0 || list[len(list) - 1] < val[1] {
            //if the next value is greater than the biggest one, just append it in the end
            list = append(list, val[1])
        } else {
            //find the right place in the list
            helper(list, val[1])
        }
    }
    
    return len(list)
}

//binary search, find the right place to insert the x into list
func helper(list []int, x int) {
    l := 0
    r := len(list) - 1
    for l + 1 < r {
        mid := (l + r) / 2
        //already existing, return
        if list[mid] == x {
            return
        }
        if list[mid] > x {
            r = mid
        } else {
            l = mid
        }
    }
    //already existing, return
    if list[l] == x || list[r] == x {
        return
    }
    
    //if list[l] > x, update it with value x
    if list[l] > x {
        list[l] = x
        return
    }
    
    //if list[r] > x (and x > list[r] for sure), update it with value x
    if list[r] > x {
        list[r] = x
    }
    return
}