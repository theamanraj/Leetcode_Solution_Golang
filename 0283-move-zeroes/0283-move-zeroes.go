func moveZeroes(nums []int)  {
    zcounter := 0
    ncounter := 0
    for _,val := range nums{
        if val == 0{
            zcounter++
        }
        if val != 0{
            nums[ncounter] = val
            ncounter++
        }
    }
    for i:=0;i<zcounter;i++{
        nums[ncounter+i] = 0
    }
}