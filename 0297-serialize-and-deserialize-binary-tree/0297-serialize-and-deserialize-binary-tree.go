/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    index int
}

func Constructor() Codec {
    return Codec{}
}

/*
    1
   / \
  2   3
  
  serialized: 1L2L#R#R3L#R#
*/

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "#"
    }
    
    val := strconv.Itoa(root.Val)
    left := this.serialize(root.Left)
    right := this.serialize(root.Right)
    return fmt.Sprintf("%sL%sR%s", val, left, right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if data[this.index] == '#' {
        this.index += 1
        return nil
    }
    
    i := this.index
    numStr := ""
    multiplier := 1
    if data[i] == '-' {
        i += 1
        multiplier = -1
    }
    
    // Take root value.
    for i<len(data) && data[i] >= '0' && data[i] <= '9' {
        numStr += string(data[i])
        i += 1
    }
    this.index = i
    num, _ :=  strconv.Atoi(numStr)
    currNode := &TreeNode{Val: multiplier*num}
    
    // Take left node.
    this.index += 1
    currNode.Left = this.deserialize(data)
    
    // Take right node.
    this.index += 1
    currNode.Right = this.deserialize(data)
    return currNode
}