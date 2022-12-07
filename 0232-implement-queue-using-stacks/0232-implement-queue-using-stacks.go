type MyQueue struct {
    stackPush, stackPop []int
}

func Constructor() MyQueue {
    return MyQueue{}
}

func (this *MyQueue) Push(x int)  {
    this.stackPush = append(this.stackPush, x)
}

func (this *MyQueue) Pop() int {
    ans := this.Peek()
    this.stackPop = this.stackPop[:len(this.stackPop)-1]
    return ans
}

func (this *MyQueue) Peek() int {
    if len(this.stackPop) == 0 {
        for len(this.stackPush) != 0 {
            this.stackPop = append(this.stackPop, this.stackPush[len(this.stackPush)-1])
            this.stackPush = this.stackPush[:len(this.stackPush)-1]
        }
    }
    return this.stackPop[len(this.stackPop)-1]
}

func (this *MyQueue) Empty() bool {
    return len(this.stackPush) == 0 && len(this.stackPop) == 0
}