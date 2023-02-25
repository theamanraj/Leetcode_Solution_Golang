func maxChunksToSorted(arr []int) int {
    st := &iStack{}
    
    for i := 0; i < len(arr); i++ {
        m1, m2 := arr[i], arr[i]
        
        for st.Len() > 0 {
            it := st.Top()
            if it.max <= m1 {
                break
            }
            
            st.Pop()
            m1 = minInt(m1, it.min)
            m2 = maxInt(m2, it.max)
        }
        st.Push(&item{min:m1, max:m2})
    }
    
    return st.Len()
}

func minInt(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func maxInt(x, y int) int {
    if x > y {
        return x
    }
    return y
}

type item struct {
    min, max int
}

type iStack struct {
    arr []*item
}

func (s *iStack) Len() int {
    return len(s.arr)
}

func (s *iStack) Push(it *item) {
    s.arr = append(s.arr, it)
}

func (s *iStack) Pop() {
    s.arr = s.arr[:len(s.arr)-1]
}

func (s *iStack) Top() *item {
    return s.arr[len(s.arr)-1]
}