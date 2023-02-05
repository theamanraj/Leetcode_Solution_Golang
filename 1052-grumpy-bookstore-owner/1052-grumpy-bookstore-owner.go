func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    var happy int
  var res int
  for i := 0; i < len(customers); i++ {
    if grumpy[i] == 0 {
      happy += customers[i]
    }
  }
  var extra int
  for i := 0; i < len(customers); i++ {
    if grumpy[i] == 1 {
      extra += customers[i]
    }
    res = max(res, happy + extra)
    if i >= minutes-1 {
      if grumpy[i-minutes+1] == 1 {
        extra -= customers[i-minutes+1]
      }
    }
  }
  return res
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}