func nearestValidPoint(x int, y int, points [][]int) int {
    result := -1
    dist := math.MaxInt32
    
    for i, point := range points {
        if point[0] == x || point[1] == y {
            temp := abs(point[0]-x) + abs(point[1]-y)
            if temp < dist {
                result = i
                dist = temp
            }
        }
    }
    
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}