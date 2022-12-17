type RandomizedCollection struct {
	values         []int
	indexesByValue map[int][]int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		indexesByValue: make(map[int][]int),
		values:         make([]int, 0),
	}
}

func (this *RandomizedCollection) Insert(val int) bool {
	this.values = append(this.values, val)

	indexes, exists := this.indexesByValue[val]
	this.indexesByValue[val] = append(indexes, len(this.values)-1)

	return !exists
}

func (this *RandomizedCollection) Remove(val int) bool {
	indexes := this.indexesByValue[val]
	if len(indexes) == 0 {
		return false
	}

	removedValueIdx := indexes[len(indexes)-1]
	indexes = indexes[:len(indexes)-1]

	if removedValueIdx < len(this.values)-1 {
		lastIdx := len(this.values) - 1
		valueToMove := this.values[lastIdx]

		this.values[removedValueIdx] = valueToMove

		for i, idx := range this.indexesByValue[valueToMove] {
			if idx == lastIdx {
				this.indexesByValue[valueToMove][i] = removedValueIdx
				break
			}
		}
	}

	this.values = this.values[:len(this.values)-1]

	if len(indexes) > 0 {
		this.indexesByValue[val] = indexes
	} else {
		delete(this.indexesByValue, val)
	}

	return true
}

func (this *RandomizedCollection) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}
