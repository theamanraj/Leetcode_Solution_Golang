import "strconv"

func subtractProductAndSum(n int) int {
    stringN := strconv.Itoa(n)
    intSlice := []int{}
    
    for _,val := range stringN{
        intVal := int(val)-48
        intSlice = append(intSlice,intVal)
    }
    return product(intSlice) - sum(intSlice)
}

//returns the product of a slice of ints 
func product(x []int) int {
    start := 1
    for _,val := range x{
        start = start*val
    }
    return start
}

//returns the sum of a slice of ints
func sum(y []int) int {
    start := 0
    for _,val := range y {
        start = start+val
    }
    return start
}