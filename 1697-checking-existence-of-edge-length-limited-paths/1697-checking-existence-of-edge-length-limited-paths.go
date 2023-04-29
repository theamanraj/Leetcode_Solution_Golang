type UnionFind struct {
	parent, depth []int
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	i, u, ans := 0, NewUnionFind(n), make([]bool, len(queries))
	for _, q := range queries {
		for i < len(edgeList) && edgeList[i][2] < q[2] {
			u.join(edgeList[i][0], edgeList[i][1])
			i++
		}
		if u.root(q[0]) == u.root(q[1]) {
			ans[q[3]] = true
		}
	}
	return ans
}

func NewUnionFind(n int) *UnionFind {
	parent, depth := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i], depth[i] = i, 0
	}
	return &UnionFind{parent, depth}
}

func (u *UnionFind) join(a, b int) {
	a, b = u.root(a), u.root(b)
	if u.depth[a] < u.depth[b] {
		a, b = b, a
	} else if u.depth[a] == u.depth[b] {
		u.depth[a]++
	}
	u.parent[b] = a
}

func (u *UnionFind) root(n int) int {
	if u.parent[n] == n {
		return n
	}
	u.parent[n] = u.root(u.parent[n])
	return u.parent[n]
}