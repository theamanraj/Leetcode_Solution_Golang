type point struct {
	val  int
	snap int
}
type SnapshotArray struct {
	values        [][]point
	snapshotCount int
}

func Constructor(length int) SnapshotArray {
	values := make([][]point, length)
	for i := 0; i < length; i++ {
		values[i] = []point{
			point{
				val:  0,
				snap: -1,
			},
		}
	}
	return SnapshotArray{
		values:        values,
		snapshotCount: -1,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	p := this.values[index]
	if p[len(p)-1].snap == this.snapshotCount {
		p[len(p)-1].val = val
	} else {
		this.values[index] = append(this.values[index], point{
			val:  val,
			snap: this.snapshotCount,
		})
	}
}

func (this *SnapshotArray) Snap() int {
	this.snapshotCount++
	return this.snapshotCount
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	p := this.values[index]
	if snap_id <= -1 {
		return -1 // impossible
	}

	left, right := 0, len(p)-1
	for left <= right {
		mid := (left + right) / 2
		if p[mid].snap >= snap_id {
			right = mid - 1
		} else if p[mid].snap < snap_id {
			left = mid + 1
		}
	}

	return p[right].val
}