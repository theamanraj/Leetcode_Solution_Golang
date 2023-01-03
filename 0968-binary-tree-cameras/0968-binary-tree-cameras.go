func minCameraCover(root *TreeNode) int {
	ret := 0
	father := make(map[*TreeNode]*TreeNode)
	covered := make(map[*TreeNode]struct{})
	levels := make([][]*TreeNode, 0)

	// get the levels slice, build father map
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		next := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Left != nil {
				father[v.Left] = v
				next = append(next, v.Left)
			}
			if v.Right != nil {
				father[v.Right] = v
				next = append(next, v.Right)
			}
		}
		levels = append(levels, queue)
		queue = next
	}

	// go through from bottom to top
	for i := len(levels) - 1; i >= 0; i-- {
		queue = levels[i]
		for _, v := range queue {
			// if the node has been covered
			if _, ok := covered[v]; ok {
				continue
			}
			ret = ret + 1
			// if it has children not covered, put camera at it
			if _, ok := covered[v.Left]; !ok && v.Left != nil {
				putCamera(v, covered, father)
				continue
			}
			if _, ok := covered[v.Right]; !ok && v.Right != nil {
				putCamera(v, covered, father)
				continue
			}
			// if it has a father, put camera at its father
			if fNode, ok := father[v]; ok && fNode != nil {
				putCamera(fNode, covered, father)
				continue
			}
			putCamera(v, covered, father)
		}
	}
	return ret
}

func putCamera(node *TreeNode, covered map[*TreeNode]struct{}, father map[*TreeNode]*TreeNode) {
	covered[node] = struct{}{}
	covered[father[node]] = struct{}{}
	covered[node.Left] = struct{}{}
	covered[node.Right] = struct{}{}
}