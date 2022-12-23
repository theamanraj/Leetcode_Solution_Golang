func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	sort.Ints(worker)
	type job struct {
		difficulty int
		profit     int
	}
	buildJobs := func(difficulty []int, profit []int) []job {
		j := []job{}
		for i := 0; i < len(difficulty); i++ {
			j = append(j, job{difficulty: difficulty[i], profit: profit[i]})
		}
		return j
	}
	jobs := buildJobs(difficulty, profit)
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].difficulty < jobs[j].difficulty
	})
	sum,j,bestProfit := 0,0,0
	for w := 0; w < len(worker); w++ {
		for j < len(jobs) && jobs[j].difficulty <= worker[w] {
			bestProfit = int(math.Max(float64(bestProfit), float64(jobs[j].profit)))
			j++
		}
		sum += bestProfit
	}
	return sum
}