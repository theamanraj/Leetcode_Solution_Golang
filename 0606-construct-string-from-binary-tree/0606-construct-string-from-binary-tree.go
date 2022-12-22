func preOrder(t *TreeNode) string {
	if t == nil {
		return ""
	}
	var left, right string //default to be empty string

	// pre-order traversing
	if t.Left != nil {
		left = preOrder(t.Left)
	}

	if t.Right != nil {
		right = preOrder(t.Right)
	}

	// if left is empty, it should be reserved in preorder-traversing as "()"
	// if right is empty, it can be omit in preorder-traversing (right is default to be empty).
	if left == "" {
		if right != "" {
			left = "()"
		}
	}

	return "(" + strconv.Itoa(t.Val) + left + right + ")"
}

func tree2str(t *TreeNode) string {
	result := preOrder(t)

	if len(result) >= 2 {
		return result[1 : len(result)-1] // split the outermost layer of parenthesis.
	} else {
		return result
	}
}