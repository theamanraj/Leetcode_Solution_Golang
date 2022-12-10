type myChars []byte

func (inData myChars) Less(i int, j int) bool {
        return (inData[i] < inData[j])
}

func (inData myChars) Swap(i int, j int) {
        inData[i], inData[j] = inData[j], inData[i]
}

func (inData myChars) Len() int {
        return len(inData)
}

func SortChars(inStr string) string {
        lBytes := []byte(inStr)
        sort.Sort(myChars(lBytes))
        return string(lBytes)
}

func findAnagrams(s string, p string) []int {
    lMap := make(map[string]string)
    var lOut []int
    lLenP := len(p)
    lInputSort := SortChars(p)
    for i:=0;i<len(s);i++ {
        //slice i to lLenP
        if (i+lLenP) > len(s) {
            break
        }
        lSubStr := s[i:i+lLenP]
        lSubSort := ""
        //fmt.Println("lBuStr=", lSubStr)
        if val, ok := lMap[lSubStr]; ok {
            //we have already sorted this, get its value
            lSubSort = val
            
        } else {
            lSubSort = SortChars(lSubStr)
            lMap[lSubStr] = lSubSort
        }
        
        
        if lSubSort == lInputSort {
            lOut = append(lOut, i)
        }
        
    }
    return lOut
}