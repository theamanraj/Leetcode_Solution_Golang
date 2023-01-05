func largestTriangleArea(points [][]int) float64 {
    n := len(points)
    var ans float64
    for i:=0;i<n;i++ {
        for j:=i+1;j<n;j++ {
            for k:=j+1;k<n;k++ {
                ans = max(ans, area(points[i],points[j],points[k]))
        }
    }
       
}
     return ans
}
    
    
    func area(p []int, q []int, r []int) float64 {
        y := p[0] * q[1] + q[0] * r[1] + r[0]* p[1]- (p[1]*q[0] + q[1]*r[0] + r[1]*p[0])
        
        return 0.5 * float64(abs(y))
    }

    func max(a, b float64) float64 {
        if a > b {
            return a
        }
        
        return b
    }
    
    func abs( a int) int {
        if a < 0 {
            return -1 * a
        }
        return a
    }