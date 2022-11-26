type Trie struct {
    childs[26] *Trie
    Count int
}

func (t *Trie) Insert(word string) {
    node  := t
    lenW := len(word)
    for i:= 0; i < lenW; i++ {
        idx := word[i] - 'a'
        if node.childs[idx] == nil {
            node.childs[idx] =  &Trie{}
        }
        node = node.childs[idx]
    }
    node.Count++
}

func (t *Trie)StartWith(prefix string) *Trie {
    node  := t
    lenP := len(prefix)
    for i:= 0; i < lenP;i++ {
        idx := prefix[i] - 'a'
        if node.childs[idx] == nil {
            return nil
        }
        node = node.childs[idx]
    }
    return node
}

func findWords(board [][]byte, words []string) []string {
    trie := &Trie{}
    //build trie tree
    for i :=0; i < len(words); i++ {
        trie.Insert(words[i])
    }
    
    result := map[string]bool{}
    //start to bactracing
    for x := 0; x < len(board); x++ {
        for y := 0; y < len(board[0]); y++ {
            backTrace(&board, x, y,  trie, "", &result)
        }
    }
    
    list := []string{}
    for k, _ := range result {
        list = append(list, k)
    }
    return list
}

func backTrace(board *([][]byte),  x, y int, trie *Trie, prefix string, result *map[string]bool) {
    lenX := len(*board)
    lenY := len((*board)[0])
   
    if x >= 0 && x < lenX && y >= 0 && y < lenY {
         char := (*board)[x][y]
        if char ==  '#' {
            return
        }
         (*board)[x][y] = '#' //mark visited
        newPrefix := prefix + string(char)
        node := trie.StartWith(newPrefix)

        if node != nil {
            backTrace(board , x-1, y, trie, newPrefix, result)
            backTrace(board , x+1, y, trie, newPrefix, result)  
            backTrace(board , x, y -1 , trie, newPrefix, result)
            backTrace(board , x, y+1, trie, newPrefix, result)  

            
            if node.Count > 0 {
                (*result)[newPrefix] = true
            }
        }
        (*board)[x][y] = char
        
    }
}