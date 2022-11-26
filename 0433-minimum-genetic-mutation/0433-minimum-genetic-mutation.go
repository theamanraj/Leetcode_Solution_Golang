func minMutation(beginWord string, endWord string, wordList []string) int {
    adjList := process(wordList,beginWord) //map[string][]string
     if _, valid := adjList[endWord]; !valid {
        return -1
    }    
    queue := []string{}
    visited := make(map[string]struct{})
    queue = append(queue, beginWord)
    path := 1
    for len(queue) > 0 {
        l := len(queue)
        for i :=0;i<l;i++{
            top := queue[0]
            queue = queue[1:]
            for _, w := range adjList[top] {
                if _, ok:= visited[w] ; ok {
                    continue
                }
                visited[w] = struct{}{}
                if w == endWord {
                    return path
                }
                queue = append(queue, w)
            }
        }
        path++
    }
    return -1
}

func process( wordList []string, start string ) map[string][]string {
    ret := make(map[string][]string)
    genes := []byte{'A','C','T','G'}
    for _,w := range wordList {
        ret[w] = []string{}
    }
    if _,ok := ret[start] ; !ok {
        wordList = append(wordList, start)
        ret[start] = []string{}
    }
    for _,w := range wordList {
        for i := range w {
            byteRep := []byte(w)
            for _,j := range genes {
                byteRep[i] = byte(j) 
                if string(byteRep) == w {
                    continue
                }
                if _, ok := ret[string(byteRep)] ; ok {
                    ret[w] = append(ret[w], string(byteRep))
                }
            }
            
        }
    }
    return ret
    
}