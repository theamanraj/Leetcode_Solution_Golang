/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    mergedList := &ListNode{}
    dummy := mergedList
    
    for list1 != nil || list2 != nil{
        
    if list1 == nil{
        dummy.Next = list2
        break
    }
    
        if list2 == nil{
        dummy.Next = list1
            break
    }
        if list1.Val > list2.Val{
            dummy.Next = list2
            list2 = list2.Next
        } else {
            dummy.Next = list1
            list1 = list1.Next 
        }
        dummy = dummy.Next
    }
    

    
    return mergedList.Next
    
}