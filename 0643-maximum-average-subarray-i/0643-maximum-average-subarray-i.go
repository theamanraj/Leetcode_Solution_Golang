package main

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	var max float64
	for i := 0; i < len(nums); i++ {
		if i < k {
			sum += nums[i]
			max = float64(sum) / float64(i+1)
		} else {
			sum = sum - nums[i-k] + nums[i]
			if float64(sum)/float64(k) > max {
				max = float64(sum) / float64(k)
			}
		}
	}
	return max
}