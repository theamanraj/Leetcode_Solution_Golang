type BrowserHistory struct {
    curr *DoubleNode
}

func Constructor(homepage string) BrowserHistory {
    homePageNode := &DoubleNode {
        Val : homepage,
    }
    
    return BrowserHistory {
        curr : homePageNode, 
    }
}

func (this *BrowserHistory) Visit(url string)  {
    visitingNode := &DoubleNode {
        Val : url,
        Prev : this.curr,
    }
    
    this.curr.Next = visitingNode
    this.curr = this.curr.Next 
}


func (this *BrowserHistory) Back(steps int) string {
    for i := 0; i < steps; i++ {
        if this.curr.Prev != nil {
            this.curr = this.curr.Prev
        } 
    }
    return this.curr.Val
}


func (this *BrowserHistory) Forward(steps int) string {
        for i := 0; i < steps; i++ {
        if this.curr.Next != nil {
            this.curr = this.curr.Next
        } 
    }
    return this.curr.Val
}


type DoubleNode struct {
    Val string
    Next *DoubleNode
    Prev *DoubleNode    
}

