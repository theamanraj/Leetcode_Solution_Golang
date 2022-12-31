type Solution struct {
    sum      int    // sumOfWeight
    weight   []int  //  sorted Slice [ preSum(0), preSum(1), preSum(2)]
}
// to build below data structure

func Constructor(w []int) Solution {       
    tmp := make([]int, len(w)+1) // need +1: need set the tmp[0] always as 0, which is start point    
    total := 0
    for i , v := range w {
       total+=v        
       tmp[i+1]=total
    }    
    return Solution{
        sum:   total,
        weight: tmp,
    }
}

func (this *Solution) PickIndex() int { // O log(n)
    // pick random [0, sum], then binary search  
    r := rand.Intn(this.sum)
    for left, right := 0, len(this.weight)-1;left<= right; {
        mid := left+(right-left)/2
        // search from [0, preSum(1)),[ prSum(1),  prSum(2)), [preSum(2), prSum(3)) . 
        // as the requirement is [0, weight-1] inclusive---> [0,weight) half open 
        if r>=this.weight[mid] && r<this.weight[mid+1]{
            return mid
        }else if this.weight[mid]>r{
            right=mid-1
        }else{
            left=mid+1
        }        
    }
    return 0
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */