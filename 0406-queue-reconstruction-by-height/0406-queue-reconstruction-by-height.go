type Fenwick struct {
    Size int
    Tree []int
}

func NewFenwick(size int) *Fenwick {
    f := &Fenwick{}
    f.Size = size
    f.Tree = make([]int, size+1)
    return f
}

func (f *Fenwick) Update(i, v int) {
    for i < f.Size {
        f.Tree[i] += v
        i += (i & -i)
    }
}

func (f *Fenwick) Query(i int) int {
    total := 0
    for i > 0 {
        total += f.Tree[i]
        i -= (i & -i)
    }
    return total
}


func reconstructQueue(people [][]int) [][]int {
    n := len(people)
    queue := make([][]int, n)
    
    sort.Slice(people, func(i,j int) bool {
        if people[i][0] < people[j][0] {
            return true
        } else if people[i][0] > people[j][0] {
            return false
        } else {
            return people[i][1] > people[j][1]
        }
    })
    
    
    fenwick := NewFenwick(n)
    for i := 1; i < n+1; i++ {
        fenwick.Update(i, 1)
    }
    
    for _, p := range people {
        h, k := p[0], p[1]
        lo, hi, index := 1, n-1, 1
        for lo <= hi {
            mid := lo + (hi - lo) / 2
            if fenwick.Query(mid) <= k {
                lo = mid + 1
                index = mid + 1
            } else {
                hi = mid - 1
            }
        }
        queue[index-1] = []int{h, k}
        fenwick.Update(index, - 1)
    }
  
    return queue
}