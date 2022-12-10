import "sort"

type String []rune

func (s String) Len() int           { return len(s) }
func (s String) Less(i, j int) bool { return s[i] < s[j] }
func (s String) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func sortString(inp string) string {
	r := []rune(inp)
	sort.Sort(String(r))
	return string(r)
}

func findCharFreq(inp string) []int {
    out := make([]int, 26)
    for _, each := range inp {
        out[each-'a'] += 1
    }
    return out
}

func isAnagram(temp, pFreq []int) bool {
    for idx, freq := range temp {
        if pFreq[idx] != freq { return false }
    }
    return true
}

// K is the length of p
// N is the length of s
func findAnagrams(s string, p string) []int {
    if len(s) < len(p) { return []int{} }
    pFreq := findCharFreq(p) // K is the length of p; K operations
    out := []int{}
    temp := findCharFreq(string(s[:len(p)])) // K Steps (once)
    for i:= 0; i <= len(s)-len(p); i++ { // N-K Steps
        if isAnagram(temp, pFreq) { // Checks for anagram K steps
            out = append(out, i)
        }
		temp[s[i]-'a'] -= 1 // Constant update
		if i + len(p) >= len(s) { //constant check
			continue
		}
		temp[s[i+len(p)]-'a'] += 1 //Constant update
    }
    return out
}
