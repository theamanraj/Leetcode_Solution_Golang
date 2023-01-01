import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type RandomizedSet struct {
	m    map[int]int
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		m:    make(map[int]int),
		nums: []int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	this.m[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}

	length := len(this.nums)
	p := this.m[val]
	this.m[this.nums[length-1]] = p
	this.nums[p], this.nums[length-1] = this.nums[length-1], this.nums[p]
	this.nums = this.nums[:length-1]
	delete(this.m, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}