func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if nums[mid] == target {
			loi, hii, loii, hiii := mid, mid, mid, mid
			for loi != -1 {
				loii = loi
				loi = binarySearch(nums, 0, loi-1, target)
			}
			for hii != -1 {
				hiii = hii
				hii = binarySearch(nums, hii+1, len(nums)-1, target)
			}

			return []int{loii, hiii}
		}
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, lo, hi, target int) int {
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}