func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
    ans := make([]bool, 0)
	
    for i := range l{
        if r[i] - l[i]+1 <= 2{
            ans = append(ans, true)
            continue
        }
        
        x := make([]int, r[i] - l[i]+1)
		copy(x, nums[l[i]:r[i]+1])
		sort.Ints(x)
        
        commonDiff := x[1] - x[0]
		tmp := true
		for j := 1; j < len(x) - 1;j++{
			if x[j+1] - x[j] != commonDiff{
				tmp = false
				break
			}
		}

		ans = append(ans, tmp)
	}

	return ans
}