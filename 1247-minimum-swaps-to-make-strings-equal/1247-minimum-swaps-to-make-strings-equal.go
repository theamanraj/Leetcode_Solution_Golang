func minimumSwap(s1 string, s2 string) int {
    length := len(s1)
    b1 := make([]byte, length)
    b2 := make([]byte, length)
    countx, county := 0, 0
    
    for i := 0; i < len(s1); i++ {
        if s1[i] == 'x' { countx++ } else { county++ }
        if s2[i] == 'x' { countx++ } else { county++ }
        b1[i] = byte(s1[i])
        b2[i] = byte(s2[i])
    }
    
    if countx % 2 != 0 || county % 2 != 0 { return -1 }
    
    countx, county = 0, 0
    count := 0
    
    for ; countx < length - 1; countx++ {
        if b1[countx] == b2[countx] { continue }
        
        altIndex := -1
        for county = countx + 1; county < length; county++ {
            if b1[countx] == b1[county] && b2[countx] == b2[county] {
                b1[countx], b2[county] = b2[county], b1[countx]
                count++
                break
            } else if altIndex == -1 && b1[countx] == b2[county] && b2[countx] == b1[county] {
                altIndex = county
            }
        }
        
        if county == length {
            if altIndex != -1 {
                count += 2
                b1[countx], b1[altIndex] = b1[altIndex], b1[countx]
            } else {
                return -1
            }
        }
    }
    
    return count
}