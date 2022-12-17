func trapRainWater(heightMap [][]int) int {
    var ret int
    jlen := len(heightMap)
    ilen := len(heightMap[0])
    if jlen <3 {
        return 0
    }
    var mark []int = make([]int, ilen*jlen)
    var cellHeap CellHeap
    heap.Init(&cellHeap)
    for i,height := range heightMap[0]{
        var cell = Cell{height:height, i:i, j:0}
        heap.Push(&cellHeap, cell)
        markIndex := i
        mark[markIndex] = -1
    }
    for j:=1; j<len(heightMap)-1; j++{
        jlen := len(heightMap[j])-1
        height1 := heightMap[j][0]
        height2 := heightMap[j][jlen]
        var cell1 = Cell{height:height1, i:0, j:j}
        var cell2 = Cell{height:height2, i:jlen, j:j}
        heap.Push(&cellHeap, cell1)
        heap.Push(&cellHeap, cell2)
        markIndex1 := j*ilen+0
        markIndex2 := j*ilen+jlen
        mark[markIndex1] = -1
        mark[markIndex2] = -1
    }
    for i,height := range heightMap[jlen-1]{
        var cell = Cell{height:height, i:i, j:jlen-1}
        heap.Push(&cellHeap, cell)
        markIndex := (jlen-1)*ilen + i
        mark[markIndex] = -1
    }
    for cellHeap.Len()!=0{
        min := heap.Pop(&cellHeap).(Cell)
        getTrap(min.i+1, min.j, min.height, mark, heightMap, &ret, &cellHeap)
        getTrap(min.i-1, min.j, min.height, mark, heightMap, &ret, &cellHeap)
        getTrap(min.i, min.j+1, min.height, mark, heightMap, &ret, &cellHeap)
        getTrap(min.i, min.j-1, min.height, mark, heightMap, &ret, &cellHeap)
    }
    return ret
}

func getTrap(i int, j int, height int, mark []int, heightMap [][]int, ret *int, cellHeap *CellHeap){
    jlen := len(heightMap)
    ilen := len(heightMap[0])
    if i<0||j<0||i>=ilen||j>=jlen{
        return
    }
    markIndex := j*ilen + i
    if mark[markIndex] != -1{
        trap := height - heightMap[j][i]
        if trap>0 {
            *ret = *ret + trap
            cell := Cell{height:height, i:i, j:j}
            heap.Push(cellHeap, cell)
        }else{
            cell := Cell{height:heightMap[j][i], i:i, j:j}
            heap.Push(cellHeap, cell) 
        }
        mark[markIndex] = -1
    }
}


type Cell struct{
    height int
    i int
    j int
}
type CellHeap []Cell
func (ch CellHeap) Len()int{return len(ch)}
func (ch CellHeap) Less(i, j int)bool{return ch[i].height<=ch[j].height}
func (ch CellHeap) Swap(i, j int){ch[i],ch[j] = ch[j],ch[i]}
func (ch *CellHeap) Push(v interface{}) {*ch = append(*ch, v.(Cell))}
func (ch *CellHeap) Pop()interface{}{
    index := len(*ch) - 1
    ret := (*ch)[index]
    *ch = (*ch)[:index]
    return ret
}