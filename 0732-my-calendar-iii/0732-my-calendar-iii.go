type MyCalendarThree struct {
    vals map[int]int
    lazy map[int]int
}

func Constructor() MyCalendarThree {
    vals := make(map[int]int)
    lazy := make(map[int]int)
    
    return MyCalendarThree{
        vals: vals,
        lazy: lazy,
    }
}

func (this *MyCalendarThree) Book(start int, end int) int {
    this.Update(start, end-1, 0, int(math.Pow(10,9)), 1)
    return this.vals[1]
}

func (this *MyCalendarThree) Update(start int, end int, left int, right int, idx int) {
    if start > right || end < left {
        return
    }
    
    if start <= left && right <= end {
        this.vals[idx]++
        this.lazy[idx]++
    } else {
        mid := (right - left) / 2 + left
        this.Update(start, end, left, mid, idx * 2)
        this.Update(start, end, mid+1, right, idx * 2 + 1)
        this.vals[idx] = this.lazy[idx] + max(this.vals[2*idx], this.vals[2*idx+1])
    }    
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}