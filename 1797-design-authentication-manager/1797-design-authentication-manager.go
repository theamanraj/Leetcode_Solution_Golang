type Token struct {
  id string
  expiry int
  
  index int // index of the item in the heap
}

type PriorityQueue []*Token

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
  return pq[i].expiry < pq[j].expiry
}
func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
  pq[i].index = i // maintain indices after the swap
  pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
  item := x.(*Token)
  item.index = len(*pq)
  *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
  var item *Token
  n := len(*pq)
  *pq, item = (*pq)[:n-1], (*pq)[n-1]
  item.index = -1 // for safety
  return item
}

func (pq PriorityQueue) Top() interface{} {
  return pq[0]
}

type AuthenticationManager struct {
  ttl int
  tokens map[string]*Token // tokens indexed by its id
  heap *PriorityQueue // tokens in a min heap based on the expiry
}


func Constructor(timeToLive int) AuthenticationManager {
  return AuthenticationManager{
    ttl: timeToLive,
    tokens: make(map[string]*Token),
    heap: &PriorityQueue{},
  }
}


func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {
  token := &Token{id: tokenId, expiry: currentTime + this.ttl}
  this.tokens[tokenId] = token
  heap.Push(this.heap, token)
}


func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {
  token, ok := this.tokens[tokenId]
  if !ok || token.expiry <= currentTime {
    return // token doesn't exist; can't renew
  }
  token.expiry = currentTime + this.ttl
  heap.Fix(this.heap, token.index)
}


func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
  for this.heap.Len() > 0 && this.heap.Top().(*Token).expiry <= currentTime {
    heap.Pop(this.heap)
  }
  return this.heap.Len()
}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */