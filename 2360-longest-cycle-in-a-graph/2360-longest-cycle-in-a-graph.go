func longestCycle(edges []int) int {
    n:=len(edges)
    visited:=make([]int,n)
    max:=-1
    for i:=0;i<n;i++{
        if visited[i]!=0 {
            continue
        }
        j:=i
        for visited[j]==0 {
            visited[j] = i+1
            j=edges[j]
            if j==-1 {
                break
            }
            if visited[j]==i+1 {
                l:=1
                k:=edges[j]
                for k!=j {
                    k=edges[k]
                    l++
                }
                if l>max{
                    max = l
                }
                break
            }
        }
    }
    return max
}