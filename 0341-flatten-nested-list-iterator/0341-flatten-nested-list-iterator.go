/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type node struct {
    list []*NestedInteger
    p int
}


type NestedIterator struct {
    stack []*node
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{[]*node{&node{nestedList, 0}}} 
}

func (this *NestedIterator) Next() int {
    this.HasNext()
    c := this.stack[len(this.stack)-1]
    ret := c.list[c.p].GetInteger()
    c.p++
    return ret
}

func (this *NestedIterator) HasNext() bool {
    for len(this.stack) != 0 {
        c := this.stack[len(this.stack)-1]
        if c.p >= len(c.list) {
            this.stack = this.stack[:len(this.stack)-1]
        } else if c.list[c.p].IsInteger() {
            return true
        } else {
            this.stack = append(this.stack, &node{c.list[c.p].GetList(), 0})
            c.p++
        }
    }
    return false
}
