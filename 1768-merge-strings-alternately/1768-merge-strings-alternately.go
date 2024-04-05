func mergeAlternately(word1 string, word2 string) string {
    word := ""
    a, b := 0,0
    for{
    if (a < len(word1) && b < len(word2)){
        word = word + string(word1[a])+string(word2[b])
        a++
        b++
        continue
    } 
        break
    }
    v := checkGreater(word1,word2)
    if v{
          return word + sliceString(word1, a)
    }
    return word + sliceString(word2, b)
    }


func checkGreater(word1 string, word2 string) bool {
    return len(word1) > len(word2)
}

func sliceString (s string, index int) string {
    res := ""
    for i:=index;i<len(s);i++{
        res = res + string(s[i])
    }
    return res
}