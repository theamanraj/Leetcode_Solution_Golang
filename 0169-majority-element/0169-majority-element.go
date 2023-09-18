func majorityElement(nums []int) int {
    count,res := 0,0
    for i := 0; i< len(nums); i++ {
        //base case to handle 1st element
        if i == 0 || count == 0 {
            res = nums[i]
        }
        // checking if num == res or not, if it is then incrementing count as 1
        // else will decrease
        //[2,2,1,1,1,2,2] -> ans is 2
        if nums[i] == res {
            count++
        } else {
        count--
        }
    }
    return res
}