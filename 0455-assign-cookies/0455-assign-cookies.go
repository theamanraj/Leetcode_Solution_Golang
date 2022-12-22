import "sort"

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    chil, cook := 0, 0
    for chil < len(g) && cook < len(s) {
        if g[chil] <= s[cook] {
            chil++
        }
        cook++                    //increase num of cookies until the 1st child is content
    }
    return chil
}