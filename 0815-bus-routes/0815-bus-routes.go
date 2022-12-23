func numBusesToDestination(routes [][]int, source int, target int) int {
    if source==target{ return 0 }    
    m := make(map[int][]int)    // key is stop, value is buses 
    for bus, ss := range routes{        
        for _, station := range ss{                        
            m[station]=append(m[station], bus) // build the map!
        }
    }        
    var busCount int
    curStops := []int{source} // Queue for BFS      
    visited := make([]int, len(routes))     // orvisited := map[int]bool{} // to record the visited bus.       
    for len(curStops)>0{  // Typical BFS template.      
        var next []int    // next layer 
        busCount++        // need 
        for _, station := range curStops{  // iterate current level
            for _, bus := range m[station]{   // all the buses of this level
                if visited[bus]==1{ continue}  // if we have taken this bus, need try next bus
                visited[bus]=1                 
                for _, s := range routes[bus]{  // With the 2 for, get the all connected stations.
                    if s==target { 
                        return busCount
                    }
                    next=append(next,s)       // all the connected stations  are stations of next level.            
                }
            }
        }        
        curStops = next        
    }
    return -1    
}
