/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type IdSpace struct {
    record map[string]string
    currId int
}
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    idCountMap := make(map[string]int)
    space := &IdSpace{
        record: make(map[string]string),
        currId: 1,
    }
    result := make([]*TreeNode, 0)
    
    mark(root, idCountMap, space, &result)
    
    return result
}

func mark(subroot *TreeNode, idCountMap map[string]int, space *IdSpace, result *[]*TreeNode) string {
    if subroot == nil { return "0" }
    sid := strconv.Itoa(subroot.Val) + "-" +
        mark(subroot.Left, idCountMap, space, result) + "-" +
        mark(subroot.Right, idCountMap, space, result)
    id, ok := space.record[sid]
    if !ok {
        id = strconv.Itoa(space.currId)
        space.currId++
        space.record[sid] = id
    }
    idCountMap[id] += 1
    if idCountMap[id] == 2 {
        *result = append(*result, subroot)
    }
    
    return id
}