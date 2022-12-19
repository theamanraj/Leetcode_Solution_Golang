type Node struct {
    keys map[string]int
    val  int
    next *Node
    prev *Node
}

type AllOne struct {
    // key2node allows to find node by the key
    key2node map[string]*Node
    // val2node allows to find node for a particular value
    val2node map[int]*Node
    // Head points to a node with min value
    head *Node
    // Tail points to a node with max value
    tail *Node
}

/** Initialize your data structure here. */
func Constructor() AllOne {
    return AllOne{
        key2node: make(map[string]*Node),
        val2node: make(map[int]*Node),
        head: nil,
        tail: nil,
    }
}

// newNode returns a new node for value v and key k
func newNode(k string, v int) *Node {
    n := &Node{
                keys: make(map[string]int),
                val: v,
                next: nil,
                prev: nil,
            }
    n.keys[k] = v
    return n
}

// insert adds node to the double-linked list after the base, or before head if the base is nil
func (this *AllOne) insert(base *Node, node *Node) {
    if node == nil {
        return
    }
    
    if base == nil {
        // We want to insert at head
        node.next = this.head
        if node.next == nil {  
            // This also must be our new tail
            this.tail = node
        } else {
            // Not a tail, so point next element back to us
            node.next.prev = node
        }
        
        this.head = node
        // Just in case clean prev pointer for the head, in case provided node had some garbage in its structure
        this.head.prev = nil
        return
    }
    
    
    next := base.next
    base.next = node
    node.prev = base
    node.next = next
    if next != nil {
        next.prev = node
    } else {
        // There was nothing after this element, so this becomes our new tail.
        this.tail = node
    }
}

// remove removes node from the double-linked list
func (this *AllOne) remove(node *Node) {
    if node == nil {
        return
    }
    
    prev := node.prev
    next := node.next
    
    if prev == nil && next == nil {
        // Last element in the list, just reset head/tail pointers.
        this.tail = nil
        this.head = nil
        return
    }
    
    if node == this.tail {
        // If we are at the tail, simply repoint tail and reset its next pointer
        this.tail = prev
        this.tail.next = nil
        return
    }
    
    if node == this.head { 
        // If we are at the head, simply repoint head and reset its prev pointer
        this.head = next
        this.head.prev = nil
        return
    }
    
    // At this point we know we are at the middle of the list, just swap pointers
    prev.next = next
    next.prev = prev
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
    if key == "" {
        return
    }
    
    node, ok := this.key2node[key]
    if !ok {
        // It is a new key, let's add it with value 1.
        if n, ok := this.val2node[1]; ok {
            // There is already a node for value 1, add the key to its list. 
            n.keys[key] = 1
            this.key2node[key] = n
        } else {
            // No node for the value 1, must create a new one and add it as a head of the list, update mappings.
            n := newNode(key, 1)
            this.key2node[key] = n
            this.val2node[1] = n
            this.insert(nil, n)
        }
        
        return
    }
    
    // It is an existing key, so increment the value and determine what node it should be place into.
    next := node.next
    newVal := node.val + 1
    if next == nil || next.val > newVal {
        // Next node either does not exist, or its value more than current value + 1.
        // We must create a new node with new value and place it after the current node. Update mappings for the key and new value.
        n := newNode(key, newVal)
        this.key2node[key] = n
        this.val2node[newVal] = n
        this.insert(node, n)
    } else {
        // Next node has the right value to just be added to. We don't have to do much.
        next.keys[key] = newVal
        this.key2node[key] = next
    }
    
    // Must clean up old node from the key, since it does not belong to it anymore.
    delete(node.keys, key)
    if len(node.keys) == 0 {
        // That was the last key in the node, delete the node.
        this.remove(node)
        delete(this.val2node, node.val)
    }
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
    if key == "" {
        return
    }
    
    // No key, nothing to decrement
    node, ok := this.key2node[key]
    if !ok {
        return
    }
    
    newVal := node.val - 1    
    prev := node.prev
    
    if newVal == 0 {
      // If value dropped to 0, delete this key as known.
      delete(this.key2node, key)
    } else if prev == nil || prev.val < newVal {
        // Value is not zero, but previous node on the list is smaller than newVal,
        // so we need to insert a new node with new value after the previous one for this key.
        n := newNode(key, newVal)
        this.key2node[key] = n
        this.val2node[newVal] = n
        this.insert(prev, n)
    }  else {
        // Previous node value is what we are looking for, just add to it.
        prev.keys[key] = newVal
        this.key2node[key] = prev
    }
    
    // Need to clean up node from the old bucket.
    delete(node.keys, key)
    if len(node.keys) == 0 {
        // Old node is empty, delete it
        delete(this.val2node, node.val)
        this.remove(node)
    }
    
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
    if this.tail != nil {
        // Just pick first element in the map and return it.
        for k, _ := range this.tail.keys {
            return k
        }
    }
    return ""
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
    if this.head != nil {
        // Just pick first element in the map and return it.
        for k, _ := range this.head.keys {
            return k
        }
    }
    return ""
}