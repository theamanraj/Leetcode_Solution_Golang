type Node struct {
    Children map[byte]*Node
    End bool
}

type WordDictionary struct {
    root *Node
}

func NewNode() *Node {
    return &Node{
        Children: make(map[byte]*Node, 26),
    }
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{root: NewNode()}
}

func (this *WordDictionary) AddWord(word string)  {
    addWord(this.root, word, 0)
}

func addWord(node *Node, word string, i int) {
    if i >= len(word) {
        node.End = true
        return
    }
    c := word[i]
    if _, ok := node.Children[c]; !ok {
        node.Children[c] = NewNode()
    }
        
    addWord(node.Children[c], word, i+1)
}

func (this *WordDictionary) Search(word string) bool {
    return search(this.root, word, 0)
}

func search(node *Node, word string, i int) bool {
    if i >= len(word) {
        return node.End
    }
    
    c := word[i]
    if c != '.' {
        if _, ok := node.Children[c]; !ok {
            return false
        }
        return search(node.Children[c], word, i+1)
    } else {
        for _, c := range node.Children {
            if search(c, word, i+1) {
                return true
            }
        }
        return false
    }
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */