func splitListToParts(root *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	length := getLength(root)
	nodeLengthList := nodesPerGroup(length, k)
	for groupID := 0; groupID < k; groupID++ {
		res[groupID] = root
		for j := 0; root != nil && j < nodeLengthList[groupID]-1; j++ { //get to the tail
			root = root.Next
		}
		if root != nil { //split with nil
			tmp := root.Next
			root.Next = nil
			root = tmp
		}
	}
	return res
}

func getLength(head *ListNode) int {
	res := 0
	node := head
	for node != nil {
		res++
		node = node.Next
	}
	return res
}

func nodesPerGroup(length, k int) []int {
	groupMinLen := length / k
	left := length % k
	var res []int
	for i := 0; i < k; i++ {
		tmp := groupMinLen
		if left > 0 {
			tmp++
			left--
		}
		res = append(res, tmp)
	}
	return res
}