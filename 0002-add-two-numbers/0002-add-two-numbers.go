/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var temp = &ListNode{}
    
    current := temp
    carry := 0
    for(l1 != nil || l2 != nil || carry > 0){
        sum := 0
        if (l1 != nil){
           sum = l1.Val 
        }
        if (l2 != nil){
           sum += l2.Val 
        }
        sum += carry
        carry = sum/10                 
        current.Next=&ListNode{
            Val:sum%10,
        }
        current = current.Next
        if (l1 != nil) {
            l1=l1.Next
        }
        if (l2 != nil) {
            l2=l2.Next
        }
    }
    return temp.Next
}