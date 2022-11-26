func wordBreak(s string, wordDict []string) []string {
    if len(s) == 0 || len(wordDict) == 0 {
        return nil
    }
    mp := make(map[string]*list.List)
    reslist := dfs(s, wordDict, mp)

    var ret []string

    // Iterate through list and print its contents.
    for e := reslist.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
        ret = append(ret, e.Value.(string))
    }
    return ret
}

func dfs(s string, wordDict []string, mp map[string]*list.List) *list.List {
    if v, ok := mp[s]; ok {
        return v
    }

    res := list.New()
    if len(s) == 0 {
        res.PushBack("")
        return res
    }

    for _, word := range wordDict {
        if strings.HasPrefix(s, word) {
            sublist := dfs(s[len(word):], wordDict, mp)
            for e := sublist.Front(); e != nil; e = e.Next() {
                if len(e.Value.(string)) != 0 {
                    res.PushBack(word + " " + e.Value.(string))
                } else {
                    res.PushBack(word + "" + e.Value.(string))
                }
            }
        }
    }
    mp[s] = res
    return res
}