type Solution struct {
    radius float64
    xMin float64
    xMax float64
    yMin float64
    yMax float64
    
}


func Constructor(radius float64, x_center float64, y_center float64) Solution {
    rand.Seed(time.Now().UnixNano())
    
    return Solution {
        radius : radius,
        xMin : x_center - radius,
        xMax : x_center + radius,
        yMin : y_center - radius,
        yMax : y_center + radius,
    }
}

func getRand(min float64, max float64) float64 {
    
    randomizedRange := min + ((max - min) * rand.Float64())
    return randomizedRange
}

func (this *Solution) RandPoint() []float64 {
        
    x1 := this.xMax - this.radius
    y1 := this.yMax - this.radius
    
    result := []float64{}

    for {
        x2 := getRand(this.xMin, this.xMax)
        y2 := getRand(this.yMin, this.yMax)
        
        if (x2-x1) * (x2-x1) + (y2-y1) * (y2-y1) <= this.radius * this.radius {
            result = append(result, x2, y2)
            break
        }        
    }
    
    return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */