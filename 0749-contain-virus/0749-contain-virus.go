func containVirus(isInfected [][]int) int {
  m, n := len(isInfected), len(isInfected[0])
  var res int
  for {
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
      visited[i] = make([]bool, n)
    }
    var wallsNeeded [][4]int 
    for i := 0; i < m; i++ {
      for j := 0; j < n; j++ {
        if isInfected[i][j] == 1 && !visited[i][j] {
          visited[i][j] = true
          uninfectedNeighbors := make(map[[2]int]bool)
          count := countWalls(isInfected, visited, i, j, uninfectedNeighbors)
          wallsNeeded = append(wallsNeeded, [4]int{i, j, len(uninfectedNeighbors), count})
        }
      }
    }
    if len(wallsNeeded) == 0 {
      break
    }
    sort.Slice(wallsNeeded, func(i, j int) bool {
      return wallsNeeded[i][2] > wallsNeeded[j][2]
    })
    res += wallsNeeded[0][3]
    buildWall(isInfected, wallsNeeded[0][0], wallsNeeded[0][1])
    spread(isInfected)
  }
  return res
}

var directions = [4][2]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

func countWalls(isInfected [][]int, visited [][]bool, row, col int, uninfectedNeighbors map[[2]int]bool) int {
  m, n := len(isInfected), len(isInfected[0])
  var walls int
  for _, dir := range directions {
    x, y := row + dir[0], col + dir[1]
    if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] {
      if isInfected[x][y] == 0 {
        uninfectedNeighbors[[2]int{x, y}] = true
        walls++
      } else if isInfected[x][y] == 1 {
        visited[x][y] = true
        walls += countWalls(isInfected, visited, x, y, uninfectedNeighbors)
      }
    }
  }
  return walls
}

func buildWall(isInfected [][]int, row, col int) {
  m, n := len(isInfected), len(isInfected[0])
  isInfected[row][col] = -1
  for _, dir := range directions {
    x, y := row+dir[0], col+dir[1]
    if x >= 0 && x < m && y >= 0 && y < n && isInfected[x][y] == 1 {
      buildWall(isInfected, x, y)
    }
  }
}

func spread(isInfected [][]int) {
  m, n := len(isInfected), len(isInfected[0])
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      if isInfected[i][j] == 1 {
        for _, dir := range directions {
          x, y := i+dir[0], j+dir[1]
          if x >= 0 && x < m && y >= 0 && y < n && isInfected[x][y] == 0 {
            isInfected[x][y] = 2 
          }
        }
      }
    }
  }
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      if isInfected[i][j] == 2 {
        isInfected[i][j] = 1
      }
    }
  }
}