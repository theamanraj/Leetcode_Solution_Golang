/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	// a stack to store node, an output to store result
	stack, output := []*Node{root}, []int{}
	for len(stack) > 0 {
		root = stack[0] // get the first root
		output = append(output, root.Val) // store a value
		stack = stack[1:] // update stack, like unshift array
		if len(root.Children) > 0 {
			stack = append(root.Children, stack...) // prepend Children to stack
		}
	}
	return output
}