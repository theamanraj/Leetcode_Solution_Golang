func majorityElement(nums []int) []int {
    var a, b int
    ac := 0
    bc := 0
    for _, num := range nums {
        if ac == 0 {
            if bc == 0 || b != num {
                a = num
                ac++
            } else {
                bc++
            }
        } else {
            if a == num {
                ac++
            } else {
                if bc == 0 {
                    b = num
                    bc++
                } else {
                    if b == num {
                        bc++
                    } else {
                        ac--
                        bc--
                    }
                }
            }
        }
    }
    
    var rst []int
    if ac > 0 {
        cnt := 0
        for _, num := range nums {
            if num == a {
                cnt++
            }
        }
        if cnt > len(nums)/3 {
            rst = append(rst, a)
        }
    }
    if bc > 0 {
        cnt := 0
        for _, num := range nums {
            if num == b {
                cnt++
            }
        }
        if cnt > len(nums)/3 {
            rst = append(rst, b)
        }
    }
    return rst
}