func maxSumOfThreeSubarrays(nums []int, k int) []int {
    // 求窗口和
    sum := 0
    // 一共可以有len(nums)-l+1个窗口和
    accArray := make([]int, len(nums)-k+1)
    for i:=0;i<len(nums);i++ {
        sum += nums[i]
        if i+1>k {
            sum -= nums[i-k]
        }
        if i+1>=k {
            accArray[i-k+1] = sum
        }
    }

    left, right := make([]int, len(accArray)), make([]int, len(accArray))
    // 从左到右，求截止第i个位置，左侧最大的累积和数组
    maxAccNum, maxIndex := 0, 0
    for index, val := range accArray {
        if val > maxAccNum {
            maxAccNum = val
            maxIndex = index
        }
        left[index] = maxIndex
    }

    // 求从右到左，截止第i个位置，右侧最大的累积和数组
    maxAccNum, maxIndex = 0, len(accArray)-1
    for i:=len(accArray)-1;i>=0;i-- {
        if accArray[i] >= maxAccNum {
            maxAccNum = accArray[i]
            maxIndex = i
        }
        right[i] = maxIndex
    }

    maxSum, answer := 0, []int{-1, -1, -1}
    // 遍历中间数组，[k, len(nums)-k)，找出左侧和右侧的最大值
    // 左侧满足：leftIndex=i-k
    // 右侧满足： i+k-1<len(accArray)-1 -> i<len(accArray)
    for i:=k;i<len(accArray)-k;i++ {
        midSum := accArray[i]
        leftIndex := left[i-k]
        rightIndex := right[i+k]
        if accArray[leftIndex] + accArray[rightIndex] + midSum > maxSum {
            maxSum = accArray[leftIndex] + accArray[rightIndex] + midSum
            answer[0], answer[1], answer[2] = leftIndex, i, rightIndex
        }
    }
    return answer
}