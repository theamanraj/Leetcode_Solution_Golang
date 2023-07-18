func dayOfTheWeek(day int, month int, year int) string {
	t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	if month < 3 {
		year -= 1
	}
	date := []string{"Sunday","Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return date[(year+year/4+year/400-year/100+t[month-1]+day)%7]
}