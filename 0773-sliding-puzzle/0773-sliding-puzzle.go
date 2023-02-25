type Pos struct {
    PosOf0 int8
    PosOf1 int8
    PosOf2 int8
    PosOf3 int8
    PosOf4 int8
    PosOf5 int8
}

func isWishedState(p *Pos) bool {
    return p.PosOf0 == 5 && p.PosOf1 == 0 && p.PosOf2 == 1 && p.PosOf3 == 2 && p.PosOf4 == 3 &&  p.PosOf5 == 4
}

func calcIdx(r int, c int) int8 {
    return int8(r*3 + c)
}

func isAdjacent(t int8, standard int8) bool {
    return (standard != 3 && t == standard -1) || (standard != 2 && t == standard + 1) || t == standard - 3 || t == standard + 3
}


func slidingPuzzle(board [][]int) int {    
    isVisited := make(map[Pos]struct{})
    
    var init Pos
    for r, row := range board {
        for c, cell := range row {
            switch cell {
                case 0:
                    init.PosOf0 = calcIdx(r,c)
                case 1:
                    init.PosOf1 = calcIdx(r,c)
                case 2:
                    init.PosOf2 = calcIdx(r,c)
                case 3:
                    init.PosOf3 = calcIdx(r,c)
                case 4:
                    init.PosOf4 = calcIdx(r,c)
                case 5:
                    init.PosOf5 = calcIdx(r,c)
            }
        }
    } 
    
    curr := []Pos{init}
    isVisited[init] = struct{}{}
    step := 0
    
    for len(curr) != 0 { 
        next := []Pos{}
        for _, p := range curr {
            if isWishedState(&p) {
                return step
            }
            
            if(isAdjacent(p.PosOf1, p.PosOf0)) {
                nextPos := Pos{p.PosOf1, p.PosOf0, p.PosOf2, p.PosOf3, p.PosOf4, p.PosOf5}
                _, ok := isVisited[nextPos]
                if !ok {
                    isVisited[nextPos] = struct{}{}
                    next = append(next, nextPos)
                }
            }
            if(isAdjacent(p.PosOf2, p.PosOf0)) {
                nextPos := Pos{p.PosOf2, p.PosOf1, p.PosOf0, p.PosOf3, p.PosOf4, p.PosOf5}
                _, ok := isVisited[nextPos]
                if !ok {
                    isVisited[nextPos] = struct{}{}
                    next = append(next, nextPos)
                }
            }
            if(isAdjacent(p.PosOf3, p.PosOf0)) {
                nextPos := Pos{p.PosOf3, p.PosOf1, p.PosOf2, p.PosOf0, p.PosOf4, p.PosOf5}
                _, ok := isVisited[nextPos]
                if !ok {
                    isVisited[nextPos] = struct{}{}
                    next = append(next, nextPos)
                }
            }
            if(isAdjacent(p.PosOf4, p.PosOf0)) {
                nextPos := Pos{p.PosOf4, p.PosOf1, p.PosOf2, p.PosOf3, p.PosOf0, p.PosOf5}
                _, ok := isVisited[nextPos]
                if !ok {
                    isVisited[nextPos] = struct{}{}
                    next = append(next, nextPos)
                }
            }
            if(isAdjacent(p.PosOf5, p.PosOf0)) {
                nextPos := Pos{p.PosOf5, p.PosOf1, p.PosOf2, p.PosOf3, p.PosOf4, p.PosOf0}
                _, ok := isVisited[nextPos]
                if !ok {
                    isVisited[nextPos] = struct{}{}
                    next = append(next, nextPos)
                }
            }
        }
        
        curr = next
        step++
    }
    
    return -1
}