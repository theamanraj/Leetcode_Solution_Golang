/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    nums []int
    p int
}

func Constructor(iter *Iterator) *PeekingIterator {
    var res PeekingIterator
    for iter.hasNext() {
        res.nums = append(res.nums, iter.next())
    }    
    res.p = 0
    return &res
}

func (this *PeekingIterator) hasNext() bool {
    return this.p < len(this.nums)
}

func (this *PeekingIterator) next() int {
    this.p ++
    return this.nums[this.p - 1]
}

func (this *PeekingIterator) peek() int {
    return this.nums[this.p]
}