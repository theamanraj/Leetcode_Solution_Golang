func minTimeToVisitAllPoints(points [][]int) int {
    seconds := 0
    for i, _ := range points {
        if i == len(points) - 1 {
            return seconds
        }
        dx := points[i+1][0] - points[i][0]
        dy := points[i+1][1] - points[i][1]

	    dx = int(math.Abs(float64(dx)))
    	dy = int(math.Abs(float64(dy)))

        if dx == dy {
            seconds += int(math.Abs(float64(dx)))
        } else {
            if dx < dy {
                seconds += int(math.Abs(float64(dy)))
            } else {
                seconds += int(math.Abs(float64(dx)))
            }
        }
    }
    return seconds
}