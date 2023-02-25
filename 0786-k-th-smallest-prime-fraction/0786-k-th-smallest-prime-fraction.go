type fraction struct {
	numerator   int
	denominator int
	value       float64
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	fractions := make([]fraction, 0, len(arr)*len(arr))
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			fractions = append(fractions, fraction{
				numerator:   arr[i],
				denominator: arr[j],
				value:       float64(arr[i]) / float64(arr[j]),
			})
		}
	}
	sort.Slice(fractions, func(i, j int) bool {
		return fractions[i].value < fractions[j].value
	})
	return []int{fractions[k-1].numerator, fractions[k-1].denominator}
}