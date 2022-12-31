func rangeSum(nums []int, n int, left int, right int) int {
    prefix := make([]int,len(nums))
    sums := make([]int, (n*(n+1)/2))
    prefix[0]=nums[0]
    sums[0]=prefix[0]
    for i, _ := range nums {
        if i==0 {
            continue
        }
        prefix[i] = prefix[i-1]+nums[i]
        sums[i]=prefix[i]
    }
    ptr:=len(nums)
    for i:=0;i<len(nums)-1;i++ {
        for j:=i+1;j<len(nums);j++ {
            sums[ptr]=(prefix[j]-prefix[i])
            ptr++
        }
    }
    sort.Ints(sums)
    final:=0
    for i:=left;i<=right;i++ {
        final+=sums[i-1]
    }
    return final%1000000007
}