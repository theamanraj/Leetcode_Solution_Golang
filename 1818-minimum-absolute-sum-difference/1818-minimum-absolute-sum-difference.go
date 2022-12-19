type ParallelIntSlices [][]int

func (p *ParallelIntSlices) Init(slices ...[]int) {
	*p = make([][]int, len(slices))
	for i := range slices {
		(*p)[i] = slices[i]
	}
}
func (p ParallelIntSlices) Len() int {
	return len(p[0])
}
func (p ParallelIntSlices) Less(a, b int) bool {
	return p[0][a] < p[0][b]
}
func (p *ParallelIntSlices) Swap(a, b int) {
	pp := *p
	for i := range pp {
		pp[i][a], pp[i][b] = pp[i][b], pp[i][a]
	}
}

func findClosestRecurse(nums []int, i, j, want int) (index, num, absDiff int) {
	mid := (i + j) / 2
	if nums[mid] == want {
		return mid, want, 0
	}

	if i-j == 0 {
		di := nums[i] - want
		if di < 0 {
			di *= -1
		}
		return i, nums[i], di
	}

	lIdx, lNum, lDiff := findClosestRecurse(nums, i, mid, want)
	rIdx, rNum, rDiff := findClosestRecurse(nums, mid+1, j, want)
	if lDiff < rDiff {
		return lIdx, lNum, lDiff
	}
	return rIdx, rNum, rDiff
}

func findClosest(nums []int, want int) (idx, num, absDiff int) {
	return findClosestRecurse(nums, 0, len(nums)-1, want)
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	nums := ParallelIntSlices{}
	nums.Init(nums1, nums2)
	sort.Sort(&nums)

	var sum uint64

	maxSwapSavings := 0

	for i, num1 := range nums1 {
		num2 := nums2[i]

		diff := num2 - num1
		if diff < 0 {
			diff *= -1
		}

		sum += uint64(diff)
		if diff <= maxSwapSavings {
			continue
		}

		_, _, replaceDiff := findClosest(nums1, num2)
		savings := diff - replaceDiff
		if savings > maxSwapSavings {
			maxSwapSavings = savings
		}
	}

	sum -= uint64(maxSwapSavings)
	modulo := uint64(math.Pow10(9)) + 7
	return int(sum % modulo)
}