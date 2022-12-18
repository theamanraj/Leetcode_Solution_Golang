type Solution struct {
    original []int
}

func Constructor(nums []int) Solution {
    rand.Seed(time.Now().UnixNano())
    s := Solution {
        original: nums,
    }
    return s
}


func (this *Solution) Reset() []int {
    return this.original
}

func (this *Solution) Shuffle() []int {
	arr := make([]int, len(this.original))
	copy(arr, this.original)
	for i := 0; i < len(arr); i++ {
		random := rand.Intn(len(arr))
		arr[i], arr[random] = arr[random], arr[i]
	}

	return arr
}