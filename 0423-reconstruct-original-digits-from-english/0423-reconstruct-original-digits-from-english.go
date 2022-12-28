func remove(str string, n int, count *[]int) {
    for _, s := range str {
        (*count)[s - 'a'] = (*count)[s - 'a'] - n 
    }
} 

func originalDigits(s string) string {
    count := make([]int, 26)
    nums := make([]int, 10)
    for _, c := range s {
        count[c - 'a'] ++
    }    
    nums[0] = count['z' - 'a'] // zero
    remove("zero", nums[0], &count)

    nums[2] = count['w' - 'a'] // two
    remove("two", nums[2], &count)
    
    nums[8] = count['g' - 'a'] // eight
    remove("eight", nums[8], &count)
    
    nums[6] = count['x' - 'a'] // six
    remove("six", nums[6], &count)
    
    nums[3] = count['h' - 'a'] // three
    remove("three", nums[3], &count)

    nums[4] = count['r' - 'a'] // eight
    remove("four", nums[4], &count)

    nums[1] = count['o' - 'a'] // one
    remove("one", nums[1], &count)

    nums[7] = count['s' - 'a'] // eight
    remove("seven", nums[7], &count)

    nums[5] = count['v' - 'a'] // eight
    remove("five", nums[5], &count)

    nums[9] = count['e' - 'a'] // eight
    remove("nine", nums[9], &count)

    res := ""
    for i := 0; i < 10; i ++ {
        for j := 0; j < nums[i]; j ++{
            res = res + string(i + '0')
        }
    }
    return res
}