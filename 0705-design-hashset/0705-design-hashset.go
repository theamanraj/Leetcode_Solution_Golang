const (
    Buckets = 1024
)

var hashFn = func(x int) int {
	return x * 2021
}

type Node struct {
    Val   int
    Next  *Node
}

type HashFunc func(int)int

type MyHashSet struct {
    node   []*Node
    size   int
    fn     HashFunc
}

func Constructor() MyHashSet {
    return MyHashSet{
        node:   make([]*Node, Buckets),
        size:   Buckets,
        fn:     hashFn,
    }
}


func (this *MyHashSet) Add(key int)  {
    hashed := this.fn(key) % this.size
    if this.node[hashed] == nil {
        this.node[hashed] = &Node { Val: key, }
    } else {
        prev, cur := this.node[hashed], this.node[hashed]
        for cur != nil && cur.Val != key {
            prev = cur
            cur = cur.Next
        }
        if cur == nil {
            prev.Next = &Node{ Val: key }
        }
    }
}


func (this *MyHashSet) Remove(key int)  {
    hashed := this.fn(key) % this.size
    if this.node[hashed] == nil { return }
    if this.node[hashed].Val == key {
        this.node[hashed] = this.node[hashed].Next
        return
    }
    prev, cur := this.node[hashed], this.node[hashed]
    for cur != nil && cur.Val != key {
        prev = cur
        cur = cur.Next
    }
    if cur != nil {
        prev.Next = cur.Next
    }
}

func (this *MyHashSet) Contains(key int) bool {
    hashed := this.fn(key) % this.size
    cur := this.node[hashed]
    for cur != nil && cur.Val != key {
        cur = cur.Next
    }
    return cur != nil && cur.Val == key
}