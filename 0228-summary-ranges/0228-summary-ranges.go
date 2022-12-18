import "fmt"

func summaryRanges(nums []int) []string {
    var ans []string
    for i, j := 0, 0; j < len(nums); j++ {
        i = j
        for j + 1 < len(nums) && nums[j] + 1 == nums[j + 1] {
            j++
        }
        if i == j {
            ans = append(ans, fmt.Sprint(nums[i]))
        } else {
            ans = append(ans, fmt.Sprintf("%v->%v", nums[i], nums[j]))
        }
    }   
    return ans
}