/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    tempNode := &ListNode{}
    curNode := tempNode
    var carryOn,temp int
    for l1 != nil || l2 != nil{
         i1 , i2 :=0, 0   
        if l1!=nil{            
            i1=l1.Val
        }  
        if l2!=nil{            
            i2=l2.Val
        }  
          lval := i1 + i2 + carryOn     
        carryOn = lval/10
        temp = lval%10
        curNode.Next=&ListNode{
            Val:temp,
        }    
        curNode = curNode.Next
         fmt.Println("carryon & temp value is",carryOn,temp,tempNode )
        if l1!=nil{
            l1=l1.Next
        }
        if l2!=nil {
            l2=l2.Next
        }  
    }
    if carryOn >0{
                curNode.Next=&ListNode{
            Val:carryOn,
        }    
        curNode = curNode.Next
    }
    return tempNode.Next    
}