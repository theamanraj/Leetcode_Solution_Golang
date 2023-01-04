func numberOfBoomerangs(points [][]int) int {
    summe := 0

    for i:=0; i < len(points); i++ {
        hx := make(map[int]int)
        for j:=0; j < len(points); j++ {
            if i != j {
                dist := distance(points[i][0] - points[j][0], points[i][1] - points[j][1])
                if _, ok:= hx[dist]; ok {
                    hx[dist] += 1
                } else {
                    hx[dist] = 1
                }
            }
        }
        for _, v := range hx {
            summe += v * (v-1)
        }
        hx = nil
    }
    return summe
}

func distance(x int, y int) int {
    return (x*x + y*y)
}
