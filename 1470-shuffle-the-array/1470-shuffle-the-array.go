func shuffle(nums []int, n int) []int {
    if len(nums) == 1 {
        return nums
    }
    
    var s1,s2,result []int
    
    for i:=0;i<n;i++{
        s1 = append(s1,nums[i])
    }
    for j:=n;j<n*2;j++{
        s2 = append(s2,nums[j])
    }
    
    var count1,count2,num int
    
    for {
        if count1 == n && count2 == n {
            return result
        }
        
        if num % 2 == 0 {
            result = append(result, s1[count1])
            count1++
            num++
            
            continue
        }
        result = append(result, s2[count2])
        count2++
        num++
    }
}