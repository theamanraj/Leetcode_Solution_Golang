func numTilings(n int) int {
  if n <=2 {
  return n
  }
  count, count1, prevCount := 2,2,1
  for i:=3;i<=n;i++ {
    temp := (count+prevCount+count1)%(1e9+7)
    temp1 := (prevCount*2 + count1)%(1e9+7)
    count, count1, prevCount = temp, temp1, count
  }
  return count
}