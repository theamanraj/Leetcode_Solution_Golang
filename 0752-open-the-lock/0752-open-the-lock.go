func openLock(deadends []string, target string) int {
    mapper := make(map[string]int)
    for _, v := range deadends {
        mapper[v] = 1
    }
    if mapper["0000"] == 1 {
        return -1
    }
    if target == "0000" {
        return 0
    }
    return BFS([]string{"0000"}, mapper, target)
}

func BFS(cur []string, be map[string]int, target string) int {
    ans := 0
    for len(cur) > 0 {
        tem := []string {}
        ans += 1
        for _, v := range cur {
            next := Next(v)
            for _, vv := range next {
                if vv == target {
                    return ans
                }
                if be[vv] == 0 {
                    be[vv] = 1
                    tem = append(tem, vv)
                }
            }
        }
        cur = []string{}
        cur = append(cur, tem...)
    }
    return -1
}

func Next(s string) []string {
    ans := []string{}
    for i := range s {
        if s[i] != '0' && s[i] != '9' {
            t := s
            t = t[:i] + string(t[i] + 1) + t[i+1:]
            
            ans = append(ans, t)
            t = s
            t = t[:i] + string(t[i] - 1) + t[i+1:]
            ans = append(ans, t)
        } else {
            if s[i] == '0' {
                t := s
                t = t[:i] +  string(t[i] + 1) + t[i+1:]
                ans = append(ans, t)
                t= s
                t = t[:i] + string(t[i] + 9) + t[i+1:]
                ans = append(ans, t)
            } else {
                if s[i] == '9' {
                    t := s
                    t = t[:i] + string(t[i] - 1) + t[i+1:]
                    ans = append(ans, t)
                    t = s
                    t = t[:i] + string(t[i] - 9) + t[i+1:]
                    ans = append(ans, t)
                }
            }
        }
    }
    return ans
}
