func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
    areaA:=(C-A) * (D-B)
    areaB:=(G-E) * (H-F)
    
    left := A
    if E > A {
        left = E }
    right := G
    if C < G {
        right = C }
    bottom := F
    if B > F{
        bottom = B }
    top := D
    if D > H {
        top = H }
    
    overlap := 0
    if right > left && top > bottom {
        overlap = (right - left ) * (top - bottom)
    }
    return areaA + areaB - overlap
}