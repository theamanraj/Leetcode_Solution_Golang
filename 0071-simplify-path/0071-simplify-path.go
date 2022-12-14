func simplifyPath(path string) string {
    var stack []string
    res, arr := "", strings.Split(path, "/")
    
    for i := 0; i < len(arr); i++ {
        if arr[i] == ".." && len(arr) > 0 {
            if len(stack) > 0 {
                stack = stack[:len(stack) - 1]
            }            
        } else if arr[i] != "" && arr[i] != "." && arr[i] != ".." {
            stack = append(stack, arr[i])
        }
    }
    
    if len(stack) == 0 {
        return "/"
    }
    
    for len(stack) > 0 {
        res += "/" + stack[0]
        stack = stack[1:]
    }
    
    return res
}