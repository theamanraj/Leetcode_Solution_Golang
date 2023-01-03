func isPrintable(targetGrid [][]int) bool {
    list := make([]*clrRange, 61)
    
    for i := 0; i < len(targetGrid); i++ {
        for j := 0; j < len(targetGrid[i]); j++ {
            clr := targetGrid[i][j]
            
            if list[clr] == nil {
                list[clr] = &clrRange{clr:clr, x1:61, y1:61}
            }
            list[clr].setRange(i, j)
        }
    }
    
    for {
        state := 1
        for i := 0; i < len(list); i++ {
            if list[i] == nil || list[i].handle {
                continue
            }
            
            if state == 1 {
                state = 2
            }
            if list[i].can(targetGrid) {
                state = 3
                list[i].setGrid(targetGrid)
            }
        }
        
        if state == 2 {
            return false
        } else if state == 1 {
            return true
        }
    }
    return true
}

type clrRange struct {
    clr int
    x1, x2, y1, y2 int
    handle bool
}

func (c *clrRange) setRange(x, y int) {
    if c.x1 > x {
        c.x1 = x
    }
    if c.x2 < x {
        c.x2 = x
    }
    if c.y1 > y {
        c.y1 = y
    }
    if c.y2 < y {
        c.y2 = y
    }
}

func (c *clrRange) can (grid [][]int) bool {
    for x := c.x1; x <= c.x2; x++ {
        for y := c.y1; y <= c.y2; y++ {
            if grid[x][y] == 0 || grid[x][y] == c.clr {
                continue
            }
            return false
        }
    }
    return true
}

func (c *clrRange) setGrid(grid [][]int) {
    for x := c.x1; x <= c.x2; x++ {
        for y := c.y1; y <= c.y2; y++ {
            grid[x][y] = 0
        }
    }
    c.handle = true
}