package main

var m map[string]int

func romanToInt(s string) int {
	// 1 before 5 and 10
	// 10 before 50 and 100
	// 100 before 500 and 1000
	// LVIII => 58
	// MCMXCIV => 1994 => M CM XC IV => 1000 + (1000 - 100) + (100 - 10) + (5 - 1)
	sum := 0
	fillMap()
	l := len(s)
	goNext := false
	for pos, char := range s {
		if goNext {
			goNext = false
			continue
		}
		charStr := string(char)
		if l-1 > pos {
			nextCharStr := string(s[pos+1])
			if charStr == "I" {
				if nextCharStr == "V" {
					sum += m["V"] - m["I"]
					goNext = true
					continue
				}
				if nextCharStr == "X" {
					sum += m["X"] - m["I"]
					goNext = true
					continue
				}
				sum += m[charStr]
				continue
			}
			if charStr == "X" {
				if nextCharStr == "L" {
					sum += m["L"] - m["X"]
					goNext = true
					continue
				}
				if nextCharStr == "C" {
					sum += m["C"] - m["X"]
					goNext = true
					continue
				}
				sum += m[charStr]
				continue
			}
			if charStr == "C" {
				if nextCharStr == "D" {
					sum += m["D"] - m["C"]
					goNext = true
					continue
				}
				if nextCharStr == "M" {
					sum += m["M"] - m["C"]
					goNext = true
					continue
				}
				sum += m[charStr]
				continue
			}
		}
		sum += m[charStr]
	}
	return sum
}
func fillMap() map[string]int {
	m = make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	return m
}
