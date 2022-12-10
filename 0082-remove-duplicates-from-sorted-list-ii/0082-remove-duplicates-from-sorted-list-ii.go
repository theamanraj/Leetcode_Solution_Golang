/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    dummy:=&ListNode{Next:head}
    pre, cur := dummy, dummy.Next    
    for cur!=nil && cur.Next!=nil{
        if cur.Val==cur.Next.Val{
            for cur.Next!=nil && cur.Val==cur.Next.Val{ cur=cur.Next } // This is the key!
            cur, pre.Next=pre.Next,cur.Next
            continue
        }
        pre, cur =cur, cur.Next
    }    
    return dummy.Next    
}