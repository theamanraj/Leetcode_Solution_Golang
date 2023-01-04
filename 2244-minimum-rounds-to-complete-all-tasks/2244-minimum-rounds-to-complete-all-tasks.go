func minimumRounds(tasks []int) int {
    sort.Ints(tasks)
    c := 1
    cnt := 0

    for i:=1;i<=len(tasks);i++ {
        
        if i == len(tasks) || tasks[i] != tasks[i-1] {
        
            if c == 1 {
                return -1
            } 
            
            cnt += c / 3 
            if c % 3 != 0 {
                cnt++
            } 
            
            c = 0
        }
        c++
    }
    
    return cnt
}