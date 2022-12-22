func distributeCandies(candies []int) int {
    k:= make(map[int]int)
    for _, kvs := range candies {
        if _,ok:=k[kvs];ok {
            k[kvs]+=1
        } else{
            k[kvs]=1
        }
    }
    
    kinds:= len(k)
    perNums:= len(candies)/2
    if perNums < kinds{
        return perNums
    }
    return kinds
}