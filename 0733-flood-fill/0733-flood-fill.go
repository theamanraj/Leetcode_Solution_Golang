func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    solve(image, sr, sc, image[sr][sc], newColor)
    return image
}

func solve(inp [][]int, sr int, sc int, prevColor int, newColor int) {
    if sr < 0 || sc < 0 || sr >= len(inp) || sc >= len(inp[0]) || inp[sr][sc] != prevColor || prevColor == newColor {
        return
    }
    inp[sr][sc] = newColor
    solve(inp, sr+1, sc, prevColor, newColor)
    solve(inp, sr-1, sc, prevColor, newColor)
    solve(inp, sr, sc+1, prevColor, newColor)
    solve(inp, sr, sc-1, prevColor, newColor)
}