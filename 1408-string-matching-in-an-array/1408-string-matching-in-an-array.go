func stringMatching(words []string) []string {
    res := []string{}
    for i, sub := range words{
        for j, word := range words{
            if i == j || len(sub)>len(word){
                continue
            }
            if strings.Contains(word, sub){
                res = append(res, sub)
                break
            }
        }   
    }
    return res
}