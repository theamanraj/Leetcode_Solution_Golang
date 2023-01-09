type FindSumPairs struct {
	num1 []int
	num2 []int
	hash map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	hash := make(map[int]int)
	for _, v := range nums2 {
		hash[v]++
	}
    fsp := FindSumPairs{
		num1: nums1,
		num2: nums2,
		hash: hash,
	}
    sort.Ints(fsp.num1)
    return fsp
}

func (fsp *FindSumPairs) Add(index int, val int)  {
	fsp.hash[fsp.num2[index]]--
	fsp.num2[index] += val
	fsp.hash[fsp.num2[index]]++
}

func (fsp *FindSumPairs) Count(tot int) int {
	count := 0
	for _, v := range fsp.num1 {
        if tot-v <= 0 {
            break
        }
		count += fsp.hash[tot-v]
	}
	return count
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */