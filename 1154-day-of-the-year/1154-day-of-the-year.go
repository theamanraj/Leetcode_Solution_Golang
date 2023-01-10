func dayOfYear(date string) int {
    t, _ := time.Parse(time.RFC3339, date+"T00:00:00.000Z")
    return int(t.Sub(time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())).Hours()) / 24 + 1
}