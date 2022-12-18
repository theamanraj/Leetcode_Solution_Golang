func invertArray( a []int) []int {
    for i:=0;i<len(a);i++ {
        a[i]=a[i]^1
    }
    return a
}

func reverseArray( a []int) []int {

    l := len(a)
    for i:=0;i<l/2;i++{
        tmp := a[i]
        a[i]=a[l-i-1]
        a[l-i-1]=tmp
    }
    return a
}
func flipAndInvertImage(A [][]int) [][]int {

    for i:=0;i<len(A);i++ {
        reverseArray(A[i])
        invertArray(A[i])
    }
    return A
}
