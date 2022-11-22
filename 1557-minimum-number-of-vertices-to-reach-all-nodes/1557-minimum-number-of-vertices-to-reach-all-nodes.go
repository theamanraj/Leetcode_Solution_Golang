func findSmallestSetOfVertices(n int, edges [][]int) []int {
  var ans []int
  
  s := make(map[int]bool)
  r := make(map[int]bool)
  
  for _, i := range edges {
    s[i[1]] = true
  }
  
  for _, i := range edges {
    if ok, _ := s[i[0]]; !ok {
      if ok, _ := r[i[0]]; !ok {
        ans = append(ans, i[0])
        r[i[0]] = true
      }
    }
  }
  
  return ans
}