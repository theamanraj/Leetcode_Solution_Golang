func isEvenOddTree(root *TreeNode) bool {
			// Edge case
			if root == nil {
				return false
			}
			
			levelOrder := LevelOrder(root)

			// check level 0
			if levelOrder[0][0]%2 == 0 {
				return false
			}

			// check level 1 ~ end(i starts at 0, so +2)
			for i, level := range levelOrder[1:] {
				// even and strictly decreasing order
				if (i+2)%2 == 0 {
					previousElement := level[0]
					if previousElement%2 != 0 {
						return false
					}

					for _, currentElement := range level[1:] {
						if currentElement%2 != 0 || previousElement <= currentElement {
							return false
						}
						previousElement = currentElement
					}
				} else {
					// odd and strictly increasing order
					previousElement := level[0]
					if previousElement%2 == 0 {
						return false
					}
					for _, currentElement := range level[1:] {
						if currentElement%2 == 0 || currentElement <= previousElement {
							return false
						}
						previousElement = currentElement
					}
				}
			}

			return true
		}

	// Normal Level Order Func
	func LevelOrder(root *TreeNode) [][]int {
		levelOrder := [][]int{}

		queue := []*TreeNode{root}
		for len(queue) != 0 {
			size := len(queue)
			level := []int{}

			for i := 0; i < size; i++ {
				node := queue[0]
				queue = append(queue[:0], queue[1:]...)

				if node.Left != nil {
					queue = append(queue, node.Left)
				}

				if node.Right != nil {
					queue = append(queue, node.Right)
				}

				level = append(level, node.Val)
			}

			levelOrder = append(levelOrder, level)
		}

		return levelOrder
	}