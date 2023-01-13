/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Token struct {
    depth int
    val   int
}

func recoverFromPreorder(str string) *TreeNode {
    tokens := tokenize(str)
    
    stack := make([]*TreeNode, 0)
    for i := range tokens {
        fmt.Printf("token: %d (d:%d) \n", tokens[i].val, tokens[i].depth)
        currNode := TreeNode{
            Val: tokens[i].val,
            Left: nil,
            Right: nil,
        }
        if i > 0 {
            if tokens[i].depth > tokens[i - 1].depth {
                prevStackItem := stack[len(stack) - 1]
                prevStackItem.Left = &currNode
                fmt.Printf(" --Left to %d\n", prevStackItem.Val)
            } else if tokens[i].depth == tokens[i - 1].depth {
                stack = stack[:len(stack) - 1]
                prevStackItem := stack[len(stack) - 1]
                prevStackItem.Right = &currNode
            } else if tokens[i].depth < tokens[i - 1].depth {
                cutLen := tokens[i - 1].depth - tokens[i].depth + 1
                stack = stack[:len(stack) - cutLen]
                prevStackItem := stack[len(stack) - 1]
                prevStackItem.Right = &currNode
            }
        }
        stack = append(stack, &currNode)
    }
    
    for len(stack) > 1 {
        stack = stack[:len(stack) - 1]
    }
    
    return stack[0]
}

func tokenize(str string) []Token {
    tokens := make([]Token, 0)
    token := Token { depth: 0, val: 0 }
    d := 0
    v := ""
    i := 0
    for i < len(str) {
        // Detecting a beginnig of new depth seq
        if i > 0 && str[i - 1] != '-' && str[i] == '-' {
            for i < len(str) && str[i] == '-' {
                d += 1
                i += 1
            }
            token.depth = d
            d = 0
        }
        // Beginning of new number
        if i == 0 || str[i - 1] == '-' && str[i] != '-' {
            for i < len(str) && str[i] != '-' {
                v += string(str[i])
                i += 1
            }
            token.val, _ = strconv.Atoi(v)
            v = ""
        }
        
        tokens = append(tokens, token)
    }
    return tokens
}