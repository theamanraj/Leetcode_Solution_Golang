func maxFrequency(nums []int, k int) int {
    hash := make(map[int]int)
    for _, v := range nums {
        hash[v]++
    }
    nums = make([]int, 0, len(hash))
    for k := range hash {
        nums = append(nums, k)
    }
    sort.Ints(nums)
    count := 0
    freq := hash[nums[0]]
    prev, prevSize := 0, hash[nums[0]]
    for i := 1; i < len(nums); i++ {
        count += hash[nums[i-1]]
        k -= count * (nums[i]-nums[i-1])
        for k < 0 {
            // prev move right until k is not negative
            k += (nums[i] - nums[prev]) * prevSize
            count -= prevSize
            prev++
            prevSize = hash[nums[prev]]
        }
        if k > 0 && prev > 0 && k / (nums[i] - nums[prev-1]) > 0 {
            // prev move left
            prev--
            prevSize = k / (nums[i] - nums[prev])
            if prevSize > hash[nums[prev]] {
                prevSize = hash[nums[prev]]
            }
            k -= prevSize * (nums[i] - nums[prev])
            count += prevSize
        }
        if currFreq := hash[nums[i]] + count; currFreq > freq {
            freq = currFreq
        }
    }
    return freq
}