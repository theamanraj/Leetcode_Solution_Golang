func sumOfUnique(nums []int) int {
    
    sort.Ints(nums)
    repeat:=0
    sum:=0
    i:=0
    if len(nums)==1{ return nums[0] }
    
    for i=0;i<len(nums)-1;i++{
         
        if nums[i]==nums[i+1]||nums[i]==repeat{
            repeat=nums[i]   
        }else{
          sum+=nums[i]  
        }
    }
    if nums[i]!=nums[i-1]{
        sum+=nums[i]
    } 
    return sum
}