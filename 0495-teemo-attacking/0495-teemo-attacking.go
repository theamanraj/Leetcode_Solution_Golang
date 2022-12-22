func findPoisonedDuration(timeSeries []int, duration int) int {
	poisonedSeconds := duration
	for i:= 1; i < len(timeSeries); i++ {
		diff := timeSeries[i] - timeSeries[i-1]
		if  diff < duration {
			poisonedSeconds -= duration - diff
		}
		poisonedSeconds += duration
	}
	return poisonedSeconds
}