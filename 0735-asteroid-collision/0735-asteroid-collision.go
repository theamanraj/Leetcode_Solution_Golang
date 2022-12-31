func asteroidCollision(asteroids []int) []int {
    stack := []int{}
    for _, v := range asteroids {
        if len(stack) == 0 || v > 0 {
            stack = append(stack, v)
        } else {
            for {
                peek := stack[len(stack) - 1]
                if peek < 0 {
                    stack = append(stack, v); break
                } else if peek == -v {
                    stack = stack[:len(stack) - 1]; break
                } else if peek > -v {
                    break
                } else {
                    stack = stack[:len(stack) - 1]
                    if len(stack) == 0 { stack = append(stack, v); break }
                }
            }
        }
    }
    return stack
}