func closestMeetingNode(edges []int, node1 int, node2 int) int {
    
    dist1:=make([]int,len(edges))
    dist2:=make([]int,len(edges))
    
    d:=1
    for node1!=-1 {
        // detect cycle
        if dist1[node1]>0{
            break
        }
        // write current distance
        dist1[node1]=d
        d+=1
        node1=edges[node1]
    }
    d=1
    for node2!=-1 {
        if dist2[node2]>0{
            break
        }
        dist2[node2]=d
        d+=1
        node2=edges[node2]
    }
    
    res:=-1
    gmx:=1<<31
    
    for i:=0;i<len(dist1);i++{
        if dist1[i]>0 && dist2[i]>0{
            mx:= max(dist1[i],dist2[i])
            if mx<gmx{
                gmx=mx
                res=i
            }
        }
    }
    
    return res
}

func max(a,b int) int{
    if a>b{
        return a
    }
    return b
}