func guessNumber(n int) int {
    left := 1
    right := n
    
    for { // no need to test if left <= right, because it's impossible that we cann't find the target
        mid := int(uint(left+right)>>1)
        switch guess(mid) {
        case 0:
            return mid
        case -1:
            right = mid-1
        case 1:
            left = mid+1
        }
    }
    
    return -1
}