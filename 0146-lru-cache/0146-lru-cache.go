import (
	"container/list"
	"fmt"
)

type entry struct {
	key int
	value int
}

type LRUCache struct {
	Capacity int
	list *list.List
	hashMap map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{}
	lruCache.list = list.New()
	lruCache.Capacity = capacity
	lruCache.hashMap = make(map[int]*list.Element, capacity)
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.hashMap[key]; ok {
		this.list.MoveToFront(elem)
		fmt.Printf("Get %d, Front (%d, %d), Back (%d, %d)\n", key,
			this.list.Front().Value.(*entry).key, this.list.Front().Value.(*entry).value,
			this.list.Back().Value.(*entry).key, this.list.Back().Value.(*entry).value)
		return elem.Value.(*entry).value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int)  {
	if this.Get(key) != -1 {
		// update
		this.hashMap[key].Value.(*entry).value = value
	} else {
		// evict tail if capacity is reached out
		if len(this.hashMap) == this.Capacity {
			fmt.Printf(" -- delete tail (%d, %d)\n", this.list.Back().Value.(*entry).key, this.list.Back().Value.(*entry).value)
			back := this.list.Back()
			this.list.Remove(back)
			delete(this.hashMap, back.Value.(*entry).key)
		}
		// insert new one
		e := this.list.PushFront(&entry{key,value})
		this.hashMap[key] = e
		fmt.Printf("Put (%d, %d), Front (%d, %d), Back (%d, %d), cap: %d len map: %d\n ", key, value,
			this.list.Front().Value.(*entry).key, this.list.Front().Value.(*entry).value,
			this.list.Back().Value.(*entry).key, this.list.Back().Value.(*entry).value,
			this.Capacity, len(this.hashMap))
	}
}