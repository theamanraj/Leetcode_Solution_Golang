type RecentCounter struct {
    t []int
}


func Constructor() RecentCounter {
    return RecentCounter{ []int{} }
}


func (this *RecentCounter) Ping(t int) int {
    this.t = append(this.t, t)
    for len(this.t) > 0 && this.t[0] + 3000 < t { this.t = this.t[1:] }
    return len(this.t)
}