func circularArrayLoop(nums []int) bool {
    getNext := func (i int) int{
        res:=(i%len(nums) + nums[i]%len(nums))%len(nums)
        if (res>=0) {return res}else {return (res+len(nums))%len(nums)}
    }
    for i:=0; i<len(nums); i++ {
        if nums[i] == 0 {
            continue
        }
        slow:=i
        fast:=i
        for nums[getNext(fast)]*nums[fast]>0 && nums[getNext(getNext(fast))] * nums[getNext(fast)]>0 {
            fast = getNext(getNext(fast))
            slow = getNext(slow)
            if fast == slow {
                if getNext(slow) == slow {
                    break
                }
                return true
            }
        }
        init := nums[slow]
        for init*nums[slow] >0 {
            oldSlow := slow
            slow = getNext(slow)
            nums[oldSlow] = 0
        }
    }
    return false
}