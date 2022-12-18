func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
items := map[int]int{}
res := [][]int{}
for _, v1 := range items1 {
items[v1[0]] += v1[1]
}
for _, v1 := range items2 {
items[v1[0]] += v1[1]
}

for val, weight := range items{
    res = append(res, []int{val, weight})
}
sort.SliceStable(res, func(i, j int) bool {
    return res[i][0] < res[j][0]
})

return res
}