func minimumTime(time []int, totalTrips int) int64 {
    var maxTime int64
    timeInt64 := make([]int64, len(time))
    for i := range time {
        timeInt64[i] = int64(time[i])
        maxTime = min(maxTime, timeInt64[i]) 
    }

    var lo, hi int64 = 0, maxTime * int64(totalTrips)
    target := int64(totalTrips)
    for lo < hi {
        mid := (hi-lo)/2+lo
        numTrips := tripsMade(mid, timeInt64)
        if numTrips < target {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    return lo
}

func tripsMade(mid int64, timeInt64 []int64) int64 {
    var numTrips int64
    for i := range timeInt64 {
        numTrips += mid / timeInt64[i]
    }
    return numTrips
}

func min(x, y int64) int64 {
    if x < y {
        return y
    }
    return x
}