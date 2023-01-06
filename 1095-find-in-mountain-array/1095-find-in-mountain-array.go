/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    l := 0
    r := mountainArr.length() - 1
    count := 0
    peak := -1

    //step 1 
    for l + 1 < r {
        count ++
        mid := l + (r - l) / 2
        midVal := mountainArr.get(mid)
        midLeftVal := mountainArr.get(mid - 1)
        midRightVal := mountainArr.get(mid + 1)
        
        if midVal > midLeftVal && midVal > midRightVal {
            peak = mid
            break
        } 
        if midVal > midLeftVal && midVal < midRightVal {
            l = mid
        }
        if midVal < midLeftVal && midVal > midRightVal {
            r = mid
        }
    }
    if peak == -1 {
        if mountainArr.get(l) > mountainArr.get(r) {
            peak = l
        } else {
            peak = r
        }
    }
    
    //step 2 
    l = 0
    r = peak
    for l + 1 < r {
        mid := l + (r - l) / 2
        midVal := mountainArr.get(mid)
        
        if midVal == target {
            return mid
        }
        if midVal > target {
            r = mid
        } else {
            l = mid
        }
    }
    if mountainArr.get(l) == target {
        return l
    }
    if mountainArr.get(r) == target {
        return r
    }
    
    //step 3 
    l = peak
    r = mountainArr.length() - 1
    for l + 1 < r {
        mid := l + (r - l) / 2
        midVal := mountainArr.get(mid)
        if midVal == target {
            return mid
        }
        if midVal > target {
            l = mid
        } else {
            r = mid
        }
    }
    if mountainArr.get(l) == target {
        return l
    }
    if mountainArr.get(r) == target {
        return r
    }
    
    
    return -1
}