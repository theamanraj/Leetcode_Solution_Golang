/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// BFS to handle binary
// steps level by level: (1) set Next Pointer (2) enQueue (3)deQueue 
func connect(root *Node) *Node {
    if root == nil {
        return nil
    }    
    //build a queue  to insert the tree, one level by another level.
    queue := []*Node{root}
    
    // as far as the next level lenth is 0, stop the application.
    for len(queue)>0 {
        qLen := len(queue)
        
        //set Next pointer
        for i:=1;i<qLen; i++ {
            queue[i-1].Next=queue[i]            
        }
        
        //enQueue: add the next level nodes
        for i:=0;i<qLen; i++ {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)    
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            } 
        }
        
        //deQueue: remove current level nodes        
        queue = queue[qLen:]
    } 
     
    return root
}