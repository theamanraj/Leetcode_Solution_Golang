func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    ruleKeys := map[string]int{
        "type":0,
        "color":1,
        "name":2,
    }
    
    result := 0
    
    for _,v := range items{
        if v[ruleKeys[ruleKey]] == ruleValue{
            result++
        }
    }
    
    return result
}