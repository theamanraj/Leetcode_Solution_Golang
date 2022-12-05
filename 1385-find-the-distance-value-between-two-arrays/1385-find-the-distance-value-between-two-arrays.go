func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    var res int
    for _,a1:= range arr1 {
        count:=0
        for _, a2:= range arr2 {
            if int(math.Abs(float64(a1-a2))) <= d {
                break
            }
            count++ 
        }
        if count == len(arr2) {
            res++
        }
    } 
    return res
}