func search(a []int, target int) bool {
    lo, hi := 0, len(a) - 1
    for lo <= hi {
        mid := (lo + hi) / 2
        if a[mid] == target || target == a[lo] || target == a[hi] {
            return true
        }
        if a[lo] < a[mid] && target < a[lo] {
            lo = mid + 1
        } else if a[lo] < a[mid] && target > a[mid] {
            lo = mid + 1
        } else if a[lo] < a[mid] && target > a[lo] {
            hi = mid - 1
        } else if a[mid] < a[hi] && target < a[mid] {
            hi = mid - 1
        } else if a[mid] < a[hi] && target > a[hi] {
            hi = mid - 1
        } else if a[mid] < a[hi] && target > a[mid] {
            lo = mid + 1
        } else if a[lo] == a[mid] || a[mid] == a[hi] {
            for _, v := range a[lo:hi+1] {
                if v == target {
                    return true
                }
            }
            return false
        }
    }
    return false
}