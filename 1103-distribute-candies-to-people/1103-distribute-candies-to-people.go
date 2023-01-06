func distributeCandies(candies int, num_people int) []int {
ret := make([]int,num_people)
sum := 0
a := 0
w := 0
n := num_people
start := 1

for j:=1;j<=n;j++{
	w+=j
}
for sum < candies{
	a = w+ (start-1)*n*n
	sum+=a
	start++
}
start--

left := candies - (sum -a)
for i:= 0; i< num_people;i++{
	temp := (start-1)*n +i+1
	if left >=temp{
		ret[i] = start*(i+1) + start*(start-1)*n/2
		left = left - temp

	}else{
		ret[i] = (start-1)*(i+1) + (start-1)*(start-2)*n/2
		ret[i] += left
        left = 0
	}
}

return ret
}