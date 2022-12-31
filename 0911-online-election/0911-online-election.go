type TopVotedCandidate struct {
	time_person [][2]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var obj TopVotedCandidate
	var person_cnt map[int]int = make(map[int]int)
	var l int = len(persons)
	var max_cnt int = 0
	var max_person int = -1
	for i := 0;i < l;i++{
		person_cnt[persons[i]]++
		if person_cnt[persons[i]] >= max_cnt{
			max_cnt = person_cnt[persons[i]]
			max_person = persons[i]
		}
		obj.time_person = append(obj.time_person,[2]int{times[i],max_person})
	}
	return obj
}

func (this *TopVotedCandidate) Q(t int) int {
	var l int = len(this.time_person)
	find_idx := sort.Search(l, func(i int) bool {
		return this.time_person[i][0] > t
	})
	return this.time_person[find_idx - 1][1]
}