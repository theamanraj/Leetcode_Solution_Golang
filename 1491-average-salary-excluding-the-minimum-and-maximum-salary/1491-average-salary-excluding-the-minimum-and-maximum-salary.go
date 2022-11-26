func average(salary []int) float64 {

	sort.Slice(salary, func(i, j int) bool {
		return salary[i] < salary[j]
	})

	salary = salary[1:]
	salary = salary[:len(salary)-1]

	var sum int
	for i := 0; i < len(salary); i++ {
		sum += salary[i]
	}
	s64 := float64(sum)
	l64 := float64(len(salary))

	return s64 / l64
}