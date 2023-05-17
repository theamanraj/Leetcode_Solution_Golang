/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    var rev *ListNode
    fast,slow := head,head
    sum,maxSum := 0,0
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    for slow != nil {
        temp := slow.Next
        slow.Next = rev
        rev = slow
        slow = temp
    }
    for head.Next != nil && rev != nil {
        sum = head.Val + rev.Val
        if sum > maxSum {
            maxSum = sum
        }
        head = head.Next
        rev = rev.Next
    }
    return maxSum
}