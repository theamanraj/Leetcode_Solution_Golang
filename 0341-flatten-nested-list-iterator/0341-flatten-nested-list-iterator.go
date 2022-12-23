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

type NestedIterator struct {
    vals    []int
    idx int    
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    var nums []int
    flattern( nestedList, &nums)
    return &NestedIterator{
        vals: nums,        
    }
}

func flattern(  nestedList []*NestedInteger, nums *[]int){
    for _, nest := range nestedList{
        if nest.IsInteger(){
            *nums=append(*nums, nest.GetInteger())
            continue
        }
        flattern( nest.GetList(), nums)        
    }
}



func (this *NestedIterator) Next() int {
    if this.idx==len(this.vals) { return -1} //?
    v := this.vals[this.idx]
    this.idx++    
    return v
}

func (this *NestedIterator) HasNext() bool {
    return this.idx<len(this.vals)
}