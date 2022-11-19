func climbStairs(n int) int {
if n==1{
return 1
}else if n==2{
return 2
}
i:=1
j:=2
k:=3
var result int
for ;k<=n;k++{
result = i+j
i=j
j=result
}

return result
}