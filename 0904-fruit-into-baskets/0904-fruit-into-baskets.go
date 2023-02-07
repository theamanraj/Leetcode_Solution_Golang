
func totalFruit(fruits []int) int {
    fmap := make(map[int]int)
    l := 0
    r := 0
    max := 0
    for r < len(fruits){
        fmap[fruits[r]]++
        r++
        for len(fmap) > 2{
            fmap[fruits[l]]--
            if fmap[fruits[l]] == 0{
                delete(fmap,fruits[l])
            }
            l++
        }
        max = maxF(max,r-l)
    }
    
    return max
}

func maxF(a,b int)int{
    if a > b{
        return a
    }
    return b
}