type Tree struct {
    r, c int
}

func outerTrees(trees [][]int) [][]int {
    hull := make(map[Tree]struct{})
    n := len(trees)
    
    if n < 4 {
        return trees
    }
    
    leftMost := 0
    for i, tree := range trees {
        if tree[0] < trees[leftMost][0] {
            leftMost = i
        }
    }
    
    p := leftMost
    for {
        q := (p + 1) % n
        for i, _ := range trees {
            if orientation(trees[p], trees[i], trees[q]) < 0 {
                q = i
            }
        }
        
        for i, _ := range trees {
            if (i != q && i != p && orientation(trees[p], trees[i], trees[q]) == 0 && inBetween(trees[p], trees[i], trees[q])) {
                hull[Tree{trees[i][0], trees[i][1]}] = struct{}{}
            }
        }
        
        hull[Tree{trees[q][0], trees[q][1]}] = struct{}{}
        p = q
        
        if _, ok := hull[Tree{trees[leftMost][0], trees[leftMost][1]}]; ok{
            break    
        }
    }
    
    output := make([][]int, 0)
    for k, _ := range hull {
        output = append(output, []int{k.r, k.c})
    }
    return output
}

func orientation(p []int, q []int, r []int) int {
    return (q[1] - p[1]) * (r[0] - q[0]) - (q[0] - p[0]) * (r[1] - q[1])
}

func inBetween(p []int, i []int, q []int) bool {
    a := (i[0] >= p[0] && i[0] <= q[0]) || (i[0] <= p[0] && i[0] >= q[0])
    b := (i[1] >= p[1] && i[1] <= q[1]) || (i[1] <= p[1] && i[1] >= q[1])
    return a && b
}