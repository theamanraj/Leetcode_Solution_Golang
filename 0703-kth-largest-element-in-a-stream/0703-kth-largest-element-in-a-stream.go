type IntSlice []int
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }
type KthLargest struct {
	k int
	nums []int
    largest []int
}

func Constructor(k int, nums []int) KthLargest {
	var ele KthLargest = KthLargest{
		k:k,
		nums:nums,
	}
	sort.Sort(IntSlice(ele.nums))
    if len(nums) > k{
        ele.largest = make([]int, k, k)
        copy(ele.largest,ele.nums[0:k])
    }else{
        ele.largest = make([]int, len(nums), len(nums)) 
        copy(ele.largest,ele.nums[0:])
    }
	return ele
}

func (this *KthLargest) Add(val int) int {
    if len(this.nums) < this.k{
        this.nums = append(this.nums, val)
        this.largest = append(this.largest,val)
        sort.Sort(IntSlice(this.largest))
        return this.largest[len(this.largest) - 1]
    }else{
        if val > this.largest[this.k - 1]{
            this.largest = append(this.largest, val)
            sort.Sort(IntSlice(this.largest))
            this.largest = this.largest[0:this.k]
        }
        return this.largest[this.k - 1]
    }	
}