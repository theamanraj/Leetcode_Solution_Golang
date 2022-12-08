func nearestExit(maze [][]byte, entrance []int) int {
    q := make([][]int,0)
    q=append(q, []int{entrance[0], entrance[1],0})
    for len(q) > 0 {
        i, j, ans  := q[0][0], q[0][1], q[0][2]
        if !(i==entrance[0] && j==entrance[1]) && (i==0 || j==0 || i==len(maze)-1 || j==len(maze[0])-1) {
            return ans
        }
        dir := [][]int{{1,0},{0,1},{0,-1},{-1,0}}
        for _, v := range dir {
            k, l := i+v[0], j+v[1]
            if k>=0 && l>=0 && k<len(maze) && l<len(maze[0]) && maze[k][l] == '.' {
                maze[k][l] = '+'
                q = append(q, []int{k,l,ans+1})
            }
        }
        q = q[1:]
    }
    return -1
}