func containsDuplicate(nums []int) bool {
    var mymap = make(map[int]bool)
    for _,i := range nums{
        if !mymap[i]{
        mymap[i] = true  
        } else {
        return true
        }
    }
    return false
}