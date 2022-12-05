func checkStraightLine(coordinates [][]int) bool {
    if len(coordinates)==2{
        return true
    }
    
    xGap:=(coordinates[1][0]-coordinates[0][0])
    yGap:=(coordinates[1][1]-coordinates[0][1])
    for idx:=2;idx<len(coordinates);idx++{
        if yGap*(coordinates[idx][0]-coordinates[0][0])!=(coordinates[idx][1]-coordinates[0][1])*xGap{
            return false
        }
    }
    
    return true
}