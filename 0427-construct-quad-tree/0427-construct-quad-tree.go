/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

package main

func construct(grid [][]int) *Node {
	leaves := map[int]*Node{
		0: {
			Val:    false,
			IsLeaf: true,
		},
		1: {
			Val:    true,
			IsLeaf: true,
		},
	}
	var dfs func(i, j, n int) *Node
	dfs = func(i, j, n int) *Node {
		if n == 1 {
			return leaves[grid[i][j]]
		}
		n >>= 1
		quads := [4]*Node{dfs(i, j, n), dfs(i, j+n, n), dfs(i+n, j, n), dfs(i+n, j+n, n)}
		for i := 0; i < len(quads); i++ {
			if !(quads[i].IsLeaf && quads[i].Val == quads[0].Val) {
				return &Node{
					Val:         true, // any value
					IsLeaf:      false,
					TopLeft:     quads[0],
					TopRight:    quads[1],
					BottomLeft:  quads[2],
					BottomRight: quads[3],
				}
			}
		}
		return quads[0]
	}
	return dfs(0, 0, len(grid))
}