import (
	"math"
	"strconv"
	"strings"
)

func Solve(idx, dots int, combination []string, s string, final *[]string) {
	if dots == 4 && idx == len(s) {
		temp := strings.Join(combination, "")
		*final = append(*final, temp[:len(temp)-1])
		return
	}
	if dots > 4 {
		return
	}
	for i := idx; i < int(math.Min(float64(idx+3), float64(len(s)))); i++ {
		temp, _ := strconv.Atoi(s[idx : i+1])
		if temp < 256 && (idx == i || s[idx] != '0') {
			tempStr := s[idx:i+1] + "."
			Solve(i+1, dots+1, append(combination, tempStr), s, final)
		}
	}
}

func restoreIpAddresses(s string) []string {
	final := []string{}
	Solve(0, 0, []string{}, s, &final)
	return final
}