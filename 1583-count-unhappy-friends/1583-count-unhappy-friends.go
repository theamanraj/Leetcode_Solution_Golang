func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
  pairMap := make(map[int]int)
  for _, pair := range pairs {
    pairMap[pair[0]] = pair[1]
    pairMap[pair[1]] = pair[0]
  }
  better := make([]map[int]bool, n)
  for i := 0; i < n; i++ {
    better[i] = make(map[int]bool)
    for j := 0; j < len(preferences[i]) && preferences[i][j] != pairMap[i]; j++ {
      better[i][preferences[i][j]] = true
    }
  }
  var res int
  for i := 0; i < n; i++ {
    for b := range better[i] {
      if better[b][i] {
        res++
        break
      }     
    }
  }
  return res
}