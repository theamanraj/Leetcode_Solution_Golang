type MinStack struct {
    nums []node
}

func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(val int)  {
    if len(this.nums)==0{
        this.nums=append(this.nums, node{val: val,min: val})  
        return
    }
    
    min := this.nums[len(this.nums)-1].min
    if val<min{ min=val }    
    this.nums=append(this.nums, node{val: val,min: min})    
}


func (this *MinStack) Pop()  {
    this.nums= this.nums[:len(this.nums)-1]
}


func (this *MinStack) Top() int {
    return this.nums[len(this.nums)-1].val
}


func (this *MinStack) GetMin() int {
    return this.nums[len(this.nums)-1].min
}

type node struct{
    min int
    val int
}
