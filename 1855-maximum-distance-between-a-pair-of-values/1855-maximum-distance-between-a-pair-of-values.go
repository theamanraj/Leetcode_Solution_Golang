func maxDistance(nums1 []int, nums2 []int) int {
    l1 := len(nums1)
    l2 := len(nums2)
    
    p1 := 0
    p2 := 0
    result := 0
    for p1<l1 && p2<l2 {
        if nums1[p1]>nums2[p2] {
            p1++
        } else {
            result = p2-p1
        }
        p2++
    }
    return result
}