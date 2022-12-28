type Solution struct {
    m map[int][]int    
}


func Constructor(nums []int) Solution {
    mMap := make(map[int][]int)
    for i, n := range nums {
        mMap[n] = append(mMap[n], i)
    }
    return Solution{mMap}
}


func (this *Solution) Pick(target int) int {
    idxs := (*this).m[target]
    p := rand.Intn(len(idxs))
    return idxs[p]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */