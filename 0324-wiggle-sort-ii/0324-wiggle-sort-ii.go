func wiggleSort(nums []int)  {
    sort.Ints(nums)  
    n := len(nums)
    dummy, mid := make([]int, n), (n-1)/2
	
    for index,i:=0,0;i<=mid;i,index=i+1,index+2{
        dummy[index]=nums[mid-i]    
        if index==n-1 { break }        
        dummy[index+1]=nums[n-1-i]        
    }
   copy(nums, dummy)    
}