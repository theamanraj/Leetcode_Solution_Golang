func merge(nums1 []int, m int, nums2 []int, n int)  {
    copy(nums1[m:], nums2[:n])
    sort.Ints(nums1)
}