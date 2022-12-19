func validUtf8(data []int) bool {
    i, n := 0, len(data)
    for i < n {
        if isOneByteCharacter(data[i]) {
            i++
        } else if i + 1 < n && isTwoByteCharacter(data[i], data[i+1]){
            i+=2
        } else if i + 2 < n && isThreeByteCharacter(data[i], data[i+1], data[i+2]) {
            i += 3
        } else if i + 3 < n && isFourByteCharacter(data[i], data[i+1], data[i+2], data[i+3]) {
            i += 4
        } else {
            return false
        }
    }
    
    return true
}


func isOneByteCharacter(b int) bool {
    return ((b >> 7) == 0)
}

func isTwoByteCharacter(b1, b2 int) bool {
    if (b1 >> 5) != 6 || (b2 >> 6) != 2 {
        return false
    }
    return true
}

func isThreeByteCharacter(b1, b2, b3 int) bool {
    if (b1 >> 4) != 14 || (b2 >> 6) != 2 || (b3 >> 6) != 2 {
        return false
    }
    return true
}

func isFourByteCharacter(b1, b2, b3, b4 int) bool {
    if (b1 >> 3) != 30 || (b2 >> 6) != 2 || (b3 >> 6) != 2 || (b4 >> 6) != 2 {
        return false
    }
    return true
}