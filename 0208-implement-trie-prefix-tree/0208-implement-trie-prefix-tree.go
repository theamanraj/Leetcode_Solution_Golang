type Trie struct {
    links []*Trie
    isEnd bool
}


func Constructor() Trie {    
    return Trie{ links: make([]*Trie,26)}
}


func (this *Trie) Insert(word string)  {
    cur :=  this    
    for i:=0;i<len(word);i++{     
        cIdx :=word[i]-'a'
        if cur.links[cIdx]==nil{
            cur.links[cIdx]=&Trie{links:make([]*Trie, 26) }
        }
        cur = cur.links[cIdx]
    }
    cur.isEnd=true
}


func (this *Trie) Search(word string) bool {
    cur := this
    for i:=0;i<len(word);i++ {        
        if cur = cur.links[word[i]-'a']; cur== nil{
            return false
        }
    }
    return cur.isEnd    
}


func (this *Trie) StartsWith(prefix string) bool {
    cur := this
    for i:=0;i<len(prefix);i++ {        
        if cur = cur.links[prefix[i]-'a']; cur== nil{ return false}
    }
    return cur!=nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */