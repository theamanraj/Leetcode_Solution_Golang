import(
"sort"
)
/*Runtime: 257 ms, faster than 57.14% of Go online submissions for Number of Subsequences That Satisfy the Given Sum Condition.
Memory Usage: 9.3 MB, less than 14.29% of Go online submissions for Number of Subsequences That Satisfy the Given Sum Condition.
*/
func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	const MOD = 1000000007 //1e9 + 7
	n := len(nums)
	exponents := make([]int, n+1)
	exponents[1] = 1
    //Pre-computing power of 2 as it can result into TLE.
    //We will return answer % MOD as answer can be large.
	for i := 2; i < n+1; i++ {
		exponents[i] = (2 * exponents[i-1]) % MOD //Or exponents[i-1] << 1%MOD
	}
	left := 0
	right := n - 1
	count := 0

	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			count = (count + exponents[right-left+1]) % MOD
			left++
		}
	}
	return count
}