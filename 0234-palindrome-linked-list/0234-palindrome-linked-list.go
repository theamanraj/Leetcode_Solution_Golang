/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    
    // find length of the linked list
    l := 0
    for p1 := head; p1 != nil; p1 = p1.Next {
        l += 1
    }
    
    if l == 0 || l == 1 {
        return true
    }
    
    // break the list into two part
    // find the starting point of the second part of the linked list
    p2 := head
    for i := 0; i < l/2; i++ {
        p2 = p2.Next
    }
    
    // reverse the second part
    p2 = reverse(p2)
    p1:= head
    for p2 != nil {
        if p1.Val != p2.Val {
            return false
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    
    return true
}

func reverse(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    var pre *ListNode
    cur := head
    for cur != nil {
        tmp := cur.Next
        cur.Next = pre
        pre = cur
        cur = tmp
    }
    return pre
}