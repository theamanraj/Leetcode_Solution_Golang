func minSpeedOnTime(dist []int, hour float64) int {
	if hour <= float64(len(dist)-1) {
		return -1
	}
    // math.MaxInt32 is large enough for right-bound
	l, r := 1, math.MaxInt32
	for l <= r {
		speed := (l+r)>>1
		total := float64(dist[len(dist)-1]) / float64(speed)
		for i := 0; i < len(dist)-1; i++ {
			total += math.Ceil(float64(dist[i])/float64(speed))
		}
        // float64 equal, not 100% right
        // use fraction for last train is better and 100% right
        if math.Abs(total-hour) < 0.00000001 {
            return speed
        }
		if total > hour {
			l = speed+1
		} else {
			r = speed-1
		}
	}
	return r+1
}