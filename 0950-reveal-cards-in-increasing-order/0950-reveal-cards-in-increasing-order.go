func deckRevealedIncreasing(deck []int) []int {
    if len(deck) < 2 {
        return deck
    }
	
	sort.Ints(deck)
	
	sorted := make([]int, len(deck)*2)
    
    y := 0
    for i:=0; i<len(sorted); i++ {
        if i%2 == 0 {
            sorted[i] = deck[y];    
            y++
        } 
        
    }
    sorted = sorted[:len(sorted)-1]
    
    lastEmpty := len(sorted)-2
    lastItem := len(sorted)-1
    for lastEmpty > 0 {
        sorted[lastEmpty] = sorted[lastItem]
        lastItem -= 1
        lastEmpty -= 2
    }
    
    return sorted[:len(deck)]
}