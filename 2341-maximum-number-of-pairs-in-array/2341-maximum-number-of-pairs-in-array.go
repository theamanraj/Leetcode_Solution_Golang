func numberOfPairs(nums []int) []int {
m:=make(map[int]int)
for i:=0;i<len(nums);i++{
cur:=nums[i]
m[cur]=m[cur]+1
}
a:=0
b:=0
for _,val:=range m{
a+=val/2
b+=val%2
}
return []int{a,b}
}

