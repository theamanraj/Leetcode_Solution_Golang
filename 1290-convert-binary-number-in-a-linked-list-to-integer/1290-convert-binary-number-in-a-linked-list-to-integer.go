/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
  if head == nil {
    return 0
  }
  str := ""
  var node *ListNode = head
  for node != nil {
    str += strconv.Itoa(node.Val)
    node = node.Next
  }
  ans, _ := strconv.ParseInt(str,2,32)
  return int(ans)
}