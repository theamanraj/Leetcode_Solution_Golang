func trap(heights []int) int {
    
    left,right,water := 0,len(heights)-1,0
    
    maxLeft,maxRight := heights[left],heights[right]
    
    for left < right {
        if heights[left] <= heights[right]{
            left++
            maxLeft = max(maxLeft, heights[left])
            water += (maxLeft-heights[left])
        } else {
            right--;
            maxRight = max(maxRight, heights[right]);
            water += (maxRight - heights[right]);
        }
    }
    return water
}

func max( x int,  y int) int{
    if x>y{
        return x
    }
    return y
}