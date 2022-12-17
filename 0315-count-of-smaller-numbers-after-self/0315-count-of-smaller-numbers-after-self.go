func merge(l1, h1, l2, h2 int, a, oldIdx, count, newIdx []int) {
    start := l1
    end := h2
    
    smaller := 0 // Count of smaller numbers in the right sorted slice
    for l1 <= h1 && l2 <= h2 {
        if a[oldIdx[l1]] > a[oldIdx[l2]] {
            smaller++
            newIdx = append(newIdx, oldIdx[l2])
            l2++
        } else {
            count[oldIdx[l1]] += smaller
            newIdx = append(newIdx, oldIdx[l1])
            l1++
        }
    }
    
    for l1 <= h1 {
        count[oldIdx[l1]] += smaller
        newIdx = append(newIdx, oldIdx[l1])
        l1++
    }
    
    for l2 <= h2 {
        newIdx = append(newIdx, oldIdx[l2])
        l2++
    }
    
    copy(oldIdx[start:end+1], newIdx)
}

func mergeSort(l, h int, a, index, count, newIdxs []int) {
    if l >= h {
        return
    }
    
    mid := (l+h)/2
    mergeSort(l, mid, a, index, count, newIdxs)
    mergeSort(mid+1, h, a, index, count, newIdxs)
    
    newIdxs = newIdxs[:0]  //Reset the new idx slice so that it is ready for next merge
    merge(l, mid, mid+1, h, a, index, count, newIdxs)
}

func countSmaller(a []int) []int {
    l := len(a)
    idxs := make([]int, l, l)
    count := make([]int, l, l)
    for i := 0; i < l; i++ {
        idxs[i] = i
    }
    newIdx := make([]int, 0, len(idxs)) // Re-use the same slice to avoid memory allocation on each call
    mergeSort(0, l-1, a, idxs, count, newIdx)
    
    return count
}