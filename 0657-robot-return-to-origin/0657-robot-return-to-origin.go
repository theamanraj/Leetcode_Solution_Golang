func judgeCircle(moves string) bool {
    var k,h int
    if len(moves)%2 != 0 {
        return false
    }
    for i:=0;i<len(moves);i++ {
        switch {
            case moves[i] == 'U':
                k++
            case moves[i] == 'D':
                k--
            case moves[i] == 'R':
                h++
            case moves[i] == 'L':
                h--
        }
    }
    return k==0&&h==0
    
}