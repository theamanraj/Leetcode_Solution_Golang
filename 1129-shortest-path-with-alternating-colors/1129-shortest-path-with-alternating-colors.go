import (
	"math"
)

type Node struct {
	vertex int
	color  byte
}

var RED, BLUE = 0, 1

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {

	// convert red and blue edges into respective adjacent matrices
	var adj [2][][]int
	adj[RED] = buildAdj(n, redEdges)
	adj[BLUE] = buildAdj(n, blueEdges)

	//  initialize distance array, that stores distances as we proceed along with bfs
	var dist [2][]int
	dist[RED], dist[BLUE] = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		dist[RED][i] = math.MaxInt
		dist[BLUE][i] = math.MaxInt
	}
	dist[RED][0], dist[BLUE][0] = 0, 0 //distance for node 0 is always 0.

	queue := []Node{{0, byte(RED)}, {0, byte(BLUE)}} //insert 0 nodes into the queue and start bfs traversal
	for len(queue) > 0 {
		color, vertex := queue[0].color, queue[0].vertex
		queue = queue[1:]

		//loop thru all the neighbouring nodes (alternate coloured edges (color^1))
		for _, nbr := range adj[color^1][vertex] {
			if dist[color][nbr] == math.MaxInt {
				dist[color][nbr] = dist[color^1][vertex] + 1
				queue = append(queue, Node{nbr, color ^ 1}) //append neighbouring node to queue after updating its distance
			}
		}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = min(dist[RED][i], dist[BLUE][i])
		if ans[i] == math.MaxInt {
			ans[i] = -1
		}

	}

	return ans

}

func buildAdj(curr int, edges [][]int) [][]int {
	adj := make([][]int, curr)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	return adj
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}