import "sort"
type elems []int
func (el elems) Len() int { return len(el) }
func (el elems) Swap(i, j int) { el[i], el[j] = el[j], el[i] }
func (el elems) Less(i, j int) bool { return el[i] <= el[j] }

func largestDivisibleSubset(nums []int) []int {
    if len(nums) == 0 || len(nums) == 1 { return nums }
    divisors, prev := make([]int, len(nums)), make([]int, len(nums))
    for i:=0; i < len(nums); i++ {
        prev[i] = -1
    }
    sort.Sort(elems(nums))
    lIdx := 0 //longestDivisibleSubset index
    for i:=0;i<len(nums); i++ {
        for j:=0; j < i; j++ {
            if nums[i]%nums[j] == 0 && (divisors[j] + 1 > divisors[i]) {
                divisors[i] = divisors[j] + 1
                prev[i] = j
            }
        }
        if divisors[lIdx] < divisors[i] {
            lIdx = i
        }
    }
    res := []int{}
    i := lIdx
    for i >= 0 {
        res = append(res, nums[i])
        i = prev[i]
    }

    return res
}