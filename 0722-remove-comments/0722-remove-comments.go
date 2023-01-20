type Mode int
const (
    None Mode = iota
    Line
    Block
)

func removeComments(source []string) []string {
    output := []string{}
    
    mode := None
    i:=0
    j:=0
    out := 0
    output = append(output, "")
    for i < len(source) {
        str := source[i]
        if j == len(str) {
            // reached to end of line
            i++
            j=0
            if mode == Line {
                mode = None
            }
            if mode != Block && len(output[len(output)-1]) > 0 {
                output = append(output,"")    
                out++
            }
            continue
        }
        switch mode {
            case None:
                if str[j] == '/' {
                    if j+1 < len(str) && str[j+1] == '/' {
                        // line comment starts
                        mode = Line
                        j=j+2
                        continue
                    } else if j + 1 < len(str) && str[j+1] == '*' {
                        // block comment starts
                        j=j+2
                        mode = Block
                        continue
                    }
                }
                output[out] = fmt.Sprintf("%s%c",output[out], str[j])
                j++
            case Block:
                if str[j] == '*' {
                    if j+1 < len(str) && str[j+1] == '/' {
                        // end of block comment
                        mode = None
                        j=j+2
                        continue
                    }
                }
                j++
            case Line:
                j++
        }
    }
    
    // Check last
    if len(output[len(output)-1]) == 0 {
        output = output[:len(output)-1]
    }
    return output
}