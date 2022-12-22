/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	T := []*Node{}
	if root != nil {
		T = append(T, root)
	}
	rt := []int{}

	for {
		if len(T) == 0 {
			break
		}

		node := T[len(T)-1]
		T = T[:len(T)-1]
		T = append(T, node.Children...)
		rt = append(rt, node.Val)
	}

	for i, j := 0, len(rt)-1; i < j; i, j = i+1, j-1 {
		rt[i], rt[j] = rt[j], rt[i]
	}

	return rt
}