type Value struct {
    v string
    ts int
}
type TimeMap struct {
    store map[string][]Value
}


func Constructor() TimeMap {
    s := make(map[string][]Value)
    return TimeMap{s}
}

// O(1)
func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := this.store[key]; !ok {
        this.store[key] = make([]Value, 0)
    }
    this.store[key] = append(this.store[key], Value{v:value,ts:timestamp})
}

// O(logN)
func (this *TimeMap) Get(key string, timestamp int) string {
    index := sort.Search(len(this.store[key]), func(i int) bool {
        return this.store[key][i].ts > timestamp
    })
    if index == 0 {
        return ""
    } 
    return this.store[key][index - 1].v
}