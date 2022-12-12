func numberOfArithmeticSlices(nums []int) int {
    // If length is less than 3, no AP.
    if len(nums) <3 {
        return 0
    }

    // Initiate diff.
    diff := nums[1]-nums[0]
    // Maximum continous elements with same difference.
    maxApLen := 0
    subSeqCount := 0
    for i:=1;i<len(nums);i++ {
        if nums[i]-nums[i-1] == diff {
            maxApLen++
        } else {
            diff = nums[i]-nums[i-1]
            // Backtrack to the last element to count this in next AP series.
            i--
            // Add the subsequences that can be formed by maximum AP length.
            // If the diff is equal n times, there will be n+1 elements in the AP.
            subSeqCount += subSequences(maxApLen+1)
            // Reset the AP length counter to 0.
            maxApLen = 0
        }
    }
    // Add the last AP series.
    subSeqCount += subSequences(maxApLen+1)
    return subSeqCount
}

// subSequences will output number of sub-sequences that can be formed from AP
// of given length. For example, an AP of length 4 will give us 3 sub-sequences.
func subSequences(maxApLen int) int {
    if maxApLen < 3 {
        return 0
    }
    n := (maxApLen-3+1)
    return (n*(n+1))/2 // Sum of AP with diff 1.
}