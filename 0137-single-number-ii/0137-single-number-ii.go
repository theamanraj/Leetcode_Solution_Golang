func singleNumber(nums []int) int {
    ones := uint(0); twos := uint(0);
    for i := uint(0); i < uint(len(nums)); i++ {
        ones = (ones ^ uint(nums[i])) &^ twos
        twos = (twos ^ uint(nums[i])) &^ ones
    }
    return int(ones)
}