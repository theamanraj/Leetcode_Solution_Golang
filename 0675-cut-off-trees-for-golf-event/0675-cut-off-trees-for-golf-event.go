import "sort"
type point struct {
	x int
	y int
}

func (p point) equal(o point) bool {
	return p.x == o.x && p.y == o.y
}

type Queue []point

func (q Queue) push(p point) Queue {
	return append(q, p)
}

func (q Queue) top() point {
	return q[0]
}

func (q Queue) pop() Queue {
	return q[1:]
}

func (q Queue) empty() bool {
	return len(q) == 0
}

type Arrive struct {
	v []bool
	w int
	h int
}

func (a *Arrive) arrive(p point) {
	a.v[p.x*a.h+p.y] = true
}

func (a Arrive) hasArrived(p point) bool {
	return a.v[p.x*a.h+p.y]
}

func travel(forest [][]int, start point, end point) int {
	if start.equal(end) {
		return 0
	}
	w := len(forest)
	h := len(forest[0])
	arrived := &Arrive{
		v: make([]bool, h*w),
		h: h,
		w: w,
	}
	queue := make(Queue, 0)

	dirs := [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	queue = queue.push(start)

	ret := make([][]int, w)
	for i := 0; i < w; i++ {
		ret[i] = make([]int, h)
	}

	for !queue.empty() {
		cur := queue.top()
		queue = queue.pop()
		for _, dir := range dirs {
			newPos := point{
				x: cur.x + dir[0],
				y: cur.y + dir[1],
			}
			if newPos.x < 0 || newPos.y < 0 || newPos.x >= w || newPos.y >= h {
				continue
			}
			if forest[newPos.x][newPos.y] != 0 && !arrived.hasArrived(newPos) {
				if newPos.equal(end) {
					return ret[cur.x][cur.y] + 1
				}
				arrived.arrive(newPos)
				queue = queue.push(newPos)
				ret[newPos.x][newPos.y] = ret[cur.x][cur.y] + 1
			}
		}
	}

	return -1
}

func cutOffTree(forest [][]int) int {
	treeMap := make(map[int]point)
	values := make([]int, 0)
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			if forest[i][j] > 1 {
				treeMap[forest[i][j]] = point{
					x: i,
					y: j,
				}
				values = append(values, forest[i][j])
			}
		}
	}

	sort.Ints(values)

	curPos := point{
		x: 0,
		y: 0,
	}

	step := 0
	for len(values) > 0 {
		v := values[0]
		nextPos := treeMap[v]

		dis := travel(forest, curPos, nextPos)
		if dis == -1 {
			return -1
		}

		step += dis
		curPos = nextPos
		values = values[1:]
	}
	return step
}