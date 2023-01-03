import (
    "fmt"
    "math/rand"
)

type Solution struct {
    n int
    subst map[int]int
}

func Constructor(N int, b []int) Solution {
	s := Solution{
		n:     N - len(b),
		subst: make(map[int]int),
	}

	for _, v := range b {
		s.subst[v] = -1
	}

	for _, v := range b {
		if v < s.n {
			for N > 0 {
				N--
				if _, ok := s.subst[N]; !ok {
					s.subst[v] = N
					break
				}
			}
		}
	}


	return s
}

func (this *Solution) Pick() int {
    r := rand.Intn(this.n)
    if val, ok := this.subst[r]; ok {
        return val
    }
    return r
}



/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */