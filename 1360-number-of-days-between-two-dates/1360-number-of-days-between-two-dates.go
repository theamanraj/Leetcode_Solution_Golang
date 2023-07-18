func daysBetweenDates(date1 string, date2 string) int {
    t1, _ := time.Parse("2006-01-02", date1)
    t2, _ := time.Parse("2006-01-02", date2)
    if t1.After(t2) {
        t1, t2 = t2, t1
    }
    d := t2.Sub(t1)
    return int(d.Hours()) / 24
}