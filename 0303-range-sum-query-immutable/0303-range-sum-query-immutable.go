type NumArray struct {
   arr []int
   total int
   before []int
   after []int
}


func Constructor(nums []int) NumArray {
    numarr := NumArray{arr: nums}
    
    total := 0
    for _, x := range nums {
        total += x
    }
    numarr.total = total
    
    n := len(nums)
    b := make([]int, n)
    b[0] = 0
    for i:=1; i < n; i++ {
        b[i] = b[i-1] + nums[i-1] 
    }
    numarr.before = b
    
    a := make([]int, n)
    a[n-1] = 0
    for j := n-2; j >=0; j-- {
        a[j] = a[j+1] + nums[j+1]
    }
    numarr.after = a
    
    return numarr
}


func (this *NumArray) SumRange(left int, right int) int {
    return ((this.total-this.before[left]-this.after[right]))
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */