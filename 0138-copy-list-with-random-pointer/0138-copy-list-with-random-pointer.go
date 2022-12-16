/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

if head == nil {
    return head
}

tmp := head
for tmp !=nil {
    newTmp :=  &Node{Val:tmp.Val, Next:tmp.Next, Random: nil}
    tmp.Next = newTmp
    tmp = newTmp.Next
}

tmp = head
for tmp != nil {
    if tmp.Next != nil && tmp.Random != nil{
        tmp.Next.Random = tmp.Random.Next
    }
    if tmp.Next != nil {
     tmp = tmp.Next.Next
    }
    
}

tmp = head
tmpClone := head.Next
retHead := tmpClone

for tmp != nil && tmpClone != nil {
    tmp.Next = tmpClone.Next
    if tmpClone.Next != nil {
        tmpClone.Next = tmpClone.Next.Next
    }
    tmp = tmp.Next
    tmpClone = tmpClone.Next
}

return retHead
}