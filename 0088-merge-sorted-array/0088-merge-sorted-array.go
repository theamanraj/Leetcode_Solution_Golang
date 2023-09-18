func merge(nums1 []int, m int, nums2 []int, n int)  {
    i,j,l := m-1,n-1,m+n-1
    
    //execute loop till last index reach one of array
    for i >= 0 && j >= 0{
        // validate which num is better nd add it in last
        if nums1[i] > nums2[j] {
            nums1[l] = nums1[i]
            i--
        } else {
            nums1[l] = nums2[j]
            j--
        }
        // reduce total size
        l--
    }
    
        for j >= 0 {
        nums1[l] = nums2[j]
        l--
        j--
    }
}