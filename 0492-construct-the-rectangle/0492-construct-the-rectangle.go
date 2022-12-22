func constructRectangle(area int) []int {
    d := int(math.Sqrt(float64(area)))
    for d > 0 {
        if area % d == 0 {
            return []int{area/d, d}
        }
        d--
    }
    return nil
}