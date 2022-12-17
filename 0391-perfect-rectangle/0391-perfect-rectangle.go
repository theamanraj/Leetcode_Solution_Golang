func isRectangleCover(rects [][]int) bool {
    
    type pt struct { x, y int }
    
    var (
        left, right, top, bottom int
        area int    
        
        set = make(map[pt]int)
    )    
    
    for i, r := range rects {
        x, y, xx, yy := r[0], r[1], r[2], r[3]
        if i == 0 || x < left {
            left = x
        }
        if i == 0 || y < bottom {
            bottom = y
        }
        if i == 0 || xx > right {
            right = xx
        }
        if i == 0 || yy > top {
            top = yy
        }
        area += (xx-x)*(yy-y)
        
        set[pt{x,y}]++
        if set[pt{x,y}] == 2 {
            delete(set, pt{x,y})
        }
        set[pt{x,yy}]++
        if set[pt{x,yy}] == 2 {
            delete(set, pt{x,yy})
        }
        set[pt{xx,y}]++
        if set[pt{xx,y}] == 2 {
            delete(set,pt{xx,y})
        }
        set[pt{xx,yy}]++
        if set[pt{xx,yy}] == 2 {
            delete(set, pt{xx,yy})
        }
    }
    
    return len(set) == 4 &&
        set[pt{left, bottom}] == 1 &&
        set[pt{left, top}] == 1 &&
        set[pt{right, bottom}] == 1 &&
        set[pt{right, top}] == 1 &&
        area == (right-left)*(top-bottom)    
    
}