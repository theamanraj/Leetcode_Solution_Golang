type MyCalendarTwo struct {
    events map[int]int
	keys   []int
}


func Constructor() MyCalendarTwo {
    var cal MyCalendarTwo
	cal.events = make(map[int]int)
	return cal
}

func (cal *MyCalendarTwo) removeElement(val int, event int) {
	cal.events[val] = cal.events[val] + event
	if cal.events[val] == 0 {
		delete(cal.events, val)
		for i, v := range cal.keys {
			if v == val {
				cal.keys = append(cal.keys[:i], cal.keys[i+1:]...)
			}
		}
	}
}

func (cal *MyCalendarTwo) addKey(val int, keys []int) {
	found := false
	for _, v := range keys {
		if v == val {
			found = true
		}
	}
	if !found {
		cal.keys = append(cal.keys, val)
	}
}

func (cal *MyCalendarTwo) addTime(val int, event int) {
	if v, found := cal.events[val]; found {
		cal.events[val] = v + event
	} else {
		cal.events[val] = event
		cal.addKey(val, cal.keys)
	}
}

func (cal *MyCalendarTwo) Book(start int, end int) bool {
	cal.addTime(start, 1)
	cal.addTime(end, -1)
	sort.Ints(cal.keys)

	active := 0
	for _, d := range cal.keys {
		active += cal.events[d]
		if active >= 3 {
			cal.removeElement(start, -1)
			cal.removeElement(end, 1)
			return false
		}
	}
	return true
}