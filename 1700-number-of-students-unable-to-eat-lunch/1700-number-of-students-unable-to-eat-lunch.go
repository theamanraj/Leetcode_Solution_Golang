func countStudents(students []int, sandwiches []int) int {
	misses := 0
	for misses != len(students) {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			misses = 0
		} else {
			var s int
			s, students = students[0], students[1:]
			students = append(students, s)
			misses++
		}
	}
	return len(students)
}