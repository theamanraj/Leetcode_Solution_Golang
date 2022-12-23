func nextGreaterElements(nums []int) []int {
    res := make([]int,len(nums))
    var stack []int
    for i:=2*len(nums)-1;i>=0;i-- {
        for len(stack) != 0 && nums[stack[len(stack)-1]] <= nums[i%len(nums)] {
            stack = stack[:len(stack)-1]
        }
        if len(stack)  == 0 {
            res[i%len(nums)] =  -1
        } else {
            res[i%len(nums)] = nums[stack[len(stack)-1]]
        }
        stack = append(stack,i % len(nums))
    }
    return res
}