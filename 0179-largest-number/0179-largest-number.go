import "sort"
import "strings"
import "strconv"

func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        return concatInts(nums[i], nums[j]) >
            concatInts(nums[j], nums[i])
    })

    if nums[0] == 0 {
        return "0"
    }

    var sb strings.Builder
    for _, e := range nums {
        sb.WriteString(strconv.Itoa(e))
    }
    return sb.String()
}

func concatInts(x, y int) int {
    r, t := x*10, y
    for t>9 {
        t /= 10
        r *= 10
    }
    return r+y
}