func longestConsecutive(nums []int) int {
    mem := make(map[int]int, len(nums))
    result := 0
    
    for _, v := range nums {
        mem[v] = 1
    }
   
    for k,v := range mem {
        for _, ok := mem[k-v]; ok == true ;  {
            
            mem[k] += mem[k-v]
            
            delete(mem, k-v)
            v = mem[k]
            
            _, ok = mem[k-v]
        }
        result = max(result, mem[k])
    }
    
    return result
}

func max(a,b int) int {
    if a > b {
        return a
    }
    
    return b
}