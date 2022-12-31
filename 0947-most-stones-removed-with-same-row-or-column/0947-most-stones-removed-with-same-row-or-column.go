func find(graph map[string]string,coordinate string)string{
  if graph[coordinate] == coordinate {
    return coordinate
  }
  graph[coordinate] = find(graph,graph[coordinate])
  return graph[coordinate]
}

func union(graph map[string]string,a string,b string)bool{
  rootA := find(graph,a)
  rootB := find(graph,b)
  if rootA!=rootB {
    graph[rootA] = rootB
    return true
  }
  return false
}

func getKey(x int,y int)string{
  return fmt.Sprintf("%v:%v",x,y)
}

func removeStones(stones [][]int) int {
  rowsMap:=make(map[int][][]int)
  colMap:=make(map[int][][]int)
  graph := make(map[string]string)
  for _, val := range stones {
    x := val[0]
    y := val[1]
    coordinate:= getKey(x,y)
    rowsMap[x] = append(rowsMap[x],val)
    colMap[y] = append(colMap[y],val)
    graph[coordinate] = coordinate
  }
  conn:=0
  for _, val := range stones{
    x:=val[0]
    y:=val[1]
    coordinate := getKey(x,y)
    for _, rowNeighbour := range rowsMap[x]{
      coord := getKey(rowNeighbour[0],rowNeighbour[1])
      if union(graph,coordinate,coord){
        conn++
      }
    }
    for _, colNeighbour := range colMap[y]{
      coord := getKey(colNeighbour[0],colNeighbour[1])
      if union(graph,coordinate,coord){
        conn++
      }
    } 
  }
  return conn
}