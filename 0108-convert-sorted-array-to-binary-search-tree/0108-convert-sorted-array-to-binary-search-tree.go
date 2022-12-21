/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if nums == nil {
        return nil;
    }
    return helper(nums, 0, len(nums) - 1)
}

func helper (nums []int, low int, high int)  *TreeNode {
    if (low > high) { // Done
        return nil
    }
    mid := low + (high - low) / 2
    node := new(TreeNode)
    node.Val = nums[mid]
    node.Left = helper(nums, low, mid - 1)
    node.Right = helper(nums, mid + 1, high)
    return node
}